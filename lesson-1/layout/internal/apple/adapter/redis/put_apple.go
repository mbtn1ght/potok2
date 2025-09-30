package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/otel/tracer"
)

func (r *Redis) PutApple(ctx context.Context, a entity.Apple) error {
	ctx, span := tracer.Start(ctx, "redis PutApple")
	defer span.End()

	data, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = r.client.Set(ctx, a.ID.String(), data, ttl).Err()
	if err != nil {
		return fmt.Errorf("r.client.Set: %w", err)
	}

	return nil
}
