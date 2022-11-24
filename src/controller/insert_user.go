package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	resterr "github.com/rbsilmann/go-lang-crud-mvc/src/configuration/rest_err"
	"github.com/rbsilmann/go-lang-crud-mvc/src/model/request"
)

func InsertUser(c *gin.Context) {
	// err := resterr.NewBadRequestError("Você chamou a rota de forma errada")
	// c.JSON(err.Code, err)
	var userRequest request.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("incorrect fields, error: %s", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
}
