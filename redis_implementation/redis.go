package redis_implementation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go_databases/records"
	"log"
	"time"
)

type RedisConnection struct {
}

func (r *RedisConnection) Ping() {
	ctx, client, cancel := r.connect()
	defer cancel()

	result, err := client.Ping(ctx).Result()

	if err == nil {
		fmt.Printf("Connect Success %s\n", result)
	} else {
		fmt.Printf("Connect Error %s\n", err)
	}
}

func (r *RedisConnection) InsertDocument() {
	ctx, client, cancel := r.connect()
	defer cancel()

	data, err := json.Marshal(records.DefaultRecord1())

	if err != nil {
		log.Fatalln(err)
	}

	set, err := client.SetNX(ctx, "123-4231", data, redis.KeepTTL).Result()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful Insert %s\n", set)
	}
}

func (r *RedisConnection) SearchAllDocument() {
	ctx, client, cancel := r.connect()
	defer cancel()
	vals, err := client.Get(ctx, "123-4231").Result()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful Find %s\n", vals)
	}
}

func (r *RedisConnection) DeleteDocument() {
	ctx, client, cancel := r.connect()
	defer cancel()
	client.Del(ctx, "123-4231")

}

func (r *RedisConnection) UpdateDocument() {
	//TODO implement me
	panic("implement me")
}

func (r *RedisConnection) DeleteManyUsingFilter() {
	//TODO implement me
	panic("implement me")
}

func (r *RedisConnection) connect() (context.Context, *redis.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return ctx, client, cancel
}

func NewMongoConnection() records.Database {
	return &RedisConnection{}
}
