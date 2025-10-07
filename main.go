package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Output struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		out := Output{Message: "My name is Gary Layman", Timestamp: time.Now().Unix()}
		outString, _ := json.Marshal(out)
		return c.Send(outString)
	})

	log.Fatal(app.Listen(":80"))
}
