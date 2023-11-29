package api

import (
	"encoding/json"
	"libcrossflow/controllers/sharespace"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	r.HandleFunc("/api/sharespaces", handleShareSpaces)
	r.PathPrefix("/api/files/").HandlerFunc(handleFiles)
	r.PathPrefix("/api/download/").HandlerFunc(handleDownload)
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
		w.Write([]byte("400 bad request"))
		return
	}

	shareSpaceName, path := split[0], split[1]
	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	files, err := shareSpace.ListFiles(path)
	if err == nil {
		res, _ := json.Marshal(files)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	url, _ := strings.CutPrefix(r.URL.String(), "/api/download/")

	split := strings.SplitN(url, "/", 2)
	if len(split) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	shareSpaceName, path := split[0], split[1]
	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	fileInfo, err := shareSpace.GetFileInfo(path)
	if fileInfo.Type == "dir" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	realPath := shareSpace.GetRealPath(path)
	http.ServeFile(w, r, realPath)
}