package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbsilmann/go-lang-crud-mvc/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// fmt.Println(os.Getenv("TEST"))
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
