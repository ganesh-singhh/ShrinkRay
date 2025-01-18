package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Avon11/ShrinkRay/internal/db"
	api "github.com/Avon11/ShrinkRay/internal/handler"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var ctx = context.Background()
var rdb *redis.Client

type RedisCreds struct {
	Addr     string
	Password string
	Database int
}

func init() {
	log.Println("Initializing application...")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Get db connection from env file
	var rdsCreds RedisCreds
	rdsCreds.Addr = os.Getenv("REDIS_ADDR")
	rdsCreds.Password = os.Getenv("REDIS_PASSWORD")
	rdsCreds.Database = cast.ToInt(os.Getenv("REDIS_DB"))

	log.Printf("Redis Configurations: %+v\n", rdsCreds)
	// Initialize Redis client
	log.Println("Connecting to Redis...")
	rdb = redis.NewClient(&redis.Options{
		Addr:     rdsCreds.Addr,
		Password: rdsCreds.Password,
		DB:       rdsCreds.Database,
	})

	// Test Redis connection
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	}
	_, err = db.CheckConnectivity()
	if err != nil {
		log.Printf("Failed to connect to Mongo DB! %v", err)
	}
}

func main() {
	router, err := api.SetupAPIHandler(rdb)
	if err != nil {
		log.Fatalf("Failed to setup API handler: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Close Redis connection
	defer rdb.Close()

	srv.ListenAndServe()
	log.Println("Server exiting")
}
