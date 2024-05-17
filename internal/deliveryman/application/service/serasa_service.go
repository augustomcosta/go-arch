package applicationsvc

import (
	"fmt"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/entity"
)

func VerifySerasaSituation(deliveryman *entity.Deliveryman) error {
	if deliveryman.Age < 18 {
		return fmt.Errorf("Menores de idade não podem trabalhar de entregador.")
	}

	return nil
}
