package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Setup() {
	getRoutes()
}

func Run() {
	router.Run(":5000")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addTreeRoutes(v1)
}
