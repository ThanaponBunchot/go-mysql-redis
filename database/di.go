package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewShema(model interface{}, modelName string) *gorm.DB {
	dsn := "admin:password@tcp(127.0.0.1:3307)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("unable to connect database")
	}
	fmt.Println("database connect!")
	InitTable(model, modelName, db)
	return db
}

func InitTable(model interface{}, modelName string, db *gorm.DB) {
	err := db.AutoMigrate(model)
	if err != nil {
		fmt.Printf("fail to migrate %v", modelName)
		panic("unable to auto migrate")
	}
	fmt.Printf("migrate %v complete!", modelName)
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("redis connect!")
	return client
}
