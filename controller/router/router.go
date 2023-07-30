package router

import (
	"album/controller"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRouter(albumDB *sql.DB) *gin.Engine {
	con := controller.NewController(albumDB)

	router := gin.Default()
	router.GET("/healthz", con.GetHealthz)
	router.POST("/albums", con.AddAlbum)
	router.GET("/albums", con.GetAlbums)
	return router
}
