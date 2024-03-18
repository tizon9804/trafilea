package numberCategorization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddNumber(service Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("value")
		value, err := strconv.Atoi(param)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		service.AddNumber(value)
		context.JSON(http.StatusNoContent, nil)
	}
}

func GetNumberType(service Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("value")
		value, err := strconv.Atoi(param)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		number, err := service.GetNumber(value)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		context.JSON(http.StatusOK, number)
		return
	}
}

func GetAllNumbersType(service Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		numbers := service.GetAllNumbers()
		context.JSON(http.StatusOK, numbers)
		return
	}
}
