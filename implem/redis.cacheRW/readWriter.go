package urlRw

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/saeidraei/go-realworld-clean/uc"
)

type rw struct {
	client *redis.Client
}

func New() uc.CacheRW {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rw{
		client: client,
	}
}
func (rw rw) Set(key string, value interface{}, ttl time.Duration) error {

	err := rw.client.Set(key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rw rw) Get(key string) (interface{}, error) {
	val, err := rw.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	//fmt.Println("cache hit")

	return val, nil
}
