package server

import (
	"encoding/json"
	"fmt"
	"lesson11/internal/models"
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
	s.router.HandleFunc("/user", s.treatmentUser())
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

var persons = make([]models.Person, 0)

func (s *server) treatmentUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.respond(w, http.StatusOK, persons)
			return
		}

		if r.Method == http.MethodPost {
			var user models.Person
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				s.error(w, http.StatusBadRequest, fmt.Errorf("invalid JSON format"))
				return
			}

			if user.Age == 0 || user.Name == "" {
				s.error(w, http.StatusBadRequest, fmt.Errorf("invalid fields"))
				return
			}

			persons = append(persons, user)
			s.respond(w, http.StatusOK, true)
			return
		}
	}
}

func (s *server) respond(w http.ResponseWriter, code int, data any) {
	data = map[string]any{
		"response": data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

func (s *server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}
