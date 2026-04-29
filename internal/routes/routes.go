package routes

import (
	"lentera/internal/controller"
	"lentera/internal/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, repo repository.PgRepo) {
	c := controller.Controller{
		Db: repo,
	}

	router.POST("/attendance/check-in", c.CheckIn)
	router.POST("/attendance/check-out", c.CheckOut)
	router.GET("/attendance", nil)
}
