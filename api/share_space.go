package api

import (
	"libcrossflow/config"
	"encoding/json"
	"net/http"
)

func GetShareSpaces(w http.ResponseWriter, r *http.Request) {
	config := config.GetConfig()
	shareSpaces := config.GetShareSpaces()

	res, _ := json.Marshal(shareSpaces)
	w.Write(res)
}
