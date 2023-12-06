package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type server struct {
	router *http.ServeMux
	logger *slog.Logger
}

func NewServer() *server {
	return &server{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		router: http.NewServeMux(),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.With("method", r.Method, "path", r.URL.Path).Info("request")
	s.router.ServeHTTP(w, r)
}

func (s *server) Start(address, port string) error {
	if port == "" {
		port = "8080"
	}

	s.configureRouter()
	addr := fmt.Sprintf("%s:%s", address, port)
	fmt.Println("Server started on", addr)
	return http.ListenAndServe(addr, s)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.homeMessage())
	s.router.HandleFunc("/time", s.sendTime())
}

func (s *server) homeMessage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprint(w, "<b>Привет!<b>")
	}
}

func (s *server) sendTime() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b>Текущее время на сервере: %s<b>", time.Now().Format("15:04:05"))
	}
}

func (s *server) respond(w http.ResponseWriter, r *http.Response, code int, data any) {
	data = map[string]any{
		"response": data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
