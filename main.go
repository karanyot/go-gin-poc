package main

import (
	"github.com/karanyot/go-gin-poc/routes"

	"github.com/Depado/ginprom"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	prom := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(prom.Instrument())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router := routes.TodoV1Router()

	endless.ListenAndServe(":4242", router)
}
