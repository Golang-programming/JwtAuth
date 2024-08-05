package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/test/gingonic/database"
)

func (repo *database.Repository) AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/register", repo.registerController)

}
