package handler

import (
	"album/service"
	"net/http"
)

type AlbumHandler struct {
	svc *service.AlbumService
}

func NewAlbumHandler(svc *service.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		svc: svc,
	}
}

func (h *AlbumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
