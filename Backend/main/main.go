package main

import (
	injector "github.com/Irictm/AutoFixPortfolio/Backend/main/Injector"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	injector.InjectDependencies(router)
	router.Run("localhost:8080")
}
