package src

import (
	"waypoint/src/pages"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {

	todo := r.Group("/todo")
	{
		todo.GET("/", pages.GetTodo)
		todo.GET("/:id", pages.RetriveTodo)
		todo.POST("/", pages.CreateTodo)
		todo.PUT("/:id", pages.UpdateTodo)
		todo.PATCH("/:id", pages.UpdateTodo)
		todo.DELETE("/:id", pages.DeleteTodo)
	}
}
