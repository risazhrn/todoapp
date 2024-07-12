package controllers

import (
	"net/http"
	"todoapp/config"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func GetToDos(c *gin.Context) {
	var toDos []models.ToDo
	config.DB.Find(&toDos)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "ToDos retrieved successfully",
		"data":    toDos,
	})
}

func CreateToDo(c *gin.Context) {
	var newToDo models.ToDo
	if err := c.BindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newToDo)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "ToDo created successfully",
		"data":    newToDo,
	})
}

func UpdateToDo(c *gin.Context) {
	var updatedToDo models.ToDo
	id := c.Param("id")
	if err := c.BindJSON(&updatedToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&models.ToDo{}).Where("id = ?", id).Updates(updatedToDo)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "ToDo updated successfully",
		"data":    updatedToDo,
	})
}

func DeleteToDo(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&models.ToDo{})
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
}
