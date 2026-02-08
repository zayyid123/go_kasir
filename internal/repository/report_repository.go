package repository

import (
	"database/sql"
	"fmt"
	"kasir-api/internal/model"
	"time"
)

type ReportRepository interface {
	Report(startDate, endDate *time.Time) (*model.Report, error)
	BestSellingProduct(startDate, endDate *time.Time) (*model.BestSeller, error)
}

type reportRepo struct {
	db *sql.DB
}

func NewReportRepo(db *sql.DB) ReportRepository {
	return &reportRepo{db: db}
}

func (r *reportRepo) Report(startDate, endDate *time.Time) (*model.Report, error) {
	var (
		totalRevenue     int
		totalTransaction int
	)

	query := `
		SELECT
			COALESCE(SUM(total_amount), 0) AS total_revenue,
			COUNT(*) AS total_transaction
		FROM transaction
		WHERE 1=1
	`
	args := []any{}
	argPos := 1

	if startDate != nil {
		query += fmt.Sprintf(" AND created_at >= $%d", argPos)
		args = append(args, *startDate)
		argPos++
	}

	if endDate != nil {
		query += fmt.Sprintf(" AND created_at < $%d", argPos)
		args = append(args, *endDate)
		argPos++
	}

	// default: hari ini
	if startDate == nil && endDate == nil {
		query += `
			AND created_at >= CURRENT_DATE
			AND created_at < CURRENT_DATE + INTERVAL '1 day'
		`
	}

	err := r.db.QueryRow(query, args...).Scan(
		&totalRevenue,
		&totalTransaction,
	)
	if err != nil {
		return nil, err
	}

	return &model.Report{
		TotalRevenue:     totalRevenue,
		TotalTransaction: totalTransaction,
	}, nil
}

func (r *reportRepo) BestSellingProduct(
	startDate, endDate *time.Time,
) (*model.BestSeller, error) {

	var result model.BestSeller

	query := `
		SELECT
			td.product_id,
			p.name,
			SUM(td.quantity) AS total_qty
		FROM transaction_detail td
		JOIN transaction t ON t.id = td.transaction_id
		JOIN product p ON p.id = td.product_id
		WHERE 1=1
	`
	args := []any{}
	arg := 1

	if startDate != nil {
		query += fmt.Sprintf(" AND t.created_at >= $%d", arg)
		args = append(args, *startDate)
		arg++
	}

	if endDate != nil {
		query += fmt.Sprintf(" AND t.created_at < $%d", arg)
		args = append(args, *endDate)
		arg++
	}

	// default: hari ini
	if startDate == nil && endDate == nil {
		query += `
			AND t.created_at >= CURRENT_DATE
			AND t.created_at < CURRENT_DATE + INTERVAL '1 day'
		`
	}

	query += `
		GROUP BY td.product_id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`

	row := r.db.QueryRow(query, args...)

	err := row.Scan(
		&result.ProductID,
		&result.Name,
		&result.Sold,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
