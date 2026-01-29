package model

type Product struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CreatedAt  string `json:"created_at"`
	CategoryId int    `json:"category_id"`
}

type ProductWithCategory struct {
	Product
	Category *Category `json:"category"`
}

type CreateProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Stock      int     `json:"stock"`
	CategoryId int     `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
