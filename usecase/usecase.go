package usecase

import (
	"context"
	"esb/domain"
	"time"
)

var (
	default_size = 12
)

type Repository interface {
	GetCustomers() (*[]domain.CustomerEntity, error)
	GetItems() (*[]domain.ItemEntity, error)
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

func (u *Usecase) GetCustomers(ctx context.Context) (*domain.MultipleResponse[domain.CustomerEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	resp, err := u.Repo.GetCustomers()
	if err != nil {
		return nil, err
	}

	return &domain.MultipleResponse[domain.CustomerEntity]{
		Data: *resp,
	}, nil
}

func (u *Usecase) GetItems(ctx context.Context) (*domain.MultipleResponse[domain.ItemEntity], error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	resp, err := u.Repo.GetItems()
	if err != nil {
		return nil, err
	}

	return &domain.MultipleResponse[domain.ItemEntity]{
		Data: *resp,
	}, nil
}
