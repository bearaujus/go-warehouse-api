package postgres_cacher

import (
	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
	"time"
)

type postgresCacherImpl struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewPostgresCacher(rdb *redis.Client, ttl time.Duration) caches.Cacher {
	return &postgresCacherImpl{rdb: rdb, ttl: ttl}
}
