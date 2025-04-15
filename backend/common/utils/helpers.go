package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idString, exists := vars["id"]
	if !exists {
		return 0, http.ErrMissingFile
	}
	return strconv.Atoi(idString)
}
