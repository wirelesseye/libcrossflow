package api

import (
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	GetShareSpaces(w, r)
}
