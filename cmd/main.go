package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jair-anderson-souza/rinha_2024/cmd/rotas"
)

//	var clientes = []cliente = {
//		id: "1", valor: "100000", tipo: ""
//	}
type cliente struct {
	id        int
	valor     int
	tipo      string
	descricao string
}

func main() {
	rotas.Example()
	// numero := testNumbers()
	// fmt.Println(numero)
	api := gin.Default()
	api.Routes()

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
