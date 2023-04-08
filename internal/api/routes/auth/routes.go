package auth

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	get(router)
	post(router)
	put(router)
}

func get(router *gin.RouterGroup) {
}

func post(router *gin.RouterGroup) {

}

func put(router *gin.RouterGroup) {
}
