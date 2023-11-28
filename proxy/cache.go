package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"strings"
	"sync"
	"time"
)

var cache = &CacheStore{v: map[string]Cache{}}

type CacheStore struct {
	sync.Mutex
	v map[string]Cache
}
type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

func EvictCacheHandler(c *fiber.Ctx) error {
	key := strings.TrimPrefix(c.Path(), "/limit")
	if _, ok := cache.v[key]; !ok {
		return fiber.ErrNotFound
	}

	cache.Delete(key)
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func NewCacheProxy(key string, ttl time.Duration) CacheProxy {
	return CacheProxy{
		key: key,
		ttl: ttl,
	}
}

func (p CacheProxy) Accept(key string) bool {
	return p.key == key
}

func (p CacheProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	//key := c.Params("key")
	if r, ok := cache.v[path]; ok && r.ttl.After(time.Now()) {
		c.Response().SetBody(r.body)
		c.Response().Header.Add(fiber.HeaderCacheControl, fmt.Sprintf("max-age:%d", p.ttl/time.Second))
		c.Response().Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return nil
	}

	//url := "https://mocki.io" + strings.TrimPrefix(path, key+"/") //todo implement
	url := "https://google.com.tr"

	fmt.Printf("http request redirecting to %s \n", url)

	if err := proxy.Do(c, url); err != nil {
		return err
	}

	ch := Cache{
		ttl:  time.Now().Add(p.ttl),
		body: c.Response().Body(),
	}

	cache.Set(path, ch)

	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}

func (c *CacheStore) Set(key string, cache Cache) {
	c.Lock()
	c.v[key] = cache
	c.Unlock()
}

func (c *CacheStore) Delete(key string) {
	c.Lock()
	delete(c.v, key)
	c.Unlock()
}
