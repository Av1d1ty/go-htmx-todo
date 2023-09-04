package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	Content string
	Done    bool
}

func main() {
	fmt.Println("Hello, 世界!")

	r := gin.Default()
	r.Static("/css", "./static/css")

	r.GET("/", getTodos)
	r.POST("/add-todo", addTodo)
	// r.GET("/todo/:id", getTodoById)

	r.LoadHTMLGlob("templates/*")

	r.Run("localhost:8000")
}

func getTodos(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"Todos": todos})
}

func addTodo(c *gin.Context) {
	text := c.PostForm("newtodo")
	todo := ToDo{Content: text, Done: false}

	// Kinda add to DB
	todos = append(todos, todo)

	c.HTML(http.StatusOK, "todo-item", todo)
}

// Kinda DB
var todos = []ToDo{
	{Content: "First todo", Done: false},
	{Content: "Second todo", Done: false},
	{Content: "Third todo", Done: true},
}
