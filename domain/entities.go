package domain

import "time"

type CustomerEntity struct {
	ID        int       `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type ItemEntity struct {
	ID        int       `gorm:"column:id" json:"id"`
	Type      string    `gorm:"column:type" json:"type"`
	Desc      string    `gorm:"column:desc" json:"desc"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type InvoiceEntity struct {
	ID         string                `gorm:"column:id" json:"invoice_id"`
	Subject    string                `gorm:"column:subject" json:"subject"`
	IssueDate  time.Time             `gorm:"column:issue_date" json:"issue_date"`
	DueDate    time.Time             `gorm:"column:due_date" json:"due_date"`
	TotalItem  int                   `gorm:"column:total_item" json:"total_item"`
	Subtotal   int                   `gorm:"column:subtotal" json:"subtotal"`
	Tax        int                   `gorm:"column:tax" json:"tax"`
	Grandtotal int                   `gorm:"column:grandtotal" json:"grandtotal"`
	Payment    int                   `gorm:"column:payment;default:0" json:"payment"`
	Status     string                `gorm:"column:status;default:unpaid" json:"status"`
	CustomerID int                   `gorm:"column:customer_id" json:"customer_id"`
	CreatedAt  time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time             `gorm:"column:updated_at" json:"updated_at"`
	Customer   CustomerEntity        `gorm:"foreignKey:CustomerID"`
	Details    []InvoiceDetailEntity `gorm:"foreignKey:InvoiceID"`
}

type InvoiceDetailEntity struct {
	ID        int       `gorm:"column:id" json:"id"`
	InvoiceID string    `gorm:"column:invoice_id" json:"invoice_id"`
	ItemID    int       `gorm:"column:item_id" json:"item_id"`
	Qty       int       `gorm:"column:qty" json:"qty"`
	Price     int       `gorm:"column:price" json:"price"`
	Amount    int       `gorm:"column:amount" json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CustomerEntity) TableName() string {
	return "customers"
}

func (ItemEntity) TableName() string {
	return "items"
}

func (InvoiceEntity) TableName() string {
	return "invoices"
}

func (InvoiceDetailEntity) TableName() string {
	return "invoice_details"
}
