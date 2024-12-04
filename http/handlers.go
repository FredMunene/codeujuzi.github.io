package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"codeujuzi.github.io/services"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request){

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}
	

}

func SigninHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("unable to parse request")
		}

		email := r.FormValue("email")
		passwd := r.FormValue("password")

		fmt.Println(email,":",passwd)
		fmt.Fprintf(w,"%v:%v",email,passwd)

		hashedPasswd := services.GetHashedPasswd(passwd)

		fmt.Printf("%x",hashedPasswd)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("/template/register.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}

}

func LessonsHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("template/lessons.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}

}
func ProfileHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("template/profile.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}

}

func CourseHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("template/course-outline.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Index.html parsing error: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// internalServerErrorHandler()
		log.Println("Index.html template parsing error: ", err)
		return
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w,r," http://localhost:8000/",http.StatusSeeOther)

}

// func TranslateHandler(w http.ResponseWriter, r *http.Request){

// }

func internalServerErrorHandler(w http.ResponseWriter){
	renderErrorPage(w,http.StatusInternalServerError,"Permission denied")
}
	
func wrongMethodErrorHandler(w http.ResponseWriter){
	renderErrorPage(w,http.StatusMethodNotAllowed,"Permission denied")
}
func pageNotFoundErrorHandler(w http.ResponseWriter){
	renderErrorPage(w,http.StatusNotFound,"The page you are looking for does not exist")
}
func badRequestHandler(w http.ResponseWriter) {
	renderErrorPage(w, http.StatusBadRequest, " Try the home page")
}
func renderErrorPage(w http.ResponseWriter, statusCode int, message string){
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles("template/error.html")
	if err != nil {
		log.Println("Error page parsing error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}
	data := struct {
		Message string
	}{
		Message: message,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error: page execution:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}