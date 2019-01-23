package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tongzcomradez/go-gorm/handler"
)

func main() {
	mux := gin.Default()

	mux.GET("/users", handler.GetUser)

	mux.Run(":3030")
}
