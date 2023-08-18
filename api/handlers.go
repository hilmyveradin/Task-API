package api

import (
	"net/http"
	"task/model"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func getUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func createUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		log.Printf("User data being sent to the database: %+v\n", user)
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind user"})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func getTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []model.Task
		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func createTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind task"})
			return
		}
		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}
