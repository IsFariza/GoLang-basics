package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type UserUseCase struct {
	repo domain.UserRepo
}

func NewUserUseCase(repo domain.UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) Create(ctx context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()

	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) GetAll(ctx context.Context) ([]*domain.User, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UserUseCase) GetById(ctx context.Context, id string) (*domain.User, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *UserUseCase) Update(ctx context.Context, id string, updatedUser *domain.User) error {
	existingUser, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if updatedUser.Username != "" {
		existingUser.Username = updatedUser.Username
	}

	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}

	now := time.Now()
	existingUser.UpdatedAt = &now

	return uc.repo.Update(ctx, id, existingUser)
}

func (uc *UserUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
