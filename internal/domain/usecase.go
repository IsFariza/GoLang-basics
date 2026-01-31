package domain

import "context"

type GameUC interface {
	Create(ctx context.Context, game *Game) error
	GetAll(ctx context.Context) ([]*Game, error)
	GetById(ctx context.Context, id string) (*Game, error)
	Update(ctx context.Context, id string, updates *Game) error
	Delete(ctx context.Context, id string) error
}

type CompanyUC interface {
	Create(ctx context.Context, company *Company) error
	GetAll(ctx context.Context) ([]*Company, error)
	GetById(ctx context.Context, id string) (*Company, error)
	Update(ctx context.Context, id string, updates *Company) error
	Delete(ctx context.Context, id string) error
}

type EmulationUC interface {
	Create(ctx context.Context, emulation *Emulation) error
	GetAll(ctx context.Context) ([]*Emulation, error)
	GetById(ctx context.Context, id string) (*Emulation, error)
	Update(ctx context.Context, id string, updates *Emulation) error
	Delete(ctx context.Context, id string) error
}

type UserUC interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, updates *User) error
	Delete(ctx context.Context, id string) error
}
