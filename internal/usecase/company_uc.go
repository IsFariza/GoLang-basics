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

func (uc *CompanyUsecase) Create(ctx context.Context, company *domain.Company) error {
	company.CreatedAt = time.Now()
	return uc.companyRepo.Create(ctx, company)

}

func (uc *CompanyUsecase) GetAll(ctx context.Context) ([]*domain.Company, error) {
	return uc.companyRepo.GetAll(ctx)

}

func (uc *CompanyUsecase) GetById(ctx context.Context, id string) (*domain.Company, error) {
	return uc.companyRepo.GetById(ctx, id)
}

func (uc *CompanyUsecase) Update(ctx context.Context, id string, updatedCompany *domain.Company) error {
	existingCompany, err := uc.companyRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if updatedCompany.Name != "" {
		existingCompany.Name = updatedCompany.Name
	}

	if updatedCompany.Description != "" {
		existingCompany.Description = updatedCompany.Description
	}

	if updatedCompany.Country != "" {
		existingCompany.Country = updatedCompany.Country
	}

	if updatedCompany.Contacts.Email != "" {
		existingCompany.Contacts.Email = updatedCompany.Contacts.Email
	}

	if updatedCompany.Contacts.Phone != "" {
		existingCompany.Contacts.Phone = updatedCompany.Contacts.Phone
	}

	now := time.Now()
	existingCompany.UpdatedAt = &now

	return uc.companyRepo.Update(ctx, id, existingCompany)
}

func (uc *CompanyUsecase) Delete(ctx context.Context, id string) error {
	return uc.companyRepo.Delete(ctx, id)
}
