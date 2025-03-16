package facilities

import (
	"github.com/gin-gonic/gin"
)

func RegisterFacilityRoutes(r *gin.RouterGroup) {
	facilities := r.Group("/facilities")

	facilities.POST("/", CreateFacility)
	facilities.GET("/", GetFacilityList)
	facilities.GET("/:id", GetFacilityDetail)
	facilities.DELETE("/:id", DeleteFacility)
	facilities.PUT("/:id", UpdateFacility)
}

func CreateFacility(context *gin.Context) {
}

func DeleteFacility(context *gin.Context) {

}

func GetFacilityDetail(context *gin.Context) {

}

func GetFacilityList(context *gin.Context) {

}

func UpdateFacility(context *gin.Context) {

}
