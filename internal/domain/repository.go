package domain

import (
	"context"
)

type CompanyRepo interface {
	Create(ctx context.Context, company Company) error
	GetAll(ctx context.Context) ([]Company, error)
	GetById(ctx context.Context, id string) (Company, error)
	Update(ctx context.Context, id string, updates Company) error
	Delete(ctx context.Context, id string) error
}
