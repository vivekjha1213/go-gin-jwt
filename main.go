package main

import (
	"os"
	"github.com/gin-gonic/gin"

	routes "github.com/vivekjha1213/go-gin-jwt/routes"

)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Initialize the Gin router
	router := gin.New()

	// Start the log.
	router.Use(gin.Logger())

	// Set up routes for authentication and user-related functionalities
	routes.AuthRoutes(router)
	routes.UserRoutes(router)


	router.GET("/api-1",func( c *gin.Context){

		c.JSON(200,gin.H{"success":"Access granted api-1"})
		
	})

	router.GET("/api-2",func( c *gin.Context){

		c.JSON(200,gin.H{"success":"Access granted api-2"})
		
	})

    router.Run(":" + port)
}
