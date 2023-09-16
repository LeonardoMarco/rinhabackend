package routes

import (
	"rinhabackendleo/src/config"
	"rinhabackendleo/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	db := config.ConnectDatabase()
	client := config.ConnectRedis()

	r.Use(func(c *gin.Context) {
		c.Set("DB_CONTEXT", db)
		c.Set("CACHE_CONTEXT", client)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/pessoas", controllers.CreatePessoa)
	r.GET("/pessoas", controllers.GetPessoas)
	r.GET("/pessoas/:id", controllers.GetPessoaById)
	r.GET("/contagem-pessoas", controllers.GetPessoasCount)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}
