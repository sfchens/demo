package controller

import (
	"demo/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Home struct {
}

func (c *Home) Home(ctx *gin.Context) {
	fmt.Printf("111111")
	response.Ok(ctx)
}
