package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.RouterGroup) {
	users := r.Group("/admins")
	users.POST("/", CreateAdmin)
	users.GET("/", GetAdminList)
	users.GET("/:id", GetAdminDetail)
	users.DELETE("/:id", DeleteAdmin)
	users.PUT("/:id", UpdateAdmin)
}

func CreateAdmin(context *gin.Context) {

}

func DeleteAdmin(context *gin.Context) {

}

func GetAdminDetail(context *gin.Context) {

}

func GetAdminList(context *gin.Context) {

}

func UpdateAdmin(context *gin.Context) {

}
