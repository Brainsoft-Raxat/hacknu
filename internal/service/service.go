package service

import (
	"context"
	"hacknu/internal/repository"
	"hacknu/pkg/data"
)

type SomeService interface {
	DoSomething(ctx context.Context, request data.DoSomethingRequest) (response data.DoSomethingResponse, err error)
}

type Service struct {
	SomeService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		SomeService: NewSomeService(repos),
	}
}
