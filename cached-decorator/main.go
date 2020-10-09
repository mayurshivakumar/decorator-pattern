package main

import (
"fmt"
 "github.com/patrickmn/go-cache"
	"time"
)

type Getter struct {}

func (g *Getter)GetNameByID(ID string) string{
	// dummy get
	if ID == "123" {
		return "Name 123"
	}
	return "Name 1234 ..."
}

func NewGetter() *Getter{
	return &Getter{}
}

type Get interface {
	GetNameByID(ID string) string
}

type CachedGetter struct {
	Get
	cache *cache.Cache
}

func (c *CachedGetter) CachedGetNameByID(ID string) string {
	name, ok := c.cache.Get(ID)
	if ok {
		fmt.Println("found in cache...")
		return name.(string)
	}
	fmt.Println("caching it....")
	name = c.Get.GetNameByID(ID)
	c.cache.Set(ID, name, 5 * time.Minute)
	return name.(string)
}

func NewCachedGetter(g Get) *CachedGetter {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &CachedGetter{
		Get:g,
		cache: c,
	}
}


func main() {
	// without caching
	g := NewGetter()
	fmt.Println(g.GetNameByID("123"))

	// with caching
	gg := NewGetter()
	cg := NewCachedGetter(gg)
	// first call will not be in cache will be cached
	fmt.Println(cg.CachedGetNameByID("1234"))
	// second call will return from cache
	fmt.Println(cg.CachedGetNameByID("1234"))
}
