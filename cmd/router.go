package main

import (
	"github.com/gin-gonic/gin"
	"scanner-service/configs"
)

/**
  @author:pandi
  @date:2022-11-03
  @note:
**/
func loadRouters() *gin.Engine {
	gin.SetMode(configs.APP.Mode)
	router := gin.Default()
	return router
}
