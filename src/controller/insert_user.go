package controller

import (
	"github.com/gin-gonic/gin"
	resterr "github.com/rbsilmann/go-lang-crud-mvc/src/configuration/rest_err"
)

func InsertUser(c *gin.Context) {
	err := resterr.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
