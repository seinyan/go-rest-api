package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"errors"
)

type Controller interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func GetPathInt(c *gin.Context, name string) (uint64, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	v, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}

	return v, nil
}