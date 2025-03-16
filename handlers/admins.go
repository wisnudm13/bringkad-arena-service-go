package handlers

import (
	"bringkad-arena-service-go/databases"
	"bringkad-arena-service-go/dto"
	"bringkad-arena-service-go/models"
	"bringkad-arena-service-go/utils"
	"reflect"
	"strconv"

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
	// get uuid from param
	uuid := context.Param("uuid")

	var admin models.Admin
	if err := databases.DB.Where("uuid = ?", uuid).Where("deleted_at IS NULL").First(&admin).Error; err != nil {
		utils.BadRequestMessage(context, "Data not found")
		return
	}

	// soft delete admin
	databases.DB.Delete(&admin)

	utils.OKMessage(context, nil, "Data deleted successfully")

}

func GetAdminDetail(context *gin.Context) {
	// get uuid from param
	uuid := context.Param("uuid")

	// data validation
	var admin models.Admin
	if err := databases.DB.Where("uuid = ?", uuid).Where("deleted_at IS NULL").First(&admin).Error; err != nil {
		utils.BadRequestMessage(context, "Data not found")
		return
	}

	utils.OKMessage(context, dto.GetAdminResponse{
		Uuid:     admin.Uuid,
		Username: admin.Username,
		Email:    admin.Email,
		IsActive: admin.IsActive,
	}, "OK")
}

func GetAdminList(context *gin.Context) {
	var admins []models.Admin
	var adminsResponse []dto.GetAdminResponse

	// get offset
	offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	// get limit
	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	// get search
	search := context.Query("search")
	query := databases.DB.Model(&models.Admin{})

	if search != "" {
		query = query.Where("username LIKE ? OR email LIKE ? AND delete_at IS NULL", "%"+search+"%", "%"+search+"%")
	}

	query.Limit(limit).Offset(offset).Find(&admins)

	for _, admins := range admins {
		adminsResponse = append(adminsResponse, dto.GetAdminResponse{
			Uuid:     admins.Uuid,
			Username: admins.Username,
			Email:    admins.Email,
			IsActive: admins.IsActive,
		})
	}

	utils.OKMessageList(context, adminsResponse, offset, limit)
}

func UpdateAdmin(context *gin.Context) {
	// get uuid from param
	uuid := context.Param("uuid")

	var rq dto.UpdateAdminRequest
	var admin models.Admin

	if utils.RequestValidator(context, &rq) != nil {
		return
	}

	if err := databases.DB.Where("uuid = ?", uuid).Where("deleted_at IS NULL").First(&admin).Error; err != nil {
		utils.BadRequestMessage(context, "Data not found")
		return
	}

	if rq.Password != "" {
		if rq.ConfirmPassword == "" {
			utils.UnprocessableEntityMessage(context, "Confirm Password is required")
			return
		}

		if rq.Password != rq.ConfirmPassword {
			utils.UnprocessableEntityMessage(context, "Password and Confirm Password do not match")
			return
		}

		// hash password
		hashedPassword, _ := utils.HashPassword(rq.Password)
		rq.Password = hashedPassword

	}

	// Use reflection to update only non-empty fields
	requestValue := reflect.ValueOf(rq)
	requestType := reflect.TypeOf(rq)

	for i := range requestType.NumField() {
		field := requestType.Field(i)
		fieldName := field.Tag.Get("json")
		value := requestValue.Field(i)

		// Check if the column exists in the database
		if value.IsValid() && databases.DB.Migrator().HasColumn(&models.Admin{}, fieldName) {
			var finalValue interface{}
			if value.Kind() == reflect.Bool {
				finalValue = value.Bool()

			} else if !value.IsZero() {
				finalValue = value.Interface()

			} else {
				continue
			}

			databases.DB.Model(&admin).Update(fieldName, finalValue)

		}
	}

	utils.OKMessage(context, nil, "Data successfully updated")
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
