package v1

import (
	"bringkad-arena-service-go/routes/v1/admins"
	"bringkad-arena-service-go/routes/v1/facilities"
	"bringkad-arena-service-go/routes/v1/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	users.RegisterUserRoutes(group)
	admins.RegisterAdminRoutes(group)
	facilities.RegisterFacilityRoutes(group)
}
