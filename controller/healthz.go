package controller

import (
	"album/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (con *Controller) GetHealthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &model.HealthzResponse{Message: "OK"})
}
