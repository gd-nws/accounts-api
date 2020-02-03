package Cache

import ("github.com/patrickmn/go-cache"
	"time"
)

var MemoryCache *cache.Cache

func CreateCache() {
	MemoryCache = cache.New(5*time.Minute, 10*time.Minute)
}