package domain

import "time"

type CustomerEntity struct {
	ID        int       `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CustomerEntity) TableName() string {
	return "customers"
}
