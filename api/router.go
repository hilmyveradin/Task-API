package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping)

	// User routes
	r.GET("/users", getUsers(db))
	r.POST("/users", createUser(db))

	// Task routes
	r.GET("/tasks", getTasks(db))
	r.POST("/tasks", createTask(db))

	return r
}
