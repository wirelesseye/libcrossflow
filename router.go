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
	r := mux.NewRouter()
	api.HandleAPI(r)
	r.PathPrefix("/").HandlerFunc(handleWeb)

	if os.Getenv("APP_ENV") == "dev" {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		return handlers.CORS(originsOk, headersOk, methodsOk)(r)
	} else {
		return r
	}
}

func handleWeb(w http.ResponseWriter, r *http.Request) {
	var resDir string
	if os.Getenv("APP_ENV") == "dev" {
		resDir = "web/dist"
	} else {
		ex, _ := os.Executable()
		exPath := filepath.Dir(ex)
		resDir = filepath.Join(exPath, "res")
	}

	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(resDir, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(resDir, "index.html"))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(resDir)).ServeHTTP(w, r)
}
