package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ErrorResponse struct {
	Code int `json:"code"`
	Errors map[string]string `json:"errors"`
}

type TokenResponse struct {
	Code int `json:"code"`
	Token string `json:"token"`
}
type MessageResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}


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


func NewErrorResponse(code int, err error) ErrorResponse {
	var verr validator.ValidationErrors
	res := ErrorResponse{}
	res.Code = code
	res.Errors = make(map[string]string)

	if errors.As(err, &verr) {
		for _, f := range verr {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			res.Errors[f.Field()] = err
		}
	} else {
		res.Errors["error"] = err.Error()
	}

	return res
}