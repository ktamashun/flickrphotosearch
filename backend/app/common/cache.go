package common

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var Cache *memcache.Client

func init() {
	Cache = memcache.New(fmt.Sprintf("%s:%s", AppConfig.MemcacheHost, AppConfig.MemcachePort))
	log.Println("Memcache connection OK")
}
