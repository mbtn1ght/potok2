package open_ai

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
	"github.com/rs/zerolog/log"
)

const timeout = 120 * time.Second

type Client struct {
	client openai.Client
}

func New(key string) Client {
	return Client{
		client: openai.NewClient(option.WithAPIKey(key)),
	}
}

func (c Client) Get(query string) Builder {
	return Builder{
		client: c.client,
		params: defaultParams(query),
	}
}

func defaultParams(query string) responses.ResponseNewParams {
	return responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(query),
		},
		Model:           openai.ChatModelGPT4oMini,
		Store:           openai.Bool(false),
		Temperature:     openai.Float(0.2),
		MaxOutputTokens: openai.Int(1500),
	}
}

func addSearch(params responses.ResponseNewParams, country string) responses.ResponseNewParams {
	params.Tools = append(params.Tools, responses.ToolUnionParam{
		OfWebSearchPreview: &responses.WebSearchToolParam{
			Type: responses.WebSearchToolTypeWebSearchPreview,
			UserLocation: responses.WebSearchToolUserLocationParam{
				Country: openai.String(country),
			},
			SearchContextSize: responses.WebSearchToolSearchContextSizeMedium,
		},
	})

	return params
}

func makeSchema(v any, name, desc string) (responses.ResponseTextConfigParam, error) {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}

	schema := reflector.Reflect(v)
	if schema == nil {
		return responses.ResponseTextConfigParam{}, errors.New("schema is nil")
	}

	data, err := schema.MarshalJSON()
	if err != nil {
		return responses.ResponseTextConfigParam{}, fmt.Errorf("error marshalling schema: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return responses.ResponseTextConfigParam{}, fmt.Errorf("error unmarshalling schema: %w", err)
	}

	return responses.ResponseTextConfigParam{
		Format: responses.ResponseFormatTextConfigUnionParam{
			OfJSONSchema: &responses.ResponseFormatTextJSONSchemaConfigParam{
				Name:        name,
				Description: openai.String(desc),
				Schema:      result,
				Strict:      openai.Bool(true),
			},
		},
	}, nil
}

func logRequest(resp *responses.Response) {
	log.Info().
		Str("status", string(resp.Status)).
		Str("model", resp.Model).
		Int64("input_tokens", resp.Usage.InputTokens).
		Int64("output_tokens", resp.Usage.OutputTokens).
		Msg("open_ai: request completed")
}
