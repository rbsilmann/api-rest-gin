package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbsilmann/go-lang-crud-mvc/src/controller/routes"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Test package godotenv using variable TEST in .env
	// fmt.Println(os.Getenv("TEST"))

	// Init router using gin default (logger and recovery middlewares)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	// Syntax #1
	// if err := router.Run(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

	// Syntax #2
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
