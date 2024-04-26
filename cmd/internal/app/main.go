package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	clientes []cliente = {
//		id: "1", valor: "100000", tipo: ""
//	}
type cliente struct {
	id        int
	valor     int
	tipo      string
	descricao string
}

func main() {

	api := gin.Default()

	api.POST("/clientes/:id/extrato", func(ctx *gin.Context) {
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
	api.Run(":8080")
}
