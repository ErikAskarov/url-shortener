package main

import (
	"context"
	"fmt"
	"log"
	"url-shortener/internal/db"
)

func main() {
	log.Print("Listening on port 8080")

	redis := db.NewRedisStorage("localhost:6379", "")
	status, err := redis.Ping(context.Background())
	if err != nil {
		fmt.Printf("redis connect failed: %s", err)
	}
	fmt.Printf("redis connect success: %s\n", status)

}
