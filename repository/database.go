package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"fmt"
)
 
type Database struct {
	Client *redis.Client
}
 
var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.Background()
)
 
func NewDatabase() (*Database, error) {
	client := redis.NewClient(&redis.Options{
	   Addr: "redis:6379",
	   Password: "",
	   DB: 0,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		fmt.Println("Error Connecting to Redis")
	   	return nil, err
	}

	return &Database{
	   Client: client,
	}, nil
}
