package repository

import (
	"esb/domain"
	"time"

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

func (repo *Repository) GetInvoices(request domain.InvoiceRequest) (*[]domain.InvoiceEntity, error) {
	var invoices []domain.InvoiceEntity

	query := repo.client.Offset(request.Page * request.Limit).Limit(request.Limit).Joins("Customer")

	if len(request.InvoiceID) > 0 {
		query.Where("invoices.id like ?", ("%" + request.InvoiceID + "%"))
	}

	if len(request.Subject) > 0 {
		query.Where("subject like ?", ("%" + request.Subject + "%"))
	}

	if request.TotalItem != nil {
		query.Where("total_item = ?", *request.TotalItem)
	}

	if len(request.Customer) > 0 {
		query.Where("Customer.name like ?", ("%" + request.Customer + "%"))
	}

	if len(request.Status) > 0 {
		query.Where("status = ?", request.Status)
	}

	if len(request.IssueDate) > 0 {
		query.Where("issue_date = ?", request.IssueDate)
	}

	if len(request.DueDate) > 0 {
		query.Where("due_date = ?", request.DueDate)
	}

	result := query.Find(&invoices)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &invoices, nil
}

func (repo *Repository) FindOneInvoice(request string) (*domain.InvoiceEntity, error) {
	var invoice domain.InvoiceEntity

	result := repo.client.Preload("Customer").Preload("Details").Where("invoices.id = ?", request).Find(&invoice)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (repo *Repository) CreateInvoice(request domain.PostInvoiceRequest) (*domain.InvoiceEntity, error) {
	dateFormat := "2006-01-02"
	issueDate, err := time.Parse(dateFormat, request.IssueDate)
	if err != nil {
		return nil, err
	}
	dueDate, err := time.Parse(dateFormat, request.DueDate)
	if err != nil {
		return nil, err
	}

	var details []domain.InvoiceDetailEntity
	for _, elm := range request.Details {
		details = append(details, domain.InvoiceDetailEntity{
			ItemID: elm.ItemID,
			Qty:    elm.Qty,
			Price:  elm.Price,
			Amount: elm.Amount,
		})
	}

	invoice := &domain.InvoiceEntity{
		ID:         request.ID,
		Subject:    request.Subject,
		IssueDate:  issueDate,
		DueDate:    dueDate,
		TotalItem:  request.TotalItem,
		Subtotal:   request.Subtotal,
		Tax:        request.Tax,
		Grandtotal: request.Grandtotal,
		CustomerID: request.CustomerID,
		Details:    &details,
	}

	if err := repo.client.Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

func (repo *Repository) GetMaxInvoiceID() (*string, error) {
	var maxID string
	err := repo.client.Table("invoices").Select("COALESCE(MAX(id), '0000')").Scan(&maxID).Error
	if err != nil {
		return nil, err
	}

	return &maxID, nil
}
