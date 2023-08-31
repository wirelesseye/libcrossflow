package api

import (
	"encoding/json"
	"fmt"
	"libcrossflow/controllers/sharespace"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	r.HandleFunc("/api/sharespaces", handleShareSpaces)
	r.PathPrefix("/api/files/").HandlerFunc(handleFiles)
}

func handleShareSpaces(w http.ResponseWriter, r *http.Request) {
	shareSpaceNames := sharespace.GetShareSpaceNames()
	res, _ := json.Marshal(shareSpaceNames)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func handleFiles(w http.ResponseWriter, r *http.Request) {
	url, _ := strings.CutPrefix(r.URL.String(), "/api/files/")

	split := strings.SplitN(url, "/", 2)
	if len(split) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	shareSpace, path := split[0], split[1]
	fmt.Fprintf(w, "ShareSpace: %v\n", shareSpace)
	fmt.Fprintf(w, "Path: %v\n", path)
}