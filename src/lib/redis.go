package lib

import (
	"context"
	"fmt"
	config "heartBeat/src/config"
	"os"

	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background();
var rdb *redis.Client;

func RedisConnect() {

	fmt.Fprintln(os.Stderr,"::[redis.go]:: redis connection started...");

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASSWORD, // no password set
		DB:       config.REDIS_DATABASE,  // use default DB
	})

	fmt.Fprintln(os.Stderr,"::[redis.go]:: go-redis connection success...");

}

func RedisHGetAll(key string) (map[string]string, error) {

	result, err := rdb.HGetAll(ctx, key).Result();

	if (err != nil) {
		return nil, err;
	}

	return result, nil;

}