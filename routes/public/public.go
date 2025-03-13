package public

import "github.com/gin-gonic/gin"

func InitStorage(r *gin.RouterGroup) {
	r.Static("/storage/gift_card", "./storage/gift_card")
}
