package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(c model.CreateCategoryRequest) (model.Category, error) {
	return s.repo.Create(c)
}

func (s *CategoryService) GetByID(id int) (model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(id int, c model.UpdateCategoryRequest) (model.Category, error) {
	return s.repo.Update(id, c)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
