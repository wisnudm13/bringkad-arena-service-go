package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Offset    *int   `json:"offset,omitempty"`
	Limit     *int   `json:"limit,omitempty"`
	Data      any    `json:"data,omitempty"`
	Error     string `json:"errors,omitempty"`
}

func MakeAPIResponse(c *gin.Context, code int, message string, data interface{}, err string, offset int, limit int, showOffsetLimit bool) {
	// If data is nil, use an empty object {}
	if data == nil && (code == 201 || code == 200) {
		data = map[string]any{} // Empty object instead of null
	}

	// set offset, limit
	var offsetPtr, limitPtr *int
	if showOffsetLimit {
		offsetPtr = &offset
		limitPtr = &limit
	}

	c.JSON(code, BaseResponse{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Offset:    offsetPtr,
		Limit:     limitPtr,
		Data:      data,
		Error:     err,
	})

}

// 200
func OKMessage(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "OK"
	}

	MakeAPIResponse(c, http.StatusOK, msg, data, "", 0, 0, false)
}

// 200 [for list api]
func OKMessageList(c *gin.Context, data interface{}, offset int, limit int) {
	MakeAPIResponse(c, http.StatusOK, "OK", data, "", offset, limit, true)
}

// 201
func CreatedMessage(c *gin.Context, data interface{}) {
	MakeAPIResponse(c, http.StatusCreated, "Data Created", data, "", 0, 0, false)
}

// 400
func BadRequestMessage(c *gin.Context, err string) {
	MakeAPIResponse(c, http.StatusBadRequest, "Bad Request", nil, err, 0, 0, false)
}

// 401
func UnauthorizedMessage(c *gin.Context) {
	MakeAPIResponse(c, http.StatusUnauthorized, "Unauthorized", nil, "", 0, 0, false)
}

// 422
func UnprocessableEntityMessage(c *gin.Context, err string) {
	MakeAPIResponse(c, http.StatusUnprocessableEntity, "Unprocessable Entity", nil, err, 0, 0, false)
}

// 500
func InternalServerErrorMessage(c *gin.Context, err string) {
	MakeAPIResponse(c, http.StatusInternalServerError, "Internal Server Error", nil, err, 0, 0, false)
}
