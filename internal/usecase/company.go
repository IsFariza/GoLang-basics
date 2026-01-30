package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type CompanyUsecase struct {
	companyRepo domain.CompanyRepo
}

func NewCompanyUsecase(companyRepo domain.CompanyRepo) *CompanyUsecase {
	return &CompanyUsecase{
		companyRepo: companyRepo,
	}
}

func (uc *CompanyUsecase) Create(ctx context.Context, company domain.Company) error {
	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()
	return uc.companyRepo.Create(ctx, company)

}

func (uc *CompanyUsecase) GetAll(ctx context.Context) ([]domain.Company, error) {
	return uc.companyRepo.GetAll(ctx)

}

func (uc *CompanyUsecase) GetById(ctx context.Context, id string) (domain.Company, error) {
	return uc.companyRepo.GetById(ctx, id)
}

func (uc *CompanyUsecase) Update(ctx context.Context, id string, updates domain.Company) error {
	updates.UpdatedAt = time.Now()
	return uc.companyRepo.Update(ctx, id, updates)

}

func (uc *CompanyUsecase) Delete(ctx context.Context, id string) error {
	return uc.companyRepo.Delete(ctx, id)
}
