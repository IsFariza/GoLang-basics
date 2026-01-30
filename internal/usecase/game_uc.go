package usecase

import (
	"context"

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

func (s *GameUseCase) CreateGame(ctx context.Context, game domain.Game) error {
	//TODO: some buisness logic, maybe validation

	return s.repo.Create(ctx, game)
}

func (s *GameUseCase) GetAll(ctx context.Context) ([]domain.Game, error) {
	//TODO: some buisness logic, maybe validation

	return s.repo.GetAll(ctx)
}

func (s *GameUseCase) GetById(ctx context.Context, id string) (domain.Game, error) {
	//TODO: some buisness logic, maybe validation

	return s.repo.GetById(ctx, id)
}

func (s *GameUseCase) Update(ctx context.Context, id string, updates domain.Game) error {
	//TODO: some buisness logic, maybe validation

	return s.repo.Update(ctx, id, updates)
}

func (s *GameUseCase) Delete(ctx context.Context, id string) error {
	//TODO: some buisness logic, maybe validation

	return s.repo.Delete(ctx, id)
}
