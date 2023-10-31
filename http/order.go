package http

import (
	"context"
	"go-mysql-redis/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name   string  `json:"name"`
	Side   bool    `json:"side"`
	Price  int64   `json:"price"`
	Amount float64 `json:"amount"`
}

func (s Server) CreateOrder(c *fiber.Ctx) error {
	_, cancle := context.WithTimeout(context.Background(), time.Second*20)
	defer cancle()
	order := new(Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"fail": err})
	}
	ID, err := s.service.CreateOrder(service.Order{
		Name:   order.Name,
		Side:   order.Side,
		Price:  order.Price,
		Amount: order.Amount,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"fail": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": ID})
}

func (s Server) ListOrderRedis(c *fiber.Ctx) error {
	_, cancle := context.WithTimeout(context.Background(), time.Second*20)
	defer cancle()
	orders, err := s.service.ListOrdersRedis()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"fail": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": orders})
}
