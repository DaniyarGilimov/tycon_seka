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

// GetSID used to take sid from request
func GetSID(r *http.Request) (int, error) {
	SID := mux.Vars(r)["sid"]

	if SID == "" {
		return 0, errors.New("Bad request: int sectorId is required")
	}
	SIDI, err := strconv.Atoi(SID)
	return SIDI, err
}
