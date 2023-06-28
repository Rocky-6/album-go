package router

import (
	"album/handler"
	"album/service"
	"database/sql"
	"net/http"
)

func NewRouter(albumDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/healthz", handler.NewHealthzHandler())
	mux.Handle("/albums", handler.NewAlbumHandler(service.NewAlbumService(albumDB)))
	return mux
}
