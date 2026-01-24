package service

import (
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/store"
)

func GetProducts() []model.Product {
	return store.Products
}

func AddProduct(p model.Product) model.Product {
	p.ID = store.ProductAutoID
	store.ProductAutoID++

	store.Products = append(store.Products, p)
	return p
}

func GetProductByID(id uint) (model.Product, error) {
	for _, p := range store.Products {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Product{}, errors.New("produk tidak ditemukan")
}

func UpdateProduct(id uint, updated model.Product) (model.Product, error) {
	for i, p := range store.Products {
		if p.ID == id {
			updated.ID = id
			store.Products[i] = updated
			return updated, nil
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
