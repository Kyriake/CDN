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
	//website favicon
	app.Use(favicon.New(favicon.Config{
		File: "./images/favicon.ico",
	}))


	app.Static("/images", "./images")

	//returns the html page when you first load up the site	
	app.Get("/", func(c *fiber.Ctx) error {
		 return c.SendFile("./site/home.html")
	})

	//returns an html file when their is a 404 error
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendFile("./site/404.html")
	})

	//returns an html file when their is a 401 error
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(401).SendFile("./site/401.html")
	})
	//Change port by changing the numbers after the :
	log.Fatal(app.Listen("0.0.0.0:3000"))
}


