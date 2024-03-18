package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trafilea/numberCategorization"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	prefix := "api/trafilea"
	api := engine.Group(prefix)

	api.GET("health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.0.1",
		})
	})

	service := numberCategorization.NewService()
	api.POST("number/:value", numberCategorization.AddNumber(service))
	api.GET("number/:value", numberCategorization.GetNumberType(service))
	api.GET("number/all", numberCategorization.GetAllNumbersType(service))

	return engine
}
