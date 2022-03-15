package main

import (
	"context"
	"github.com/joho/godotenv"
	"integrate-go-with-redis/infrastructure"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	initial := infrastructure.NewRedis()

	redisCheckHealth := initial.Ping(ctx)

	log.Println(redisCheckHealth)
}
