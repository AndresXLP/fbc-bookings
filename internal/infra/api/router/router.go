package router

import (
	"fbc-bookings/internal/infra/api/handler"
	"fbc-bookings/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server   *echo.Echo
	bookings groups.Bookings
	payments groups.Payments
}

func New(server *echo.Echo, bookings groups.Bookings, payments groups.Payments) *Router {
	return &Router{
		server,
		bookings,
		payments,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.CORS())
	r.server.Use(middleware.Recover())
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, latency=${latency_human} status=${status}\n",
	}))

	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.bookings.Resource(basePath)
	r.payments.Resource(basePath)

	r.server.RouteNotFound("*", handler.NotFound)
}
