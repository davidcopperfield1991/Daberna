package fiber

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber"
	redis "main.go/internal/redis"
)

type Serveram struct {
	Httpserver fiber.App
	Db         redis.Redisam
}

func (s Serveram) Httpserverengin(db redis.Redisam) {

	app := fiber.New()

	s.Db = db

	app.Get("/call", s.call)

	log.Fatal(app.Listen(":3000"))
}

func (s Serveram) awailable() []string {
	awailable := []string{}
	for i := 1; i < 100; i++ {
		val, key := s.Db.RedisRead(strconv.Itoa(i))
		if val == "false" {
			awailable = append(awailable, key)
		}
	}
	return awailable

}

func (s Serveram) randombede() string {

	aw := s.awailable()
	max := len(aw)

	if max > 0 {
		random := rand.Intn(max)
		return aw[random]

	} else {
		return "ended"
	}

}

func (s Serveram) call(c *fiber.Ctx) {

	rand := s.randombede()

	if rand == "ended" {
		c.SendString("game ended")
	} else {
		value, _ := s.Db.RedisRead(rand)
		if value == "false" {
			s.Db.RedisInsert(rand, "true")
			c.Send(rand)
		} else {
			c.SendString("number readed befor !!!")
		}

	}
	fmt.Println(rand)

}
