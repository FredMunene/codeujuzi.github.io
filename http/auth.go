package handlers

import (
	"fmt"
	"net/http"

	db "codeujuzi.github.io/database"
	"codeujuzi.github.io/middleware"
	"codeujuzi.github.io/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("unable to parse request")
		}

		w.Header().Set("Content-Type", "application/json")
		email := r.FormValue("email")
		passwd := r.FormValue("password")


		
		// fmt.Fprintf(w, "%v:%v", email, passwd)

		userDbPasswd,userName, err := db.GetUserPasswd(email)
		if err != nil {
			fmt.Printf("user not found%v ", email)
			http.Error(w, "User not found", http.StatusUnauthorized)
			
			return
		}
		if (services.ComparePasswds(userDbPasswd,passwd)){
			fmt.Println("Pasword match")
			tokenString, err  := middleware.CreateToken(userName)
			if err != nil {
				fmt.Println("failed to generate token for user:", userName)
				internalServerErrorHandler(w)
				return
			}
			w.Header().Set("Authorization", "Bearer "+tokenString)
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("Pasword not match ")
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}	
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("unable to parse request")
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		passwd := r.FormValue("password")
		gender := r.FormValue("gender")
		// languages := r.FormValue("language")
		// CONVERT languages into []Languages
		// dob := r.FormValue("dob")
		// CONVERT STRING TO time.Time TODO
		country := r.FormValue("country")

		hashedPasswd, err := services.GetHashedPasswd(passwd)
		if err != nil{
			fmt.Println(err)
			internalServerErrorHandler(w)
			return
		}

		// check if user already exists TODO

		newUser := db.User{
			Name:           name,
			Email:          email,
			HashedPassword: hashedPasswd,
			Country:        country,
			Gender:         gender,
			// Dob: dob,
			// Languges: languages,

		}

		fmt.Println(email, ":", passwd, ":", name)
		// fmt.Fprintf(w, "%v:%v", email, passwd)

		db.AddUser(&newUser)

		CourseHandler(w,r)
	}
}
