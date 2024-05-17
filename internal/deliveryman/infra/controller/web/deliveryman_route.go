package web

import (
	"github.com/labstack/echo/v4"
)

func LoadDeliverymanRoutes(apiGroup *echo.Group, handler DeliverymanHandlerJSON) {
	group := apiGroup.Group("/deliveryman")

	group.POST("", handler.CreateDeliveryman)
}
