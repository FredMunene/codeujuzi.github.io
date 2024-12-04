package database

import "time"


type User struct {
	UserID int64 `json:"user_id"`
	Email string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Score struct {
	
}