package admins

import (
	"bringkad-arena-service-go/handlers"
	"bringkad-arena-service-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.RouterGroup) {
	admins := r.Group("/admins")

	admins.POST("/", handlers.CreateAdmin)
	admins.GET("/", middleware.ValidateAccessToken("admin"), handlers.GetAdminList)
	admins.GET("/:uuid", middleware.ValidateAccessToken("admin"), handlers.GetAdminDetail)
	admins.DELETE("/:uuid", middleware.ValidateAccessToken("admin"), handlers.DeleteAdmin)
	admins.PUT("/:uuid", middleware.ValidateAccessToken("admin"), handlers.UpdateAdmin)
	admins.POST("/login", handlers.LoginAdmin)
	admins.POST("/login/refresh", handlers.RefreshLoginAdmin)
}
