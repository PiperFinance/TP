package configs

import (
	"context"
	"fmt"
	"time"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/go-redis/redis/v8"
	"github.com/jellydator/ttlcache/v3"
	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"TP/schema"
)

var (
	// CacheManager *cache.Cache[any]
	RedisClient     *redis.Client
	CacheManager    *cache.ChainCache[any]
	TokenPriceCache *cache.ChainCache[float64]
	TTLCache        = ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](15 * time.Second),
	)
)

func LoadCache() {
	redisOption, err := redis.ParseURL(Config.RedisUrl.String())
	if err != nil {
		log.Fatalf("Cache :: %s", err)
	}

	ctx := context.Background()
	gCache := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := store.NewGoCache(gCache)
	RedisClient = redis.NewClient(redisOption)
	if cmd := RedisClient.Ping(ctx); cmd.Err() != nil {
		log.Error(cmd.Err())
		CacheManager = cache.NewChain[any](
			cache.New[any](gocacheStore),
		)
		TokenPriceCache = cache.NewChain[float64](
			cache.New[float64](gocacheStore),
		)
	} else {
		redisStore := store.NewRedis(RedisClient)
		CacheManager = cache.NewChain[any](
			cache.New[any](redisStore),
		)
		TokenPriceCache = cache.NewChain[float64](
			cache.New[float64](redisStore),
		)
	}

	err = CacheManager.Set(ctx, "Connected", "YES", store.WithExpiration(15*time.Second))

	if err != nil {
		panic(err)
	}
	_, err = CacheManager.Get(ctx, "Connected")
	if err != nil {
		log.Fatalf("unable to get cache key '%s' ", err)
	}
	// fmt.Printf("%#+v\n", value)

	err = TokenPriceCache.Set(ctx, "Connected", 0.0020922440899582904, store.WithExpiration(15*time.Second))

	if err != nil {
		panic(err)
	}
	_, err = TokenPriceCache.Get(ctx, "Connected")
	if err != nil {
		log.Fatalf("unable to get cache key '%s' ", err)
	}
	// fmt.Printf("%#+v\n", value)
}

func InsertedTSHKey(tokenId schema.TokenId, level schema.TimeSeriesLevel) (string, string) {
	return "TP:INTIDs", fmt.Sprintf("id:%s:%s", tokenId, level)
}

// return fmt.Sprintf("UTHS:%d", chain), fmt.Sprintf("%s-%s", user.String(), token.String())
