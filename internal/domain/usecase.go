package domain

import "context"

type GameUC interface {
	Create(ctx context.Context, game *Game) error
	GetAll(ctx context.Context) ([]*Game, error)
	GetById(ctx context.Context, id string) (*PopulatedGame, error)
	GetReviewsByGameId(ctx context.Context, id string) ([]*Review, error)
	Update(ctx context.Context, id string, updates *Game) error
	Delete(ctx context.Context, id string) error
	Approve(ctx context.Context, id string) error
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
	SignUp(ctx context.Context, user *User, password string) error
	Login(ctx context.Context, username, password string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, updates *User) error
	Delete(ctx context.Context, id string) error
}

type ReviewUC interface {
	Create(ctx context.Context, reviews *Review) error
	GetAll(ctx context.Context) ([]*Review, error)
	GetById(ctx context.Context, id string) (*Review, error)
	Update(ctx context.Context, id string, updates *Review) error
	Delete(ctx context.Context, id string) error
}

type PurchaseUC interface {
	Create(ctx context.Context, purchase *Purchase) error
	GetAll(ctx context.Context) ([]*Purchase, error)
	GetById(ctx context.Context, id string) (*Purchase, error)
	Delete(ctx context.Context, id string) error
}
