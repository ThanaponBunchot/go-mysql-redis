package http

import (
	"go-mysql-redis/service"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	service service.Service
	app     *fiber.App
}

type Dependency struct {
	Service service.Service
	App     *fiber.App
}

func New(d Dependency) Server {
	return Server{
		service: d.Service,
		app:     d.App,
	}
}
