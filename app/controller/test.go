package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Test struct {
}

func (c *Test) Test(ctx *gin.Context) {
	fmt.Printf("Test")
}
