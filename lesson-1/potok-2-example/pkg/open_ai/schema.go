package open_ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (b Builder) Schema(v any, name, desc string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	schema, err := makeSchema(v, name, desc)
	if err != nil {
		return fmt.Errorf("error creating schema: %w", err)
	}

	b.params.Text = schema

	resp, err := b.client.Responses.New(ctx, b.params)
	if err != nil {
		return fmt.Errorf("error creating response: %w", err)
	}

	logRequest(resp)

	if resp.Error.Message != "" {
		return fmt.Errorf("response error: %s", resp.Error.Message)
	}

	var obj string

	for _, output := range resp.Output {
		for _, content := range output.Content {
			if content.Text != "" {
				obj = content.Text
				break
			}
		}

		if output.Action.Query != "" {
			log.Info().Str("search_query", output.Action.Query).Msg("open_ai: search request!")
		}
	}

	err = json.Unmarshal([]byte(obj), v)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}
