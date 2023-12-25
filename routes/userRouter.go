


package routes

import(
	controller "github.com/vivekjha1213/go-gin-jwt/controllers"
	"github.com/vivekjha1213/go-gin-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}




