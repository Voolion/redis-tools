package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	tools "github.com/zehuamama/redis-tools"
)

func main() {
	rc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	myClient := tools.NewTools(rc)
	val, err := myClient.Cas(context.Background(), "hao", "new", "jun")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val) // true

}
