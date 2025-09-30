package open_ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
	"github.com/rs/zerolog/log"
)

type Builder struct {
	client openai.Client
	params responses.ResponseNewParams
}

func (b Builder) WithSearch(country string) Builder {
	b.params = addSearch(b.params, country)

	return b
}

func (b Builder) Text() (answer string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := b.client.Responses.New(ctx, b.params)
	if err != nil {
		return "", fmt.Errorf("error creating response: %w", err)
	}

	logRequest(resp)

	if resp.Error.Message != "" {
		return "", fmt.Errorf("response error: %s", resp.Error.Message)
	}

	for _, output := range resp.Output {
		for _, content := range output.Content {
			if content.Text != "" {
				answer += content.Text + "\n"
			}
		}

		if output.Action.Query != "" {
			log.Info().Str("search_query", output.Action.Query).Msg("open_ai: search request!")
		}
	}

	return answer, nil
}
