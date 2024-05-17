package web

import (
	"github.com/augustomcosta/go-arch/internal/deliveryman/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeliverymanHandlerJSON struct {
	CreateDeliverymanUseCase usecase.CreateDeliverymanUseCase
}

func NewDeliverymanHandlerJSON(createDeliverymanUseCase usecase.CreateDeliverymanUseCase) *DeliverymanHandlerJSON {
	return &DeliverymanHandlerJSON{
		CreateDeliverymanUseCase: createDeliverymanUseCase,
	}
}

func (ctrl *DeliverymanHandlerJSON) CreateDeliveryman(c echo.Context) error {
	var request usecase.CreateDeliverymanInputDTO
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Nao foi possivel")
	}

	output, err := ctrl.CreateDeliverymanUseCase.Execute(c.Request().Context(), &request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, output)
}
