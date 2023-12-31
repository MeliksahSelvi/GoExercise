package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

var Proxies = []Proxy{
	NewLimitProxy("limitUser", 3, 3*time.Minute),
	NewCacheProxy("cacheUser", 10*time.Second),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range Proxies {
		if v.Accept(c.Params("key")) {
			return v.Proxy(c)
		}
	}

	c.Response().SetStatusCode(404)
	return nil
}
