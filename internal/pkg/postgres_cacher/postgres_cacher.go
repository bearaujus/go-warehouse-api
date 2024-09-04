package postgres_cacher

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
)

func (r *postgresCacherImpl) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	res, err := r.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	err = q.Unmarshal([]byte(res))
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (r *postgresCacherImpl) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	r.rdb.Set(ctx, key, res, r.ttl)
	return nil
}

func (r *postgresCacherImpl) Invalidate(ctx context.Context) error {
	var (
		cursor uint64
		keys   []string
	)
	for {
		var (
			k   []string
			err error
		)
		k, cursor, err = r.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", caches.IdentifierPrefix), 0).Result()
		if err != nil {
			return err
		}
		keys = append(keys, k...)
		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		if _, err := r.rdb.Del(ctx, keys...).Result(); err != nil {
			return err
		}
	}
	return nil
}
