package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)


type dataSource struct {
	DB          *sqlx.DB
	RedisClient *redis.Client
}

func initDS() (*dataSource, error) {
	log.Printf("initializing data sources.\n")
	
	os.Getenv("")

	dsn := fmt.Sprintf("")
	
	log.Printf("Connecting to MySQL.\n")
	db, err := sqlx.Open("pq", dsn)
	if err != nil {
		return nil , fmt.Errorf("error opening db: %v", err)
	}
	
	if err := db.Ping(); err != nil {
		return nil , fmt.Errorf("error connecting to db: %v", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	
	log.Printf("Coonecting to Redis. \n")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB: 0,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil , fmt.Errorf("error connecting to redis: %v", err)
	}

	return &dataSource{
		DB:          db,
		RedisClient: rdb,
	}, nil 
}


func (d *dataSource) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing MySQL: %v", err)
	}

	return nil 
}