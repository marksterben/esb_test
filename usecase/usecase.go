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
