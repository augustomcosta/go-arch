package domainsvc

import (
	"fmt"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/entity"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/vo"
)

func CreateDeliverymanList(names []string, ages []int, addresses []vo.Address) (*[]entity.Deliveryman, error) {
	if len(names) != len(ages) {
		return nil, fmt.Errorf("Nomes e idades não tem o mesmo tamanho.")
	}

	var deliverymanList []entity.Deliveryman
	for i, _ := range names {
		deliveryman, err := entity.NewDeliveryman(names[i], ages[i], &addresses[i])

		if err != nil {
			return nil, err
		}

		deliverymanList = append(deliverymanList, *deliveryman)
	}

	return &deliverymanList, nil
}
