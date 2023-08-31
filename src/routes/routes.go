package routes

import (
	"rinhabackendleo/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/pessoas", controllers.CreatePessoa)
	r.GET("/pessoas", controllers.GetPessoas)
	r.GET("/pessoas/:id", controllers.GetPessoaById)
	r.GET("/contagem-pessoas", controllers.GetPessoasCount)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}
