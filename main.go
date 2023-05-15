package main

import (
	"fmt"
	"strconv"

	redismain "github.com/go-redis/redis"

	"main.go/internal/fiber"
	"main.go/internal/redis"
)

type Daberna struct {
	db     redis.Redisam
	server fiber.Serveram
}

func main() {
	fmt.Println("sallam")
	client := redismain.NewClient(&redismain.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	Redisam := &redis.Redisam{Db: client}

	redis.Redisam.Init(*Redisam)
	daberna := &Daberna{db: *Redisam}
	Daberna.dabernam(*daberna)

}
func (d Daberna) dabernam() {

	for i := 1; i < 100; i++ {
		d.db.RedisInsert(strconv.Itoa(i), "false")
	}

	for i := 1; i < 100; i++ {
		d.db.RedisRead(strconv.Itoa(i))
	}

	d.server.Httpserverengin(d.db)

}
