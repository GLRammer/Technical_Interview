package main

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v2"
)

type Output struct{
    Message string `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        out := Output{Message: "My name is Gary Layman", Timestamp:time.Now()}
        return c.JSON(out)
    })

    log.Fatal(app.Listen(":3000"))
}