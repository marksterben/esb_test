package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	client *gorm.DB
}

func NewRepo(client *gorm.DB) *Repository {
	return &Repository{client}
}
