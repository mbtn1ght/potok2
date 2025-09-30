package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-school/potok-2-example/pkg/logger"
	"github.com/golang-school/potok-2-example/pkg/open_ai"
	"github.com/golang-school/potok-2-example/pkg/render"
	"github.com/rs/zerolog/log"
)

const (
	// Получить ключ на https://platform.openai.com
	key  = `<YOUR_API_KEY>`
	port = "8080"
)

// Сервер
func main() {
	logger.Init("info", true)

	r := chi.NewRouter()
	r.Post("/api/v1/about_country", Handler)

	log.Info().Msgf("server started on port %s", port)

	err := http.ListenAndServe(net.JoinHostPort("", port), r)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}

// Модели
type Input struct {
	Countries []string `json:"countries"`
}

type Output struct {
	ProcessingTime string   `json:"processing_time"`
	Results        []Result `json:"results"`
}

type Result struct {
	Country    string   `json:"country" jsonschema_description:"Страна"`
	Capital    string   `json:"capital" jsonschema_description:"Столица"`
	Population int      `json:"population" jsonschema_description:"Население"`
	Area       int      `json:"area" jsonschema_description:"Площадь в квадратных километрах"`
	Currency   string   `json:"currency" jsonschema_description:"Валюта"`
	Languages  []string `json:"languages" jsonschema_description:"Официальные языки (может быть несколько)"`
}

// Обработчик HTTP-запроса
func Handler(w http.ResponseWriter, r *http.Request) {
	// Принимаем JSON в теле запроса
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("failed to decode request body")

		return
	}

	// Ошибка, если поле countries пустое
	if len(input.Countries) == 0 {
		http.Error(w, "countries field is required", http.StatusBadRequest)
		log.Error().Msg("countries field is required")

		return
	}

	now := time.Now() // Хотим замерить время обработки запроса

	// Готовим ответ
	var (
		wg       sync.WaitGroup
		mx       sync.Mutex
		results  []Result
		errGroup error
	)

	for _, country := range input.Countries {
		wg.Add(1)

		go func() {
			defer wg.Done()

			var result Result
			result, errGroup = GetCountryInfo(country)

			mx.Lock()
			results = append(results, result)
			mx.Unlock()
		}()
	}

	wg.Wait()

	if errGroup != nil {
		http.Error(w, errGroup.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("failed to get country info")

		return
	}

	output := Output{
		ProcessingTime: time.Since(now).Round(time.Second).String(),
		Results:        results,
	}

	// Отправляем ответ в формате JSON
	render.JSON(w, output, http.StatusOK)
	log.Info().Msgf("processed countries: %v", input.Countries)
}

func GetCountryInfo(country string) (Result, error) {
	openai := open_ai.New(key)

	var (
		output     Result
		schemaName = "countries_info"
		schemaDesc = "Информация о стране"
		query      = "Ты опытный географ. Предоставь точную и актуальную информацию о стране: " + country
	)

	err := openai.Get(query).Schema(&output, schemaName, schemaDesc)
	if err != nil {
		return Result{}, fmt.Errorf("failed to query openai: %w", err)
	}

	return output, nil
}
