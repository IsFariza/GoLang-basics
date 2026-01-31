package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type GameUseCase struct {
	repo domain.GameRepo
}

func NewGameUseCase(repo domain.GameRepo) *GameUseCase {
	return &GameUseCase{
		repo: repo,
	}
}

func (uc *GameUseCase) CreateGame(ctx context.Context, game domain.Game) error {
	//TODO: some buisness logic, maybe validation

	return uc.repo.Create(ctx, game)
}

func (uc *GameUseCase) GetAll(ctx context.Context) ([]domain.Game, error) {
	//TODO: some buisness logic, maybe validation

	return uc.repo.GetAll(ctx)
}

func (uc *GameUseCase) GetById(ctx context.Context, id string) (domain.Game, error) {
	//TODO: some buisness logic, maybe validation

	return uc.repo.GetById(ctx, id)
}

func (uc *GameUseCase) Update(ctx context.Context, id string, updatedGame domain.Game) error {
	existingGame, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if updatedGame.Title != "" {
		existingGame.Title = updatedGame.Title
	}

	if updatedGame.Description != "" {
		existingGame.Description = updatedGame.Description
	}

	if !updatedGame.ReleaseDate.IsZero() {
		existingGame.ReleaseDate = updatedGame.ReleaseDate
	}

	if updatedGame.Price != 0 {
		existingGame.Price = updatedGame.Price
	}

	existingGame.UpdatedAt = time.Now()
	existingGame.IsVerified = false

	return uc.repo.Update(ctx, id, existingGame)
}

func (uc *GameUseCase) Delete(ctx context.Context, id string) error {
	//TODO: some buisness logic, maybe validation

	return uc.repo.Delete(ctx, id)
}
