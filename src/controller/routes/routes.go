package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rbsilmann/go-lang-crud-mvc/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user", controller.GetUser)
	r.GET("/user/:id", controller.GetUserById)
	r.GET("/user/email/:email", controller.GetUserByEmail)
	r.POST("/user", controller.InsertUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}
