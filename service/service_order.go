package service

import (
	"go-mysql-redis/repository"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name   string
	Side   bool
	Price  int64
	Amount float64
	Time   time.Time
}

func (s service) CreateOrder(order Order) (uint, error) {
	id, err := s.orderRepo.CreateOrder(repository.Order{
		Name:   order.Name,
		Side:   order.Side,
		Price:  order.Price,
		Amount: order.Amount,
		Model: gorm.Model{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			DeletedAt: order.DeletedAt,
		},
	})
	if err != nil {
		return 0, errors.Wrap(err, "unable to create the order")
	}
	return id, nil
}

func (s service) ListOrdersRedis() ([]Order, error) {
	orders, err := s.orderRepo.ListOrderRedis()
	if err != nil {
		return nil, err
	}
	return RepotoServiceOrder(orders), nil
}

func RepotoServiceOrder(orders []repository.Order) []Order {
	arrOrder := []Order{}
	for _, order := range orders {
		arrOrder = append(arrOrder, Order{
			Name:   order.Name,
			Side:   order.Side,
			Price:  order.Price,
			Amount: order.Amount,
			Model: gorm.Model{
				ID:        order.ID,
				CreatedAt: order.CreatedAt,
				UpdatedAt: order.UpdatedAt,
				DeletedAt: order.DeletedAt,
			},
		})
	}
	return arrOrder
}
