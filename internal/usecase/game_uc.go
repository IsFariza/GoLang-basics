package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type GameUseCase struct {
	repo          domain.GameRepo
	companyRepo   domain.CompanyRepo
	emulationRepo domain.EmulationRepo
}

func NewGameUseCase(repo domain.GameRepo, companyRepo domain.CompanyRepo, emulationRepo domain.EmulationRepo) *GameUseCase {
	return &GameUseCase{
		repo:          repo,
		companyRepo:   companyRepo,
		emulationRepo: emulationRepo,
	}
}

func (uc *GameUseCase) Create(ctx context.Context, game *domain.Game) error {
	if _, err := uc.companyRepo.GetById(ctx, game.PublisherId.Hex()); err != nil {
		return domain.ErrorInvalidPublisher
	}

	if _, err := uc.companyRepo.GetById(ctx, game.DeveloperId.Hex()); err != nil {
		return domain.ErrorInvalidDeveloper
	}

	if !game.EmulationId.IsZero() {
		if _, err := uc.emulationRepo.GetById(ctx, game.EmulationId.Hex()); err != nil {
			return domain.ErrorInvalidEmulator
		}
	}

	game.CreatedAt = time.Now()

	return uc.repo.Create(ctx, game)
}

func (uc *GameUseCase) GetAll(ctx context.Context) ([]*domain.Game, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *GameUseCase) GetById(ctx context.Context, id string) (*domain.Game, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *GameUseCase) Update(ctx context.Context, id string, updatedGame *domain.Game) error {
	existingGame, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if !updatedGame.PublisherId.IsZero() {
		_, err := uc.companyRepo.GetById(ctx, updatedGame.PublisherId.Hex())
		if err != nil {
			return domain.ErrorInvalidPublisher
		}

		existingGame.PublisherId = updatedGame.PublisherId
	}

	if !updatedGame.DeveloperId.IsZero() {
		_, err := uc.companyRepo.GetById(ctx, updatedGame.DeveloperId.Hex())
		if err != nil {
			return domain.ErrorInvalidDeveloper
		}

		existingGame.DeveloperId = updatedGame.DeveloperId
	}

	if updatedGame.Title != "" {
		existingGame.Title = updatedGame.Title
	}

	if updatedGame.Description != "" {
		existingGame.Description = updatedGame.Description
	}

	if updatedGame.ReleaseDate != nil {
		existingGame.ReleaseDate = updatedGame.ReleaseDate
	}

	if updatedGame.Price != nil {
		existingGame.Price = updatedGame.Price
	}

	now := time.Now()
	existingGame.UpdatedAt = &now
	existingGame.IsVerified = false

	return uc.repo.Update(ctx, id, existingGame)
}

func (uc *GameUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *GameUseCase) Approve(ctx context.Context, id string) error {
	return uc.repo.Approve(ctx, id)
}
