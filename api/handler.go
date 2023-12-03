package api

import (
	"github.com/gorilla/mux"
)

func HandleAPI(r *mux.Router) {
	handleFile(r)
}
