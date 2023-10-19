package usecase

import (
	"context"
	"esb/domain"
	"esb/helper"
	"time"
)

var (
	default_size = 12
)

type Repository interface {
	GetCustomers() (*[]domain.CustomerEntity, error)
	GetItems() (*[]domain.ItemEntity, error)
	GetInvoices(request domain.InvoiceRequest) (*[]domain.InvoiceEntity, error)
	FindOneInvoice(request string) (*domain.InvoiceEntity, error)
	CreateInvoice(request domain.PostInvoiceRequest) (*domain.InvoiceEntity, error)
	GetMaxInvoiceID() (*string, error)
}

type Usecase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewUsecase(repo Repository, time time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		ContextTimeout: time,
	}
}

func (u *Usecase) GetCustomers(ctx context.Context) (*domain.ApiResponse[[]domain.CustomerEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	resp, err := u.Repo.GetCustomers()
	if err != nil {
		return nil, err
	}

	return &domain.ApiResponse[[]domain.CustomerEntity]{
		Data: resp,
	}, nil
}

func (u *Usecase) GetItems(ctx context.Context) (*domain.ApiResponse[[]domain.ItemEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	resp, err := u.Repo.GetItems()
	if err != nil {
		return nil, err
	}

	return &domain.ApiResponse[[]domain.ItemEntity]{
		Data: resp,
	}, nil
}

func (u *Usecase) GetInvoices(ctx context.Context, request domain.InvoiceRequest) (*domain.ApiResponse[[]domain.InvoiceEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if request.Limit < 1 || request.Limit > 100 {
		request.Limit = default_size
	}

	if request.Page < 0 {
		request.Page = 0
	}

	resp, err := u.Repo.GetInvoices(request)
	if err != nil {
		return nil, err
	}

	return &domain.ApiResponse[[]domain.InvoiceEntity]{
		Data: resp,
	}, nil
}

func (u *Usecase) FindOneInvoice(ctx context.Context, request string) (*domain.ApiResponse[domain.InvoiceEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	resp, err := u.Repo.FindOneInvoice(request)
	if err != nil {
		return nil, err
	}

	return &domain.ApiResponse[domain.InvoiceEntity]{
		Data: resp,
	}, nil
}

func (u *Usecase) CreateInvoice(ctx context.Context, request domain.PostInvoiceRequest) (*domain.ApiResponse[domain.InvoiceEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	stringID, err := u.Repo.GetMaxInvoiceID()
	if err != nil {
		return nil, err
	}

	stringID, err = helper.IncrementNumberString(*stringID)
	if err != nil {
		return nil, err
	}

	request.ID = *stringID

	resp, err := u.Repo.CreateInvoice(request)
	if err != nil {
		return nil, err
	}

	return &domain.ApiResponse[domain.InvoiceEntity]{
		Data:    resp,
		Message: "Invoice successfully created",
	}, nil
}
