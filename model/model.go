package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UID      uint   `gorm:"primaryKey" json:"uid"`
	FullName string `gorm:"column:full_name" json:"fullName"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID" json:"tasks"`
}

type Task struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	DeadlineDate time.Time `gorm:"column:deadline_date" json:"deadlineDate"`
	Status       uint      `gorm:"column:status" json:"status"`
	UserID       uint      `gorm:"column:user_id" json:"userID"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Task{})
}
