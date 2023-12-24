package adapters

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.Marshal("OK")
	if err != nil {
		log.Println("Error marshalling healthcheck response")
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Println("Error writing healthcheck response")
	}
}

func NewAppMux() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthcheck", healthcheckHandler)

	return r
}
