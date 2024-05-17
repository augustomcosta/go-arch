package repository

import (
	"context"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/entity"
)

type DeliverymanRepository interface {
	Save(ctx context.Context, deliveryman *entity.Deliveryman) error
}
