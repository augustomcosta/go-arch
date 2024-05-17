package entity

import (
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/vo"
	"github.com/google/uuid"
)

type Deliveryman struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Age     int32       `json:"number"`
	Address *vo.Address `json:"address"`
}

func NewDeliveryman(name string, age int32, address *vo.Address) (*Deliveryman, error) {
	return &Deliveryman{
		ID:      uuid.New(),
		Name:    name,
		Age:     age,
		Address: address,
	}, nil
}
