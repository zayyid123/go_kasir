package service

import (
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/store"
)

func GetCategories() []model.Category {
	return store.Categories
}

func AddCategory(c model.CreateCategoryRequest) model.Category {
	category := model.Category{
		ID:          store.CategoryAutoID,
		Name:        c.Name,
		Description: c.Description,
	}
	store.CategoryAutoID++

	store.Categories = append(store.Categories, category)
	return category
}

func GetCategoryByID(id uint) (model.Category, error) {
	for _, c := range store.Categories {
		if c.ID == id {
			return c, nil
		}
	}
	return model.Category{}, errors.New("kategori tidak ditemukan")
}

func UpdateCategory(id uint, updated model.UpdateCategoryRequest) (model.Category, error) {
	for i, c := range store.Categories {
		if c.ID == id {
			updatedCategory := model.Category{
				ID:          id,
				Name:        updated.Name,
				Description: updated.Description,
			}
			store.Categories[i] = updatedCategory
			return updatedCategory, nil
		}
	}
	return model.Category{}, errors.New("kategori tidak ditemukan")
}

func DeleteCategory(id uint) error {
	for i, c := range store.Categories {
		if c.ID == id {
			store.Categories = append(store.Categories[:i], store.Categories[i+1:]...)
			return nil
		}
	}
	return errors.New("kategori tidak ditemukan")
}
