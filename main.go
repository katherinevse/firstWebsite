package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                    string
	age                     uint16
	money                   int16
	averageGrade, happiness float64
}

func home_page(w http.ResponseWriter, r *http.Request) { // первый параметр для того, чтотбы писать что-то на страницу
	bob := User{"Bob", 25, 500, 5.0, 100}
	fmt.Fprintf(w, "Username is "+bob.name)
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Our contacts")
}

func handle_request() {
	http.HandleFunc("/", home_page) // отслеживание гдавной страницы home_page- какой-то метод, который будет вызываться при переходе на главную страницу
	http.HandleFunc("/contact/", contact_page)
	http.ListenAndServe(":8080", nil) // первый параметр порт какой-то
}
func main() {
	handle_request()
}
