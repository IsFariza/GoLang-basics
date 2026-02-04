package domain

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/delivery/dto"
)

type GameUC interface {
	Create(ctx context.Context, game *Game, id string) error
	GetAll(ctx context.Context) ([]*Game, error)
	GetAllVerified(ctx context.Context) ([]*Game, error)
	GetById(ctx context.Context, id string) (*PopulatedGame, error)
	GetByUserId(ctx context.Context, userId string) ([]*Game, error)
	GetReviewsByGameId(ctx context.Context, id string) ([]*Review, error)
	Update(ctx context.Context, id string, updates *Game, userId, userRole string) error
	Delete(ctx context.Context, id string) error
	VerifySwitch(ctx context.Context, id string) error
	SearchByTitle(ctx context.Context, title string) ([]*Game, error)
	GetUserLibraryWithDetails(ctx context.Context, userId string) ([]dto.UserLibraryItemDTO, error)
	GetStats(ctx context.Context) (*dto.GameStatsDTO, error)
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
