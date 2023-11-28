package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/:key/*", ProxyHandler)

	app.Delete("limit/:key/*", ResetLimitHandler)

	app.Delete("cache/:key/*", EvictCacheHandler)

	app.Listen(":3000")
}
