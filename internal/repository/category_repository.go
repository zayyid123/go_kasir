package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetByID(id int) (model.Category, error)
	Create(c model.CreateCategoryRequest) (model.Category, error)
	Update(id int, c model.UpdateCategoryRequest) (model.Category, error)
	Delete(id int) error
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query(`select id, name, description, created_at from public.category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)

	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *categoryRepo) GetByID(id int) (model.Category, error) {
	var c model.Category
	err := r.db.QueryRow(`
		select id, name, description, created_at
		from category where id = $1
	`, id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt)

	return c, err
}

func (r *categoryRepo) Create(c model.CreateCategoryRequest) (model.Category, error) {
	var result model.Category

	err := r.db.QueryRow(`
		insert into category(name, description)
		values ($1,$2)
		returning id, name, description, created_at
	`, c.Name, c.Description).
		Scan(&result.ID, &result.Name, &result.Description, &result.CreatedAt)

	return result, err
}

func (r *categoryRepo) Update(id int, c model.UpdateCategoryRequest) (model.Category, error) {
	var result model.Category

	err := r.db.QueryRow(`
		update category
		set name=$1, description=$2
		where id=$3
		returning id, name, description, created_at
	`, c.Name, c.Description, id).
		Scan(&result.ID, &result.Name, &result.Description, &result.CreatedAt)

	return result, err
}

func (r *categoryRepo) Delete(id int) error {
	_, err := r.db.Exec(`delete from category where id=$1`, id)
	return err
}
