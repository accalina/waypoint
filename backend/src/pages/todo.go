package pages

import (
	"net/http"
	"waypoint/src/models"
	"waypoint/src/serializers"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	var todos []models.Todo
	// models.DB.Find(&todos)
	models.DB.Where("is_delete = false").Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var input serializers.CreateTodoSerializer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Content: input.Content}
	models.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func RetriveTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input serializers.UpdateTodoSerializer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := map[string]interface{}{
		"content": input.Content,
		"is_done": input.IsDone,
	}

	models.DB.Model(&todo).Updates(&payload)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ? and is_delete = false", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := map[string]interface{}{"is_delete": true}
	models.DB.Model(&todo).Updates(&payload)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
