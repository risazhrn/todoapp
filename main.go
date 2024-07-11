package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID 		string `json:"id"`
	Task 	string `json:"task`
	Status 	string `json:"status"`
}

var toDoList = []ToDo{
	{ID: "1", Task: "Belajar Golang", Status: "Pending"},
	{ID: "2", Task: "Membuat aplikasi To-Do List", Status: "In Progress"},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getToDos)
	router.POST("/todos", createToDo)
	router.PUT("/todos/:id", updateToDo)
	router.DELETE("/todos/:id", deleteToDo)

	router.Run("localhost:8080")
}

func getToDos(c *gin.Context){
	c.IndentedJSON(http.StatusOK, toDoList)
}

func createToDo(c *gin.Context){
	var newToDo ToDo

	if err := c.BindJSON(&newToDo); err != nil {
		return
	}

	toDoList = append(toDoList, newToDo)
	c.IndentedJSON(http.StatusCreated, newToDo)
}


func updateToDo(c *gin.Context){
	id := c.Param("id")
	var updateToDo ToDo

	if err := c.BindJSON(&updateToDo); err != nil {
		return
	}

	for i, t := range toDoList {
		if t.ID == id {
			toDoList[i].Task = updateToDo.Task
			toDoList[i].Status = updateToDo.Status
			c.IndentedJSON(http.StatusOK, toDoList[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})

}

func deleteToDo(c *gin.Context){
	id := c.Param("id")

	for i, t := range toDoList {
		if t.ID == id {
			toDoList = append(toDoList[:i],toDoList[i+1:]... )
			c.IndentedJSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})

}
