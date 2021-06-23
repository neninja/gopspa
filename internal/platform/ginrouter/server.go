package ginrouter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gopspa/docs"
)

// https://github.com/swaggo/swag#general-api-info
// https://github.com/swaggo/swag#api-operation

// @title Gopspa - Biblioteca
// @version 1.0
// @description Prova de conceito (poc) da stack Go + Postgresql + SPA (React) de um gerenciamento de biblioteca

// @license.name UNLICENSE
// @license.url https://unlicense.org/
func Run() {
	r := gin.Default()

	// TODO desenvolvimento frontend
	r.Use(CORS())

	api := r.Group("/api")
	{
		Livros(api)
		Alunos(api)
		Emprestimos(api)

		r.GET("/x", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": true,
			})
		})

		r.GET("/x/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("o livro %s", c.Param("id")),
			})
		})
	}

	url := ginSwagger.URL("http://localhost:3030/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(spa.Middleware("/", "web/static"))

	r.Run(":3030")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}

		c.Next()
	}
}
