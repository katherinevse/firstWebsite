package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name                    string
	age                     uint16
	money                   int16
	averageGrade, happiness float64
}

func (u User) getInfoAboutUser() string {
	return fmt.Sprintf("Username is %s. Age is %d and he has %d money ", u.name, u.age, u.money)

}
func homePage(w http.ResponseWriter, r *http.Request) { // первый параметр для того, чтотбы писать что-то на страницу
	//bob := User{"Bob", 25, 500, 5.0, 100}
	fmt.Fprintf(w, "<b>Main TEXT</b>")

	//обработать ошибки нужно!
	tmpl, _:= template.
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Our contacts")
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
