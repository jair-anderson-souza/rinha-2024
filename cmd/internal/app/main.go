package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.POST("/clientes/:id/transacoes", func(ctx *gin.Context) {
		id := ctx.Param("id")

		ctx.JSON(http.StatusOK, gin.H{
			"message": id,
		})
	})

	api.POST("/clientes", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "message",
		})
	})
	api.Run(":9999")
}
