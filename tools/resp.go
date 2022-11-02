package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Error   = 1
)

type BasePackage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ResponseOk(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = []int{}
	}
	ctx.JSON(http.StatusOK, &BasePackage{Code: Success, Data: data, Msg: "Operation successful."})
}

func ResponseError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, &BasePackage{Code: Error, Data: nil, Msg: err.Error()})
}

func ResponseUnAuthorize(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, &BasePackage{
		Code: Error,
		Data: nil,
		Msg:  err.Error(),
	})
}

func ResponsePaginateData(ctx *gin.Context, total int64, data interface{}) {
	resp := struct {
		TotalCount int64       `json:"totalCount"`
		Records    interface{} `json:"records"`
	}{total, data}
	ctx.JSON(http.StatusOK, &BasePackage{Code: Success, Data: resp, Msg: "success"})
}
