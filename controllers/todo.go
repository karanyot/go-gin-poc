package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo GET S",
	})
}

func CreateATodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Create",
	})
}

func GetATodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo GET",
	})
}

func UpdateATodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Update",
	})
}

func DeleteATodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Delete",
	})
}
