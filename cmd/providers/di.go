package providers

import (
	"fbc-bookings/internal/app"
	"fbc-bookings/internal/infra/adapters/postgres/repo"
	"fbc-bookings/internal/infra/api/handler"
	"fbc-bookings/internal/infra/api/router"
	"fbc-bookings/internal/infra/api/router/groups"
	"fbc-bookings/internal/infra/resources/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(postgres.NewConnection)

	_ = Container.Provide(repo.NewRepository)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewBookingGroup)

	_ = Container.Provide(handler.NewBookingHandler)

	_ = Container.Provide(app.NewBookingApp)

	return Container
}
