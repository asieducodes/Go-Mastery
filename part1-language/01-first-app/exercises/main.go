package main

// Getting started with the first chapter
// Page 05 -> as reference
import (
	"fmt"
	"html/template"
	"net/http"
)

// Defining a Data Type and a collection
type Rsvp struct {
	Name, Email, Phone string
	willAttend         bool
}

var templates = make(map[string]*template.Template, 3)

// Creating a slice
var responses = make([]*Rsvp, 0, 10)

func loadTemplate() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// Creating the HTTP Handlers and Server
func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}
func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		templates["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{}, // An instance of the required data format
		})
	case http.MethodPost:
		request.ParseForm()
		responseData := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			willAttend: request.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)
			if responseData.willAttend {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}

	}
}
func main() {
	//fmt.Println("Hello, new backend engineer")
	loadTemplate()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
