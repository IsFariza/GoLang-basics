package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type ReviewUsecase struct {
	reviewRepo domain.ReviewRepo
	userRepo   domain.UserRepo
}

func NewReviewUsecase(reviewRepo domain.ReviewRepo, userRepo domain.UserRepo) *ReviewUsecase {
	return &ReviewUsecase{
		reviewRepo: reviewRepo,
		userRepo:   userRepo,
	}
}
func (uc *ReviewUsecase) Create(ctx context.Context, review *domain.Review) error {
	user, err := uc.userRepo.GetById(ctx, review.UserId.Hex())
	if err != nil {
		return err
	}
	review.Username = user.Username
	ownsGame := false
	for _, libGame := range user.Library {
		if libGame.GameId == review.GameId {
			ownsGame = true
			break
		}
	}
	if !ownsGame {
		return errors.New("review denied: you must own the game to leave a review")
	}
	now := time.Now()
	review.CreatedAt = now
	review.UpdatedAt = &now
	return uc.reviewRepo.Create(ctx, review)
}
func (uc *ReviewUsecase) GetAll(ctx context.Context) ([]*domain.Review, error) {
	return uc.reviewRepo.GetAll(ctx)
}

func (uc *ReviewUsecase) GetById(ctx context.Context, id string) (*domain.Review, error) {
	return uc.reviewRepo.GetById(ctx, id)
}

func (uc *ReviewUsecase) Update(ctx context.Context, id string, currentUserID string, updatedReview *domain.Review) error {
	review, err := uc.GetById(ctx, id)
	if err != nil {
		return err
	}
	if review.UserId.Hex() != currentUserID {
		return errors.New("permission denied: you can only edit your own reviews")
	}
	if updatedReview.Rating != nil {
		review.Rating = updatedReview.Rating
	}
	if updatedReview.Content != "" {
		review.Content = updatedReview.Content
	}
	now := time.Now()
	review.UpdatedAt = &now
	return uc.reviewRepo.Update(ctx, id, review)
}
func (uc *ReviewUsecase) Delete(ctx context.Context, id string, userId string, userRole string) error {
	review, err := uc.reviewRepo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if userRole != "admin" && review.UserId.Hex() != userId {
		return errors.New("permission denied: you can only delete your own reviews")
	}

	return uc.reviewRepo.Delete(ctx, id, userId, userRole)
}
