package repository_impl

import (
	"context"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/entity"
	db "github.com/augustomcosta/go-arch/sql/sqlc_gen"
	"log/slog"
)

type DeliverymanRepositorySQLC struct {
	db *db.Queries
}

func NewDeliverymanRepositorySQLC(db *db.Queries) *DeliverymanRepositorySQLC {
	return &DeliverymanRepositorySQLC{db: db}
}

func (r *DeliverymanRepositorySQLC) Save(ctx context.Context, deliveryman *entity.Deliveryman) error {
	params := db.CreateDeliverymanParams{
		ID:           deliveryman.ID,
		Name:         deliveryman.Name,
		Age:          deliveryman.Age,
		AddressState: deliveryman.Address.State,
		AddressCity:  deliveryman.Address.City,
	}

	err := r.db.CreateDeliveryman(ctx, params)
	if err != nil {
		return err
	}

	slog.Info("Deliveryman created successfully.",
		slog.Group("Fields",
			slog.String("ID", deliveryman.ID.String()),
			slog.String("Name", deliveryman.Name)))

	return nil
}
