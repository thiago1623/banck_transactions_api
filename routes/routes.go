package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1623/banck_transactions_api/controllers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "hello",
			})
		})
		transactionController := controllers.NewTransactionController()
		v1.POST("/transactions/upload", transactionController.ProcessCSV)
	}

	return router
}
