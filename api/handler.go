package api

import (
	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	r.HandleFunc("/api/share_spaces", ShareSpaces)
}
