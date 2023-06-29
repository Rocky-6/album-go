package router

import (
	"album/handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRouter(albumDB *sql.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/healthz", handler.GetHealthz)
	// mux.Handle("/albums", handler.NewAlbumHandler(service.NewAlbumService(albumDB)))
	return router
}
