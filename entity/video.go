package entity

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname,omitempty" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname,omitempty" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age,omitempty" binding:"gte=1,lte=100"`
	Email     string `json:"email,omitempty" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title,omitempty" binding:"min=2,max=100" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json:"description,omitempty" binding:"required,max=500" gorm:"type:varchar(500)"`
	URL         string    `json:"url,omitempty" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author,omitempty" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
