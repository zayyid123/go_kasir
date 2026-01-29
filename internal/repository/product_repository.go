package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (model.Product, error)
	Create(p model.CreateProductRequest) (model.Product, error)
	Update(id int, p model.UpdateProductRequest) (model.Product, error)
	Delete(id int) error
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) GetAll() ([]model.Product, error) {
	rows, err := r.db.Query(`select id, name, price, stock, created_at from public.product`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepo) GetByID(id int) (model.Product, error) {
	var p model.Product
	err := r.db.QueryRow(`
		select id, name, price, stock, created_at
		from product where id = $1
	`, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)

	return p, err
}

func (r *productRepo) Create(p model.CreateProductRequest) (model.Product, error) {
	var result model.Product

	err := r.db.QueryRow(`
		insert into product(name, price, stock)
		values ($1,$2,$3)
		returning id, name, price, stock, created_at
	`, p.Name, p.Price, p.Stock).
		Scan(&result.ID, &result.Name, &result.Price, &result.Stock, &result.CreatedAt)

	return result, err
}

func (r *productRepo) Update(id int, p model.UpdateProductRequest) (model.Product, error) {
	var result model.Product

	err := r.db.QueryRow(`
		update product
		set name=$1, price=$2, stock=$3
		where id=$4
		returning id, name, price, stock, created_at
	`, p.Name, p.Price, p.Stock, id).
		Scan(&result.ID, &result.Name, &result.Price, &result.Stock, &result.CreatedAt)

	return result, err
}

func (r *productRepo) Delete(id int) error {
	_, err := r.db.Exec(`delete from product where id=$1`, id)
	return err
}
