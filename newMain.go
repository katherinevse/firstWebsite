package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/homePage")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "homePage", nil) // блок зтмл файла
}

func handleRequest() {
	http.HandleFunc("/", homePage) // отслеживание гдавной страницы home_page- какой-то метод, который будет вызываться при переходе на главную страницу
	http.HandleFunc("/contact/", contactPage)
	err := http.ListenAndServe(":8080", nil) // первый параметр порт какой-то
	if err != nil {
		log.Fatalln("error to launch server:", err)
	}
}

func main() {
	handleRequest()
}
