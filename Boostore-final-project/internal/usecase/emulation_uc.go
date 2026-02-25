package usecase

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type EmulationUsecase struct {
	emulationRepo domain.EmulationRepo
}

func NewEmulationUsecase(emulationRepo domain.EmulationRepo) *EmulationUsecase {
	return &EmulationUsecase{
		emulationRepo: emulationRepo,
	}
}

func (uc *EmulationUsecase) Create(ctx context.Context, emulation *domain.Emulation) error {
	return uc.emulationRepo.Create(ctx, emulation)

}

func (uc *EmulationUsecase) GetAll(ctx context.Context) ([]*domain.Emulation, error) {
	return uc.emulationRepo.GetAll(ctx)

}

func (uc *EmulationUsecase) GetById(ctx context.Context, id string) (*domain.Emulation, error) {
	return uc.emulationRepo.GetById(ctx, id)
}

func (uc *EmulationUsecase) Update(ctx context.Context, id string, updatedEmulation *domain.Emulation) error {

	existingCompany, err := uc.emulationRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if updatedEmulation.Name != "" {
		existingCompany.Name = updatedEmulation.Name
	}

	return uc.emulationRepo.Update(ctx, id, existingCompany)
}

func (uc *EmulationUsecase) Delete(ctx context.Context, id string) error {
	return uc.emulationRepo.Delete(ctx, id)
}
