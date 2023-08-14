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
}

func New(server *echo.Echo, bookings groups.Bookings) *Router {
	return &Router{
		server,
		bookings,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.CORS())
	r.server.Use(middleware.Recover())
	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.bookings.Resource(basePath)
}
