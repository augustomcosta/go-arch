package config

import (
	"github.com/augustomcosta/go-arch/internal/deliveryman/infra/controller/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func LoadRoutes(container *DIContainer) *echo.Echo {
	mux := echo.New()

	// Middlewares
	mux.Use(middleware.Logger())
	setupCORS(mux)

	// Index Endpoint
	mux.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Foi possível")
	})

	// API Endpoints
	apiGroup := mux.Group("/api")

	web.LoadDeliverymanRoutes(apiGroup, container.ControllerContainer.DeliverymanWebController)

	return mux
}

func setupCORS(mux *echo.Echo) {
	mux.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}
