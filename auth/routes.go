package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/auth/register", registerController)

}
