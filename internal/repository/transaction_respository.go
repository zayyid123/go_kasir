package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/internal/model"
)

type TransactionRepository interface {
	Checkout(items []model.CheckoutItem) (*model.Transaction, error)
}

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("checkout items cannot be empty")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Insert transaction first (total_amount 0)
	var (
		transactionID int
		createdAt     string
	)

	err = tx.QueryRow(`
		INSERT INTO transaction (total_amount)
		VALUES (0)
		RETURNING id, created_at
	`).Scan(&transactionID, &createdAt)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Prepare statement insert transaction detail
	stmtDetail, err := tx.Prepare(`
		INSERT INTO transaction_detail
		(transaction_id, product_id, quantity, subtotal)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`)
	if err != nil {
		return nil, err
	}
	defer stmtDetail.Close()

	totalAmount := 0
	details := make([]model.TransactionDetail, 0, len(items))

	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %d", item.ProductID)
		}

		var (
			productName  string
			productPrice int
		)

		// Check is product exist
		err = tx.QueryRow(`
			SELECT name, price
			FROM product
			WHERE id = $1
		`, item.ProductID).Scan(&productName, &productPrice)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Safe stock update
		res, err := tx.Exec(`
			UPDATE product
			SET stock = stock - $1
			WHERE id = $2 AND stock >= $1
		`, item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		rows, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}
		if rows == 0 {
			return nil, fmt.Errorf("stock for product %s is not enough", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// Insert transaction detail
		var detailID int
		err = stmtDetail.QueryRow(
			transactionID,
			item.ProductID,
			item.Quantity,
			subtotal,
		).Scan(&detailID)
		if err != nil {
			return nil, err
		}

		details = append(details, model.TransactionDetail{
			ID:            detailID,
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			ProductName:   productName,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	// Update total amount after all transaction
	_, err = tx.Exec(`
		UPDATE transaction
		SET total_amount = $1
		WHERE id = $2
	`, totalAmount, transactionID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		CreatedAt:   createdAt,
		Details:     details,
	}, nil
}
