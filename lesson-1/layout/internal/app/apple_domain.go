package app

import (
	"github.com/golang-school/layout/internal/apple/adapter/kafka_producer"
	"github.com/golang-school/layout/internal/apple/adapter/postgres"
	"github.com/golang-school/layout/internal/apple/adapter/redis"
	"github.com/golang-school/layout/internal/apple/controller/http_router"
	"github.com/golang-school/layout/internal/apple/controller/kafka_consumer"
	"github.com/golang-school/layout/internal/apple/usecase"
)

func AppleDomain(d Dependencies) {
	appleUseCase := usecase.New(
		postgres.New(d.Postgres.Pool),
		kafka_producer.New(d.KafkaWriter.Writer),
		redis.New(d.Redis.Client),
	)

	http_router.AppleRouter(d.RouterHTTP, appleUseCase)

	go kafka_consumer.AppleConsumer(d.KafkaReader, appleUseCase)
}
