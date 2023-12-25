package routes

import (
	

	"github.com/gin-gonic/gin"
	controller "github.com/vivekjha1213/go-gin-jwt/controllers"
	
)

const PREFIX_URL = "/users"

// AuthRoutes defines the authentication routes
func AuthRoutes(incomingRoutes *gin.Engine) {

    incomingRoutes.POST(PREFIX_URL + "/signup",controller.Signup())
    incomingRoutes.POST(PREFIX_URL + "/login", controller.Login())
}


