package web

import (
	"github.com/augustomcosta/go-arch/internal/deliveryman/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeliverymanController struct {
	CreateDeliverymanUseCase usecase.CreateDeliverymanUseCase
}

func (ctrl *DeliverymanController) CreateDeliveryman(c echo.Context) error {
	var request usecase.CreateDeliverymanInputDTO
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Nao foi possivel")
	}

	output, err := ctrl.CreateDeliverymanUseCase.Execute(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Não foi possivel")
	}

	return c.JSON(http.StatusOK, output)
}
