package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"libcrossflow/controllers/sharespace"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	r.HandleFunc("/api/file/sharespaces", handleFileShareSpaces)
	r.PathPrefix("/api/file/list/").HandlerFunc(handleFileList)
	r.PathPrefix("/api/file/stat/").HandlerFunc(handleFileStat)
	r.PathPrefix("/api/file/download/").HandlerFunc(handleFileDownload)
	r.HandleFunc("/api/file/upload", handleFileUpload).Methods("POST")
}

func handleFileShareSpaces(w http.ResponseWriter, r *http.Request) {
	shareSpaceNames := sharespace.GetShareSpaceNames()
	res, _ := json.Marshal(shareSpaceNames)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func splitPath(path string) (sharespace.ShareSpace, string, error) {
	var shareSpaceName, relPath string

	split := strings.SplitN(path, "/", 2)
	if len(split) == 0 {
		return sharespace.ShareSpace{}, "", errors.New("Empty path")
	} else if len(split) == 1 {
		shareSpaceName = split[0]
		relPath = ""
	} else {
		shareSpaceName, relPath = split[0], split[1]
	}

	unescapedPath, err := url.QueryUnescape(relPath)
	if err != nil {
		return sharespace.ShareSpace{}, "", err
	}

	shareSpace, ok := sharespace.GetShareSpace(shareSpaceName)
	if !ok {
		return sharespace.ShareSpace{}, "", fmt.Errorf("Sharespace %s does not exist", shareSpaceName)
	}

	return shareSpace, unescapedPath, nil
}

func handleFileList(w http.ResponseWriter, r *http.Request) {
	path, _ := strings.CutPrefix(r.URL.String(), "/api/file/list/")
	shareSpace, relPath, err := splitPath(path)
	if err != nil {
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
	shareSpace, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	fileInfo, err := shareSpace.GetFileStat(relPath)
	if fileInfo.Type == "dir" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	realPath := shareSpace.GetRealPath(relPath)
	http.ServeFile(w, r, realPath)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}
	defer file.Close()

	path := r.FormValue("dirpath")
	shareSpace, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	realPath := shareSpace.GetRealPath(relPath)
	filename := filepath.Join(realPath, handler.Filename)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func handleFileStat(w http.ResponseWriter, r *http.Request) {
	path, _ := strings.CutPrefix(r.URL.String(), "/api/file/stat/")
	shareSpace, relPath, err := splitPath(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

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
