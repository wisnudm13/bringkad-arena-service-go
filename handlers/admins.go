package handlers

import (
	"bringkad-arena-service-go/databases"
	"bringkad-arena-service-go/dto"
	"bringkad-arena-service-go/models"
	"bringkad-arena-service-go/utils"

	"github.com/gin-gonic/gin"
)

func CreateAdmin(context *gin.Context) {
	// req body validation
	var rq dto.RegisterAdminRequest

	if utils.RequestValidator(context, &rq) != nil {
		return
	}

	var admins []models.Admin
	// data validation in db
	if err := databases.DB.Where("username = ? OR email = ?", rq.Username, rq.Email).First(&admins).Error; err == nil {
		utils.BadRequestMessage(context, "Username or email already used by other admins")
		return
	}

	// create admin
	admin := models.Admin{
		Username: rq.Username,
		Email:    rq.Email,
		Password: rq.Password,
	}

	if err := databases.DB.Create(&admin).Error; err != nil {
		utils.InternalServerErrorMessage(context, err.Error())
		return
	}

	utils.CreatedMessage(context, nil)
}

func DeleteAdmin(context *gin.Context) {

}

func GetAdminDetail(context *gin.Context) {
	// get uuid
	uuid := context.Param("uuid")

	// data validation
	var admin models.Admin
	if err := databases.DB.Where("uuid = ?", uuid).Where("deleted_at IS NULL").First(&admin).Error; err != nil {
		utils.BadRequestMessage(context, "Data not found")
		return
	}

	utils.OKMessage(context, dto.GetAdminDetailResponse{
		Username: admin.Username,
		Email:    admin.Email,
		IsActive: admin.IsActive,
	}, "OK")
}

func GetAdminList(context *gin.Context) {

}

func UpdateAdmin(context *gin.Context) {

}

func LoginAdmin(context *gin.Context) {
	// request body validation
	var rq dto.LoginAdminRequest

	if utils.RequestValidator(context, &rq) != nil {
		return
	}

	// data validation
	var admin models.Admin
	if err := databases.DB.Where("username = ?", rq.Username).First(&admin).Error; err != nil {
		utils.BadRequestMessage(context, "Incorrect username or password")
		return
	}

	if passwordValid := utils.CheckHashPassword(rq.Password, admin.Password); !passwordValid {
		utils.BadRequestMessage(context, "Incorrect username or password")
	}

	// generate access token and set to cookie
	accessToken := utils.GenerateAccessToken("admin", admin.Uuid, admin.ID)
	context.SetCookie("access_token", accessToken, 1800, "/", "", false, true) // TODO: set domain and secure https

	// generate refresh access token and set cookie
	refreshToken := utils.GenerateRefreshAccessToken("admin", admin.Uuid, admin.ID)
	context.SetCookie("refresh_access_token", refreshToken, 604800, "/", "", false, true) // TODO: set domain and secure https

	utils.OKMessage(context, nil, "Login Successful")
}

// RefreshLoginAdmin refreshes the access token for admin
//
// This function is used to generate a new access token when the existing one expires.
func RefreshLoginAdmin(context *gin.Context) {

}
