package usecase

import (
	"time"
)

var (
	default_size = 12
)

type Repository interface {
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
