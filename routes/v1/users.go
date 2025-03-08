package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.POST("/", CreateUser)
	users.GET("/", GetUserList)
	users.GET("/:id", GetUserDetail)
	users.DELETE("/:id", DeleteUser)
	users.PUT("/:id", UpdateUser)
}

func CreateUser(context *gin.Context) {

}

func DeleteUser(context *gin.Context) {

}

func GetUserDetail(context *gin.Context) {

}

func GetUserList(context *gin.Context) {

}

func UpdateUser(context *gin.Context) {

}
