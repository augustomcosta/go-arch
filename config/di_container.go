package config

import (
	"database/sql"
	"github.com/augustomcosta/go-arch/internal/deliveryman/application/usecase"
	"github.com/augustomcosta/go-arch/internal/deliveryman/domain/repository"
	"github.com/augustomcosta/go-arch/internal/deliveryman/infra/controller/web"
	"github.com/augustomcosta/go-arch/internal/deliveryman/infra/repository_impl"
)

// DIContainer Configure Project Initialization
type DIContainer struct {
	RepositoryContainer *RepositoryContainer
	UseCaseContainer    *UseCaseContainer
	ControllerContainer *ControllerContainer
}

func InitDIContainer(dbcon *sql.DB) *DIContainer {
	// Init repositories
	dbQueries := NewDBQueries(dbcon)
	deliverymanRepository := repository_impl.NewDeliverymanRepositorySQLC(dbQueries)

	repositoryContainer := newRepositoryContainer(
		deliverymanRepository,
	)

	// Init use cases
	createDeliverymanUseCase := usecase.NewCreateDeliverymanUseCase(repositoryContainer.DeliverymanRepository)

	useCaseContainer := newUseCaseContainer(
		createDeliverymanUseCase,
	)

	// Init controllers
	agencyHandler := web.NewDeliverymanHandlerJSON(useCaseContainer.CreateDeliverymanUseCase)

	controllerContainer := newControllerContainer(
		agencyHandler,
	)

	return &DIContainer{
		RepositoryContainer: repositoryContainer,
		UseCaseContainer:    useCaseContainer,
		ControllerContainer: controllerContainer,
	}
}

// RepositoryContainer Configure Repositories
type RepositoryContainer struct {
	DeliverymanRepository repository.DeliverymanRepository
}

func newRepositoryContainer(
	deliverymanRepository repository.DeliverymanRepository,
) *RepositoryContainer {
	return &RepositoryContainer{
		DeliverymanRepository: deliverymanRepository,
	}
}

// UseCaseContainer Configure UseCases
type UseCaseContainer struct {
	CreateDeliverymanUseCase usecase.CreateDeliverymanUseCase
}

func newUseCaseContainer(
	createDeliverymanUseCase *usecase.CreateDeliverymanUseCase,
) *UseCaseContainer {
	return &UseCaseContainer{
		CreateDeliverymanUseCase: *createDeliverymanUseCase,
	}
}

// ControllerContainer Configure Controllers
type ControllerContainer struct {
	DeliverymanWebController web.DeliverymanHandlerJSON
}

func newControllerContainer(
	deliverymanWebController *web.DeliverymanHandlerJSON,
) *ControllerContainer {
	return &ControllerContainer{
		DeliverymanWebController: *deliverymanWebController,
	}
}
