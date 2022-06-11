package controllers

import (
	"net/http"
	"testgo/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAnimals(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var animals []models.Animal
	db.Find(&animals)

	if len(animals) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "404 Not Found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": animals})
	}
}

func FindAnimal(c *gin.Context) { // Get model if exist
	var animal models.Animal
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&animal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "404 Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": animal})
}

func CreateAnimal(c *gin.Context) {
	var input models.Animal

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	animal := models.Animal{Name: input.Name, Class: input.Class, Legs: input.Legs}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&animal)
	c.JSON(http.StatusOK, gin.H{"data": animal})
}

func UpdateAnimal(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var animal models.Animal

	if err := db.Where("id = ?", c.Param("id")).First(&animal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var input models.Animal

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Animal
	updatedInput.Name = input.Name
	updatedInput.Class = input.Class
	updatedInput.Legs = input.Legs

	db.Model(&animal).Updates(updatedInput)
	c.JSON(http.StatusOK, gin.H{"data": animal})
}

func DeleteAnimal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Animal

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
