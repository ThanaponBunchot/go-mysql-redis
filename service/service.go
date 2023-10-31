package service

import "go-mysql-redis/repository"

type Service interface {
	//method
	CreateOrder(order Order) (uint, error)
	ListOrdersRedis() ([]Order, error)
}
type Dependency struct {
	OrderRepository repository.OrderRepository
}

type service struct {
	orderRepo repository.OrderRepository
}

func New(d Dependency) Service {
	return service{
		orderRepo: d.OrderRepository,
	}
}
