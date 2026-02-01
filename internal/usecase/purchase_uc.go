package usecase

import (
	"context"
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
)

type PurchaseUsecase struct {
	purchaseRepo domain.PurchaseRepo
}

func NewPurchaseUsecase(purchaseRepo domain.PurchaseRepo) *PurchaseUsecase {
	return &PurchaseUsecase{
		purchaseRepo: purchaseRepo,
	}
}

func (uc *PurchaseUsecase) Create(ctx context.Context, purchase *domain.Purchase) error {
	purchase.Payment.PaidAt = time.Now()
	return uc.purchaseRepo.Create(ctx, purchase)
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
