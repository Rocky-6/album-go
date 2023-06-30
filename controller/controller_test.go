package controller_test

import (
	"album/controller/router"
	"album/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthz(t *testing.T) {
	albumDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer albumDB.Close()

	router := router.NewRouter(albumDB)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	correctResponse := "{\"message\":\"OK\"}"
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, correctResponse, w.Body.String())
}
