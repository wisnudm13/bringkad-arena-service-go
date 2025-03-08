package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	// register v1 routes
}

func main() {
	router := gin.Default()
	RegisterRouters(router)
}
