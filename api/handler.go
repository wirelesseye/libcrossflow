package api

import (
	"encoding/json"
	"net/http"
	"os"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dirname, _ := os.UserHomeDir()
	files, _ := os.ReadDir(dirname)

	var filenames []string
	for _, e := range files {
		filenames = append(filenames, e.Name())
	}

	res, _ := json.Marshal(filenames)
	w.Write(res)
}
