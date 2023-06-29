package handler

import (
	"album/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, &model.HealthzResponse{Message: "OK"})
}
