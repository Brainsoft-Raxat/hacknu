package service

import (
	"context"
	"hacknu/internal/repository"
	"hacknu/pkg/data"
	"time"
)

type someService struct {
	someRepo *repository.Repository
}

func (s *someService) DoSomething(ctx context.Context, req data.DoSomethingRequest) (resp data.DoSomethingResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	return
}

func NewSomeService(repo *repository.Repository) SomeService {
	return &someService{
		someRepo: repo,
	}
}
