package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name                    string
	Age                     uint16
	Money                   int16
	AverageGrade, Happiness float64
	Hobbies                 []string
}

func (u User) getInfoAboutUser() string {
	return fmt.Sprintf("Username is %s. Age is %d and he has %d money ", u.Name, u.Age, u.Money)

}
func homePage(w http.ResponseWriter, r *http.Request) { // первый параметр для того, чтотбы писать что-то на страницу
	bob := User{"Bob", 25, 500, 5.0, 100, []string{"Football", "Skate", "Dance"}}
	//fmt.Fprintf(w, "<b>Main TEXT</b>")

	//обработать ошибки нужно!
	tmpl, _ := template.ParseFiles("templates/homePage.html")
	tmpl.Execute(w, bob)
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
