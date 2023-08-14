package main

import (
	"fmt"
	"log"

	"fbc-bookings/cmd/providers"
	"fbc-bookings/config"
	"fbc-bookings/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().Server.Host
	serverPort = config.Environments().Server.Port
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
