package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBID used to take bid from request
func GetBID(r *http.Request) (int, error) {
	BID := mux.Vars(r)["bid"]

	if BID == "" {
		return 0, errors.New("Bad request: int business ID is required")
	}
	BIDI, err := strconv.Atoi(BID)
	return BIDI, err
}
