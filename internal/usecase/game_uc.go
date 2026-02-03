package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"golang.org/x/sync/errgroup"
)

type GameUseCase struct {
	repo          domain.GameRepo
	companyRepo   domain.CompanyRepo
	emulationRepo domain.EmulationRepo
	reviewRepo    domain.ReviewRepo
	userRepo      domain.UserRepo
}

func NewGameUseCase(repo domain.GameRepo, companyRepo domain.CompanyRepo, emulationRepo domain.EmulationRepo, reviewRepo domain.ReviewRepo, userRepo domain.UserRepo) *GameUseCase {
	return &GameUseCase{
		repo:          repo,
		companyRepo:   companyRepo,
		emulationRepo: emulationRepo,
		reviewRepo:    reviewRepo,
		userRepo:      userRepo,
	}
}

func (uc *GameUseCase) Create(ctx context.Context, game *domain.Game, userId string) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		_, err = uc.companyRepo.GetById(ctx, game.PublisherId.Hex())
		if err != nil {
			return domain.ErrorInvalidPublisher
		}
		return nil
	})

	g.Go(func() error {
		var err error
		_, err = uc.companyRepo.GetById(ctx, game.DeveloperId.Hex())
		if err != nil {
			return domain.ErrorInvalidDeveloper
		}
		return nil
	})

	if !game.EmulationId.IsZero() {
		g.Go(func() error {
			var err error
			_, err = uc.emulationRepo.GetById(ctx, game.EmulationId.Hex())
			if err != nil {
				return domain.ErrorInvalidEmulator
			}
			return nil
		})
	}

	game.CreatedAt = time.Now()

	return uc.repo.Create(ctx, game, userId)
}

func (uc *GameUseCase) GetAll(ctx context.Context) ([]*domain.Game, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *GameUseCase) GetAllVerified(ctx context.Context) ([]*domain.Game, error) {
	return uc.repo.GetAllVerified(ctx)
}

func (uc *GameUseCase) GetById(ctx context.Context, id string) (*domain.PopulatedGame, error) {
	game, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var publisher *domain.Company
	var developer *domain.Company
	var emulation *domain.Emulation

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		publisher, err = uc.companyRepo.GetById(ctx, game.PublisherId.Hex())
		return err
	})

	g.Go(func() error {
		var err error
		developer, err = uc.companyRepo.GetById(ctx, game.DeveloperId.Hex())
		return err
	})

	// Ignore err because emulation is optional
	g.Go(func() error {
		emulation, _ = uc.emulationRepo.GetById(ctx, game.EmulationId.Hex())
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &domain.PopulatedGame{
		ID:             game.ID,
		Publisher:      publisher,
		Developer:      developer,
		Emulation:      emulation,
		OriginalSystem: game.OriginalSystem,
		Title:          game.Title,
		Description:    game.Description,
		ReleaseDate:    game.ReleaseDate,
		Price:          game.Price,
		IsVerified:     game.IsVerified,
		Category:       game.Category,
		CreatedAt:      game.CreatedAt,
		UpdatedAt:      game.UpdatedAt,
	}, nil
}

func (uc *GameUseCase) GetByUserId(ctx context.Context, userId string) ([]*domain.Game, error) {
	return uc.repo.GetByUserId(ctx, userId)
}

func (uc *GameUseCase) GetReviewsByGameId(ctx context.Context, gameId string) ([]*domain.Review, error) {
	_, err := uc.repo.GetById(ctx, gameId)
	if err != nil {
		return nil, err
	}

	return uc.reviewRepo.GetReviewsByGameId(ctx, gameId)
}

func (uc *GameUseCase) Update(ctx context.Context, id string, updatedGame *domain.Game, userId, userRole string) error {
	existingGame, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if userRole != "admin" && existingGame.UserId.Hex() != userId {
		return errors.New("permission denied: you are not the owner of this game")
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

	if !updatedGame.EmulationId.IsZero() {
		_, err := uc.emulationRepo.GetById(ctx, updatedGame.EmulationId.Hex())
		if err != nil {
			return domain.ErrorInvalidEmulator
		}

		existingGame.EmulationId = updatedGame.EmulationId
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

func (uc *GameUseCase) VerifySwitch(ctx context.Context, id string) error {
	game, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if game.IsVerified {
		return uc.repo.Unverify(ctx, id)
	}

	return uc.repo.Verify(ctx, id)
}

func (uc *GameUseCase) SearchByTitle(ctx context.Context, title string) ([]*domain.Game, error) {
	games, err := uc.repo.SearchByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	return games, nil
}
