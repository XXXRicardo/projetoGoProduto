package controllers

import (
	"api/database"
	"api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Mercados(c *gin.Context) {
	fmt.Println("entrou no mostrar com id")
	c.JSON(200, gin.H{
		"value": "meu pau",
	})
}

func ShowMercado(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("entrou no mostrar com id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID TEM Q SER INTEIRO",
		})
	}

	db := database.GetDatabase()

	var mercado models.Mercado
	err = db.First(&mercado, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "linguiça não encontrada " + err.Error(),
		})
		return
	}
	c.JSON(200, mercado)

}

func CreateProduto(c *gin.Context) {
	db := database.GetDatabase()

	var mercado models.Mercado

	err := c.ShouldBindJSON(&mercado)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "linguiça não criada " + err.Error(),
		})
		return
	}

	err = db.Create(&mercado).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "linguiça não encontrada " + err.Error(),
		})
		return
	}
	c.JSON(200, mercado)
}

func ShowProduto(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Mercado
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func UpdateProduto(c *gin.Context) {
	db := database.GetDatabase()

	var mercado models.Mercado

	err := c.ShouldBindJSON(&mercado)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "linguiça não criada " + err.Error(),
		})
		return
	}

	err = db.Save(&mercado).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "linguiça não atualizada " + err.Error(),
		})
		return
	}
	c.JSON(200, mercado)
}

func DeleteMercado(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID TEM Q SER INTEIRO",
		})
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Mercado{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "salsicha nao deletada" + err.Error(),
		})
		return
	}
	c.Status(204)
}
