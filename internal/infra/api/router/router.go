package router

import (
	"net/http"

	"fbc-bookings/internal/domain/entity"
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
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, latency=${latency_human} status=${status}\n",
	}))

	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.bookings.Resource(basePath)

	r.server.RouteNotFound("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, entity.Response{
			Message: "Sorry, page not found",
		})
	})
}
