package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/n3k0fi5t/ShortTheURL/common/config"
)

const (
	dsnFormat       = "%s:%s"
	defaultDB       = 0
	defaultPassword = ""
)

type RDB struct {
	RDB *redis.Client
}

var ctx = context.Background()

func Init(cfg *config.Config) (*RDB, error) {
	dsn := fmt.Sprintf(dsnFormat, cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: defaultPassword,
		DB:       defaultDB,
	})

	// simply test by set
	_, err := rdb.Set(ctx, "key", "0", time.Second).Result()
	return &RDB{rdb}, err
}
