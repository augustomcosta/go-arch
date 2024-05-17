package usecases

import (
	"github.com/augustomcosta/go-arch/internal/deliveryman/application/service"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/entity"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/repository"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/vo"
	"time"
)

type CreateDeliverymanInputDTO struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	AddressState string `json:"state"`
	AddressCity  string `json:"city"`
}

type CreateDeliverymanOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateDeliverymanUseCase struct {
	repository repository.DeliverymanRepository
}

func (uc *CreateDeliverymanUseCase) Execute(input *CreateDeliverymanInputDTO) (*CreateDeliverymanOutputDTO, error) {
	createdAddress, err := vo.NewAddress(input.AddressState, input.AddressCity)
	if err != nil {
		return nil, err
	}

	createdDeliveryman, err := entity.NewDeliveryman(input.Name, input.Age, createdAddress)
	if err != nil {
		return nil, err
	}

	err = applicationsvc.VerifySerasaSituation(createdDeliveryman)
	if err != nil {
		return nil, err
	}

	// EXEMPLO CONCEITUAL
	//deliverymen, err := domainsvc.CreateDeliverymanList(input.Name)
	//if err != nil {
	//	return nil, err
	//}

	return &CreateDeliverymanOutputDTO{
		ID:        createdDeliveryman.ID.String(),
		CreatedAt: time.Now(),
	}, nil
}
