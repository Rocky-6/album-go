package handler_test

import (
	"album/db"
	"album/handler/router"
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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
