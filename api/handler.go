package api

import (
	"encoding/json"
	"errors"
	"libcrossflow/controllers/sharespace"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	r.HandleFunc("/api/file/sharespaces", handleFileShareSpaces)
	r.PathPrefix("/api/file/list/").HandlerFunc(handleFileList)
	r.PathPrefix("/api/file/stat/").HandlerFunc(handleFileStat)
	r.PathPrefix("/api/file/download/").HandlerFunc(handleFileDownload)
}

func handleFileShareSpaces(w http.ResponseWriter, r *http.Request) {
	shareSpaceNames := sharespace.GetShareSpaceNames()
	res, _ := json.Marshal(shareSpaceNames)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func splitPath(path string) (string, string, error) {
	var shareSpaceName, relPath string

	split := strings.SplitN(path, "/", 2)
	if len(split) == 0 {
		return "", "", errors.New("Empty path")
	} else if len(split) == 1 {
		shareSpaceName = split[0]
		relPath = ""
	} else {
		shareSpaceName, relPath = split[0], split[1]
	}

	unescapedPath, err := url.QueryUnescape(relPath)
	if err != nil {
		return "", "", err
	}

	return shareSpaceName, unescapedPath, nil
}

func handleFileList(w http.ResponseWriter, r *http.Request) {
	path, _ := strings.CutPrefix(r.URL.String(), "/api/file/list/")
	shareSpaceName, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	files, err := shareSpace.ListFiles(relPath)
	if err == nil {
		res, _ := json.Marshal(files)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
	}
}

func handleFileDownload(w http.ResponseWriter, r *http.Request) {
	path, _ := strings.CutPrefix(r.URL.String(), "/api/file/download/")
	shareSpaceName, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	unescapedPath, err := url.QueryUnescape(relPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}
	path = unescapedPath

	fileInfo, err := shareSpace.GetFileStat(relPath)
	if fileInfo.Type == "dir" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	realPath := shareSpace.GetRealPath(relPath)
	http.ServeFile(w, r, realPath)
}

func handleFileStat(w http.ResponseWriter, r *http.Request) {
	path, _ := strings.CutPrefix(r.URL.String(), "/api/file/stat/")
	shareSpaceName, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	unescapedPath, err := url.QueryUnescape(relPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}
	path = unescapedPath

	fileInfo, err := shareSpace.GetFileStat(relPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	res, _ := json.Marshal(fileInfo)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
