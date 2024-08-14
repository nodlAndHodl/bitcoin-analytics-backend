package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nodlandhodl/bitcoin-analytics-backend/docs"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/controllers"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/service/bitcoind"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()
	bitcoindService, err := bitcoind.NewBitcoindService()
	if err != nil {
		panic(err)
	}
	bitcoindController := controllers.NewBitcoindController(bitcoindService)
	route.GET("/api/v1/blockhash/:blockHeight", bitcoindController.GetBlockHash)
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8081/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run("localhost:8081")
}
