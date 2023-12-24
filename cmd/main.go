package main

import (
	"github.com/ShababKarim/go-ddd/internal/core"
	"github.com/ShababKarim/go-ddd/internal/domain/adapters"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config := core.GetAppConfig()

	log.Printf("Started server on port %d", config.Port)
	log.Fatal("Error starting app with cause",
		http.ListenAndServe(getAddress(config.Port), adapters.NewAppMux()))
}

func getAddress(port int) string {
	return ":" + strconv.Itoa(port)
}
