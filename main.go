package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record video", Completed: false},
}

// for get request
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}

// for post request
func addTodos(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)

}

func findUsingId(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todoFound, err := findUsingId(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todoFound)
}

func patchTodo(context *gin.Context) {
	id := context.Param("id")
	todoFound, err := findUsingId(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todoFound.Completed = !todoFound.Completed
	context.IndentedJSON(http.StatusOK, todoFound)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", patchTodo)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}
