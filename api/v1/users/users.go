package users

import (
	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/domain"
	"github.com/moooner/mw-back/domain/model"
)

func Post(context *gin.Context) {
	// parse queries
	appleUserId := context.Query("appleUserId")
	googleUserId := context.Query("googleUserId")
	email := context.Query("email")

	user := model.User{GoogleUserId: googleUserId, AppleUserId: appleUserId, Email: email}

	response := domain.Db.Create(&user)

	if response.Error != nil {
		context.JSON(400, gin.H{
			"status":  400,
			"message": "Bad request.",
			"datum":   response.Error.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"status":  200,
		"message": "Signed up.",
		"datum":   user,
	})
}

func Get(context *gin.Context) {
	// parse queries
	appleUserId := context.Query("appleUserId")
	googleUserId := context.Query("googleUserId")
	email := context.Query("email")

	user := model.User{GoogleUserId: googleUserId, AppleUserId: appleUserId, Email: email}

	response := domain.Db.First(&user)

	if response.Error != nil {
		context.JSON(400, gin.H{
			"status":  400,
			"message": "Bad request.",
			"datum":   response.Error.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"status":  200,
		"message": "Signed up.",
		"datum":   user,
	})
}
