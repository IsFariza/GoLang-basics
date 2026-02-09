package domain

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/delivery/dto"
)

type GameRepo interface {
	Create(ctx context.Context, game *Game, id string) error
	GetAll(ctx context.Context) ([]*Game, error)
	GetAllVerified(ctx context.Context) ([]*Game, error)
	GetById(ctx context.Context, id string) (*Game, error)
	GetByIds(ctx context.Context, ids []string) ([]Game, error)
	GetByUserId(ctx context.Context, userId string) ([]*Game, error)
	Update(ctx context.Context, id string, updates *Game) error
	Delete(ctx context.Context, id string) error
	Verify(ctx context.Context, id string) error
	Unverify(ctx context.Context, id string) error
	SearchByTitle(ctx context.Context, title string) ([]*Game, error)
	GetStats(ctx context.Context) (*dto.GameStatsDTO, error)
}

type CompanyRepo interface {
	Create(ctx context.Context, company *Company) error
	GetAll(ctx context.Context) ([]*Company, error)
	GetById(ctx context.Context, id string) (*Company, error)
	GetVerified(ctx context.Context) ([]*Company, error)
	Verify(ctx context.Context, id string) error
	Unverify(ctx context.Context, id string) error
	Update(ctx context.Context, id string, updates *Company) error
	Delete(ctx context.Context, id string) error
}

type EmulationRepo interface {
	Create(ctx context.Context, emulation *Emulation) error
	GetAll(ctx context.Context) ([]*Emulation, error)
	GetById(ctx context.Context, id string) (*Emulation, error)
	Update(ctx context.Context, id string, updates *Emulation) error
	Delete(ctx context.Context, id string) error
}

type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, updates *User) error
	Delete(ctx context.Context, id string) error
	ChangeRoleToModerator(ctx context.Context, id string) error
	ChangeRoleToUser(ctx context.Context, id string) error
}

type ReviewRepo interface {
	Create(ctx context.Context, review *Review) error
	GetAll(ctx context.Context) ([]*Review, error)
	GetById(ctx context.Context, id string) (*Review, error)
	GetReviewsByGameId(ctx context.Context, gameId string) ([]*Review, error)
	Update(ctx context.Context, id string, updates *Review) error
	Delete(ctx context.Context, id string, userId string, userRole string) error
}

type PurchaseRepo interface {
	Create(ctx context.Context, purchase *Purchase) error
	GetAll(ctx context.Context) ([]*Purchase, error)
	GetById(ctx context.Context, id string) (*Purchase, error)
	Delete(ctx context.Context, id string) error
}
