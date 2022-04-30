package controllers

import (
	"G014/finalassignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTodo godoc
// @Summary Get Todo
// @Description get All Todo List
// @ID get-Todo
// @Produce  json
// @Success 200 "success"
// @Router /todos [get]
func GetTodo(c *gin.Context) {
	var todos []models.Todo
	models.Db.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// GetTodo godoc
// @Summary Get Todo
// @Description create new Todo
// @ID create-Todo
// @Accept  json
// @Produce  json
// @Param requestbody body models.TodoAdd true "request body json"
// @Success 200 "success"
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var create models.TodoAdd
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Task: create.Task, Completed: create.Completed}
	models.Db.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GetTodoId godoc
// @Summary Get Todo by Id
// @Description find todo by id
// @ID get-Todo-Id
// @Accept  json
// @Produce  json
// @Param id path int true "Todos Id"
// @Success 200 {object} models.Todo
// @Failure 400 {string} Token "Todo not found"
// @Router /todos/{id} [get]
func GetTodoId(c *gin.Context) {
	var getId models.Todo
	if err := models.Db.Where("id = ?", c.Param("id")).First(&getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": getId})

}

// DeleteTodoId godoc
// @Summary Delete Todo by Id
// @Description delete todo by id
// @ID delete-Todo-Id
// @Accept  json
// @Produce  json
// @Param id path int true "Todos Id"
// @Success 200 {object} models.Todo
// @Failure 400 {string} Token "Todo not found"
// @Router /todos/{id} [delete]
func DeleteTodoId(c *gin.Context) {
	var getId models.Todo
	if err := models.Db.Where("id = ?", c.Param("id")).First(&getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
	}

	models.Db.Delete(&getId)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// UpdateTodoId godoc
// @Summary Update Todo by Id
// @Description updating todo by id
// @ID update-Todo-Id
// @Accept  json
// @Produce  json
// @Param id path int true "Todos Id"
// @Param requestbody body models.TodoAdd true "request body json"
// @Success 200 {object} models.Todo
// @Failure 400 {string} Token "Todo not found"
// @Router /todos/{id} [put]
func UpdateTodoId(c *gin.Context) {
	var getId models.Todo
	if err := models.Db.Where("id = ?", c.Param("id")).First(&getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
	}

	c.BindJSON(&getId)
	models.Db.Save(getId)

	c.JSON(http.StatusOK, gin.H{"data": getId})

}
