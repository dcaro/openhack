package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	City  string
	Color string
}

//const city = "Amsterdam"
//const color = "coral"

const city = "Redmond"
const color = "gray"

func main() {

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")

	homePageVars := PageVariables{
		city,
		color,
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, homePageVars)
	if err != nil {
		log.Printf("template rendering error: ", err)
	}
}
