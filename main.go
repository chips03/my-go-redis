package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chips03/my-go-redis/database"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	godotenv.Load()
	fmt.Printf("===== %v =====\n", os.Getenv("APP_NAME"))
	rdb := database.NewRedisClient(0)

	ttl := 2 * time.Minute
	err := rdb.Set(ctx, "key-1", "Hello world", ttl)
	if err != nil {
		log.Println(err)
	}
}
