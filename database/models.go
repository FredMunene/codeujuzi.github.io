package database

import "time"


type User struct {
	UserID int64 `json:"user_id" gorm:"primaryKey; autoIncrement"`
	Name string `json:"user_name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique"`
	HashedPassword string `json:"hashed_password" gorm:"not null"`
	Dob time.Time `json:"dob" gorm:"type:date"`
	Role Role `json:"role" gorm:"default:visitor"`
	Country string `json:"country" gorm:"size:50"`
	Gender string `json:"gender" gorm:"size:10"`
	CreatedAt time.Time `json:"created_at" gorm:"autocreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autocreateTime"`
	Languges []UserLanguage `json:"languages" gorm:"foreignKey:UserID"`
}

type Score struct {
	
}

type UserLanguage struct {
	UserLanguageID int64 `json:"user_language_id" gorm:"PrimaryKey; autoIncrement"`
	UserID int64 	`json:"user_id" gorm:"not null;index"`
	Language string `json:"language" gorm:"size:50;not null"`
}

type Role string

const (
	Admin    Role = "admin"
	Visitor     Role = "visitor"
	Moderator Role = "moderator"
	Student  Role = "student" // Default role
)