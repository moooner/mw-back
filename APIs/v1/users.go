package v1

import (
	"github.com/gin-gonic/gin"
)

// SignUp
//
// POST /v1/SignUp
//
// @queries
// 	Must have at least one media_uid.
// 	- googleUserId?
// 	- appleUserId?
// @response data
//	- userId
func SignUp(context *gin.Context) {
	// parse queries
	appleUserId := context.Query("appleUserId")
	googleUserId := context.Query("googleUserId")

	// parse response data
	data := map[string]string{
		"userId":       "1",
		"appleUserId":  appleUserId,
		"googleUserId": googleUserId,
	}

	// response 200: success
	context.JSON(200, gin.H{
		"status":  200,
		"message": "Signed up.",
		"data":    data,
	})
}

// SignIn
//
// GET /v1/SignIn
//
// @queries
// 	Must have at least one media_uid.
// 	- googleUserId?
// 	- appleUserId?
// @response data
//	- userId
func SignIn(context *gin.Context) {
	// parse queries
	appleUserId := context.Query("appleUserId")
	googleUserId := context.Query("googleUserId")

	if appleUserId == "" && googleUserId == "" {
		context.JSON(403, gin.H{
			"status":  403,
			"message": "Permission denied.",
			"data":    nil,
		})
		return
	}

	// parse response data
	data := map[string]string{
		"userId": "1",
	}

	// response 200: success
	context.JSON(200, gin.H{
		"status":  200,
		"message": "Signed in.",
		"data":    data,
	})
}
