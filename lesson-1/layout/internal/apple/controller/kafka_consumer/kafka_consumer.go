package kafka_consumer

import (
	"context"
	"github.com/golang-school/layout/internal/apple/usecase"
	"github.com/golang-school/layout/pkg/kafka_reader"
	"github.com/rs/zerolog/log"
)

func AppleConsumer(reader *kafka_reader.Reader, uc *usecase.UseCase) {
	ctx := context.Background()

	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Error().Err(err).Msg("kafka_consumer.AppleConsumer: reader.FetchMessage")

			break
		}

		// UseCase call here

		log.Info().
			Str("topic", m.Topic).
			Int("partition", m.Partition).
			Int64("offset", m.Offset).
			Str("key", string(m.Key)).
			Str("value", string(m.Value)).
			Msg("kafka_consumer.AppleConsumer: reader.FetchMessage")

		if err = reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka_consumer.AppleConsumer: reader.CommitMessages")
		}
	}
}
