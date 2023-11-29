package main

import (
	"libcrossflow/api"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() http.Handler {
	var resDir string
	if os.Getenv("APP_ENV") == "dev" {
		resDir = "web/dist"
	} else {
		ex, _ := os.Executable()
		exPath := filepath.Dir(ex)
		resDir = filepath.Join(exPath, "res")
	}

	r := mux.NewRouter()
	api.HandleAPI(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(resDir)))

	if os.Getenv("APP_ENV") == "dev" {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		return handlers.CORS(originsOk, headersOk, methodsOk)(r)
	} else {
		return r
	}
}
