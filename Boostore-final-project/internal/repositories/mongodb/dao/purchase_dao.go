package dao

import (
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PurchaseDoc struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserId      bson.ObjectID `bson:"user_id"`
	TotalAmount float32       `bson:"total_amount"`
	Items       []ItemDoc     `bson:"items"`
	Payment     PaymentDoc    `bson:"payment"`
}

type ItemDoc struct {
	GameId  bson.ObjectID `bson:"game_id"`
	PriceAt float32       `bson:"price_at"`
}

type PaymentDoc struct {
	PaymentMethod string    `bson:"payment_method"`
	PaymentStatus string    `bson:"payment_status"`
	PaidAt        time.Time `bson:"paid_at"`
}

func FromPurchaseDomain(p *domain.Purchase) *PurchaseDoc {
	itemDocs := make([]ItemDoc, len(p.Items))
	for i, item := range p.Items {
		gID, _ := bson.ObjectIDFromHex(item.GameId)
		itemDocs[i] = ItemDoc{
			GameId:  gID,
			PriceAt: item.PriceAt,
		}
	}

	uID, _ := bson.ObjectIDFromHex(p.UserId)

	doc := &PurchaseDoc{
		UserId:      uID,
		TotalAmount: p.TotalAmount,
		Items:       itemDocs,
		Payment: PaymentDoc{
			PaymentMethod: p.Payment.PaymentMethod,
			PaymentStatus: p.Payment.PaymentStatus,
			PaidAt:        p.Payment.PaidAt,
		},
	}

	if p.ID != "" {
		if objID, err := bson.ObjectIDFromHex(p.ID); err == nil {
			doc.ID = objID
		}
	}
	return doc
}

func (d PurchaseDoc) ToDomain() *domain.Purchase {
	items := make([]domain.Item, len(d.Items))
	for i, item := range d.Items {
		items[i] = domain.Item{
			GameId:  item.GameId.Hex(),
			PriceAt: item.PriceAt,
		}
	}

	return &domain.Purchase{
		ID:          d.ID.Hex(),
		UserId:      d.UserId.Hex(),
		TotalAmount: d.TotalAmount,
		Items:       items,
		Payment: domain.Payment{
			PaymentMethod: d.Payment.PaymentMethod,
			PaymentStatus: d.Payment.PaymentStatus,
			PaidAt:        d.Payment.PaidAt,
		},
	}
}
