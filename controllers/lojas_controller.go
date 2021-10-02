package controllers

import (
	"go/database"
	"go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowLojas(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id not is a integer",
		})

		return
	}

	db := database.GetDatabase()

	var loja models.Loja
	err = db.First(&loja, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find Loja" + err.Error(),
		})

		return
	}

	c.JSON(200, loja)
}

func CreateLoja(c *gin.Context) {
	db := database.GetDatabase()

	var loja models.Loja

	err := c.ShouldBindJSON(&loja)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON" + err.Error(),
		})

		return
	}

	err = db.Create(&loja).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create Loja" + err.Error(),
		})

		return
	}

	c.JSON(200, loja)
}

func ListLojas(c *gin.Context) {
	db := database.GetDatabase()

	var lojas []models.Loja
	err := db.Find(&lojas).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list the Lojas" + err.Error(),
		})

		return
	}

	c.JSON(200, lojas)
}

func UpdateLoja(c *gin.Context) {
	db := database.GetDatabase()

	var loja models.Loja

	err := c.ShouldBindJSON(&loja)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON" + err.Error(),
		})

		return
	}

	err = db.Save(&loja).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update Loja" + err.Error(),
		})

		return
	}

	c.JSON(200, loja)
}

func DeleteLoja(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id not is a integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Loja{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot Delete Loja" + err.Error(),
		})

		return
	}

	c.Status(204)
}
