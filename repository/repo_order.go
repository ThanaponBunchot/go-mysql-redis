package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mysql-redis/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name   string
	Side   bool
	Price  int64
	Amount float64
}

type OrderRepository interface {
	// method
	CreateOrder(order Order) (uint, error)
	ListOrderRedis() (orders []Order, err error)
}
type orderRepository struct {
	tx          *gorm.DB
	redisClient *redis.Client
}

func NewOrderRepository(tx *gorm.DB, redisClient *redis.Client) orderRepository {
	return orderRepository{
		tx, redisClient,
	}
}

func (o orderRepository) CreateOrder(order Order) (uint, error) {
	if res := o.tx.Model(&model.Order{}).Create(&order); res.Error != nil {
		return 0, errors.Wrap(res.Error, "unable to create order")
	}
	return order.ID, nil
}

func (o orderRepository) ListOrderRedis() (orders []Order, err error) {
	//get redis
	ctx := context.Background()
	key := "repository::order"
	//unmarshall >>>>
	//string => byte => struct
	orderJSON, err := o.redisClient.Get(ctx, key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(orderJSON), &orders)
		if err == nil {
			fmt.Println("redis")
			return orders, nil
		}
	}
	//get database
	res := o.tx.Limit(100).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println("database")
	// marshall
	// struct >> btye >> string

	byteOrder, err := json.Marshal(orders)
	if err != nil {
		return nil, err
	}
	//set redis
	if err := o.redisClient.Set(ctx, key, string(byteOrder), time.Second*10).Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
