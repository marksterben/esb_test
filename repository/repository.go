package repository

import (
	"esb/domain"

	"gorm.io/gorm"
)

type Repository struct {
	client *gorm.DB
}

func NewRepo(client *gorm.DB) *Repository {
	return &Repository{client}
}

func (repo *Repository) GetCustomers() (*[]domain.CustomerEntity, error) {
	var customers []domain.CustomerEntity

	result := repo.client.Find(&customers)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &customers, nil
}

func (repo *Repository) GetItems() (*[]domain.ItemEntity, error) {
	var items []domain.ItemEntity

	result := repo.client.Find(&items)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &items, nil
}
