package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetAll() ([]model.ProductWithCategory, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id int) (model.ProductWithCategory, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(p model.CreateProductRequest) (model.ProductWithCategory, error) {
	return s.repo.Create(p)
}

func (s *ProductService) Update(id int, p model.UpdateProductRequest) (model.Product, error) {
	return s.repo.Update(id, p)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
