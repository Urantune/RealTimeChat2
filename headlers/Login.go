package headlers

import (
	"RealTimeChatApplication/services"
	"RealTimeChatApplication/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type UserRam struct {
		Username string `json:"username"`
	}

	userRam := UserRam{}

	err := c.ShouldBindJSON(&userRam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userService := services.NewUserService()
	user, err := userService.GetUserByUserName(userRam.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})

}
