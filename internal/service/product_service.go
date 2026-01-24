package service

import (
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/store"
)

func GetProducts() []model.Product {
	return store.Products
}

func AddProduct(p model.CreateProductRequest) model.Product {
	product := model.Product{
		ID:    store.ProductAutoID,
		Name:  p.Name,
		Price: int(p.Price),
		Stock: p.Stock,
	}
	store.ProductAutoID++

	store.Products = append(store.Products, product)
	return product
}

func GetProductByID(id uint) (model.Product, error) {
	for _, p := range store.Products {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Product{}, errors.New("produk tidak ditemukan")
}

func UpdateProduct(id uint, updated model.UpdateProductRequest) (model.Product, error) {
	for i, p := range store.Products {
		if p.ID == id {
			updatedProduct := model.Product{
				ID:    id,
				Name:  updated.Name,
				Price: int(updated.Price),
				Stock: updated.Stock,
			}
			store.Products[i] = updatedProduct
			return updatedProduct, nil
		}
	}
	return model.Product{}, errors.New("produk tidak ditemukan")
}

func DeleteProduct(id uint) error {
	for i, p := range store.Products {
		if p.ID == id {
			store.Products = append(store.Products[:i], store.Products[i+1:]...)
			return nil
		}
	}
	return errors.New("produk tidak ditemukan")
}
