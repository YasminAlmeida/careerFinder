package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	//init router
	router := gin.Default()

	//init routers
	InitializeRoutes(router)

	router.Run(":8081")
}
