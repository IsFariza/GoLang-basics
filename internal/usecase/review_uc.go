package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type ReviewUsecase struct {
	reviewRepo domain.ReviewRepo
}

func NewReviewUsecase(reviewRepo domain.ReviewRepo) *ReviewUsecase {
	return &ReviewUsecase{
		reviewRepo: reviewRepo,
	}
}
func (uc *ReviewUsecase) Create(ctx context.Context, review *domain.Review) error {
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

func (uc *ReviewUsecase) Update(ctx context.Context, id string, updatedReview *domain.Review) error {
	review, err := uc.GetById(ctx, id)
	if err != nil {
		return err
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
func (uc *ReviewUsecase) Delete(ctx context.Context, id string) error {
	return uc.reviewRepo.Delete(ctx, id)
}
