package http_server

import (
	"context"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	Port string `envconfig:"HTTP_PORT" default:"8080"`
}

type Server struct {
	server *http.Server
	notify chan error
}

func New(handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         net.JoinHostPort("", port),
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	go s.start()

	log.Info().Msg("HTTP Server started on port: " + port)

	return s
}

func (s *Server) start() {
	s.notify <- s.server.ListenAndServe()
	close(s.notify)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("server - Close - s.server.Shutdown")
	}

	log.Info().Msg("HTTP Server closed")
}
