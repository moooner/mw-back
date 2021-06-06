package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/domain"
	"github.com/moooner/mw-back/domain/model"
)

func Put(context *gin.Context) {
	idx, _ := strconv.ParseUint(string(context.Param("idx")), 10, 64)
	email := context.Query("email")
	appleUserId := context.Query("appleUserId")
	googleUserId := context.Query("googleUserId")

	user := model.User{
		Idx:          idx,
		Email:        email,
		AppleUserId:  appleUserId,
		GoogleUserId: googleUserId,
	}

	response := domain.Db.Model(&user).Updates(&user)

	/// Update user SET deleted_at="date_time",where ID= &user

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
		"message": "Updated.",
		"datum":   domain.Db.First(&user).Row(),
	})
}

func Delete(context *gin.Context) {
	// parse queries
	idx, _ := strconv.ParseUint(string(context.Param("idx")), 10, 64)

	user := model.User{Idx: idx}

	response := domain.Db.Delete(&user)
	/// Update user SET deleted_at="date_time",where ID= &user

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
		"message": "Deleted.",
		"datum":   user,
	})
}
