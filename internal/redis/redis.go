package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redisam struct {
	Db *redis.Client
}

func (r Redisam) Init() *Redisam {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	Redisam := &Redisam{Db: client}
	return Redisam

}

func (r Redisam) RedisInsert(key, value string) {

	err := r.Db.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r Redisam) RedisRead(key string) (string, string) {

	val, err := r.Db.Get(key).Result()
	if err != nil {
		fmt.Println("some problem")
	}
	// fmt.Println("key", key, val)
	return val, key
}
