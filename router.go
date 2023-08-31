package main

import (
	"libcrossflow/api"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	resDir, success := os.LookupEnv("RES_PATH")
	if !success {
		ex, _ := os.Executable()
		exPath := filepath.Dir(ex)
		resDir = filepath.Join(exPath, "res")
	}

	r := mux.NewRouter()
	api.HandleAPI(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(resDir)))

	return r
}
