package main

import (
	"go-mysql-redis/database"
	"go-mysql-redis/http"
	"go-mysql-redis/model"
	"go-mysql-redis/repository"
	"go-mysql-redis/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	redisClient := database.InitRedis()
	orderCol := database.NewShema(&model.Order{}, "order")

	orderRepo := repository.NewOrderRepository(orderCol, redisClient)

	sv := service.New(service.Dependency{
		OrderRepository: orderRepo,
	})

	serv := http.New(http.Dependency{
		Service: sv,
		App:     app,
	})
	serv.Route()

	if err := app.Listen(":9000"); err != nil {
		log.Panic(err, "Server is not running")
	}
}
