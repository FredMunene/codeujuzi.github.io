package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var errConn error

	DB, errConn = gorm.Open(sqlite.Open("codeujuzi.db"), &gorm.Config{})
	if errConn != nil {
		fmt.Println("failed to connect to database: ", errConn)
		return
	}

	// Migrate Schema(create table 'users')
	err := DB.AutoMigrate(&User{}, &UserLanguage{})
	if err != nil {
		fmt.Println("failed to migrate schema: ", err)
		return
	}

	fmt.Println("Db connected and schema migrated successfully")
}

func AddUser(user *User) error {

	var existingUser User

	result := DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return fmt.Errorf("email %s is already in use", user.Email)
	}

	// initialize dialect
	result = DB.Create(user)
	if result.Error != nil {
		 return fmt.Errorf("failed to add user: %w", result.Error)
	}
	fmt.Println("User added successfully", user.Email)
	return nil
}

func GetUserPasswd(email string) (string,string, error) {

	var user User
	result := DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return "","",fmt.Errorf("user not found: %w", result.Error)
   }

   

	return string(user.HashedPassword), string(user.Name), nil
}