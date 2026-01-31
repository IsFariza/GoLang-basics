package domain

import "context"

type GameUC interface {
	Create(ctx context.Context, game *Game) error
	GetAll(ctx context.Context) ([]*Game, error)
	GetById(ctx context.Context, id string) (*Game, error)
	Update(ctx context.Context, id string, updates *Game) error
	Delete(ctx context.Context, id string) error
}
