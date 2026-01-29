package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type ProductRepository interface {
	GetAll() ([]model.ProductWithCategory, error)
	GetByID(id int) (model.ProductWithCategory, error)
	Create(p model.CreateProductRequest) (model.ProductWithCategory, error)
	Update(id int, p model.UpdateProductRequest) (model.Product, error)
	Delete(id int) error
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) GetAll() ([]model.ProductWithCategory, error) {
	rows, err := r.db.Query(`
	select 
		p.id, p.name, p.price, p.stock, p.created_at, p.category_id,
		c.id, c.name, c.description, c.created_at
	from public.product p
	left join public.category c on c.id = p.category_id
	order by p.id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.ProductWithCategory, 0)

	for rows.Next() {
		var p model.ProductWithCategory
		var c model.Category

		err := rows.Scan(
			&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.CategoryId,
			&c.ID, &c.Name, &c.Description, &c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if c.ID != 0 {
			p.Category = &c
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *productRepo) GetByID(id int) (model.ProductWithCategory, error) {
	var p model.ProductWithCategory
	var cat model.Category
	err := r.db.QueryRow(`
	select 
		p.id, p.name, p.price, p.stock, p.created_at, p.category_id,
		c.id, c.name, c.description, c.created_at
	from public.product p
	left join public.category c on c.id = p.category_id
	where p.id = $1`, id).Scan(
		&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.CategoryId,
		&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt,
	)

	if cat.ID != 0 {
		p.Category = &cat
	}

	return p, err
}

func (r *productRepo) Create(p model.CreateProductRequest) (model.ProductWithCategory, error) {
	var result model.ProductWithCategory
	var cat model.Category

	err := r.db.QueryRow(`
	with inserted as (
		insert into public.product (name, price, stock, category_id)
		values ($1,$2,$3,$4)
		returning id, name, price, stock, created_at, category_id
	)
	select 
		i.id, i.name, i.price, i.stock, i.created_at, i.category_id,
		c.id, c.description, c.name, c.created_at
	from inserted i
	left join public.category c on c.id = i.category_id
`, p.Name, p.Price, p.Stock, p.CategoryId).
		Scan(
			&result.ID, &result.Name, &result.Price, &result.Stock, &result.CreatedAt,
			&result.CategoryId,
			&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt,
		)

	if cat.ID != 0 {
		result.Category = &cat
	}

	return result, err
}

func (r *productRepo) Update(id int, p model.UpdateProductRequest) (model.Product, error) {
	var result model.Product

	err := r.db.QueryRow(`
		update product
		set name=$1, price=$2, stock=$3
		where id=$4
		returning id, name, price, stock, created_at, category_id
	`, p.Name, p.Price, p.Stock, id).
		Scan(&result.ID, &result.Name, &result.Price, &result.Stock, &result.CreatedAt, &result.CategoryId)

	return result, err
}

func (r *productRepo) Delete(id int) error {
	_, err := r.db.Exec(`delete from product where id=$1`, id)
	return err
}
