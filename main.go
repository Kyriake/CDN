package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)
func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration: 30 * time.Second,
		CacheControl: true,
	
	}))
	app.Use(favicon.New(favicon.Config{
		File: "./images/favicon.ico",
	}))

	app.Static("/images", "./images")

	app.Get("/", func(c *fiber.Ctx) error {
		 return c.SendFile("./site/home.html")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendFile("./site/404.html")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(401).SendFile("./site/401.html")
	})
	log.Fatal(app.Listen("0.0.0.0:3000"))
}


