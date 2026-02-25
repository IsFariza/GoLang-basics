package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo domain.UserRepo
}

func NewUserUseCase(repo domain.UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) SignUp(ctx context.Context, user *domain.User, password string) error {
	existingUser, err := uc.repo.GetByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("User with this email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = hash
	user.Role = "user"
	user.CreatedAt = time.Now()

	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
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

func (uc *UserUseCase) RoleSwitch(ctx context.Context, id string) error {
	user, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if user.Role != "moderator" {
		return uc.repo.ChangeRoleToModerator(ctx, id)
	}

	return uc.repo.ChangeRoleToUser(ctx, id)
}
