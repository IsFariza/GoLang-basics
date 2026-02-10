package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type PurchaseUsecase struct {
	purchaseRepo domain.PurchaseRepo
	userRepo     domain.UserRepo
}

func NewPurchaseUsecase(purchaseRepo domain.PurchaseRepo, userRepo domain.UserRepo) *PurchaseUsecase {
	return &PurchaseUsecase{
		purchaseRepo: purchaseRepo,
		userRepo:     userRepo,
	}
}

func (uc *PurchaseUsecase) Create(ctx context.Context, purchase *domain.Purchase) error {
	err := uc.purchaseRepo.Create(ctx, purchase)
	if err != nil {
		return err
	}
	user, err := uc.userRepo.GetById(ctx, purchase.UserId)
	if err != nil {
		return err
	}
	now := time.Now()
	for _, item := range purchase.Items {
		isOwned := false
		for _, libGame := range user.Library {
			if libGame.GameId == item.GameId {
				isOwned = true
				break
			}
		}

		if !isOwned {
			user.Library = append(user.Library, domain.UserGame{
				GameId:        item.GameId,
				AddedAt:       now,
				PlaytimeHours: 0,
			})
		}
	}
	return uc.userRepo.Update(ctx, user.ID, user)
}
func (uc *PurchaseUsecase) GetAll(ctx context.Context) ([]*domain.Purchase, error) {
	return uc.purchaseRepo.GetAll(ctx)
}
func (uc *PurchaseUsecase) GetById(ctx context.Context, id string) (*domain.Purchase, error) {
	return uc.purchaseRepo.GetById(ctx, id)
}
func (uc *PurchaseUsecase) Delete(ctx context.Context, id string) error {
	return uc.purchaseRepo.Delete(ctx, id)
}
