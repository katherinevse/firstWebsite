package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/homePage.html", "tmpl/header.html", "tmpl/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "homePage", nil) // блок зтмл файла
}

func purchasePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/purchasePage.html", "tmpl/header.html", "tmpl/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "purchasePage", nil) // блок зтмл файла
}

// обработка данных,отображение шаблона не нужно
// получение данных из формы
func saveArticle(w http.ResponseWriter, r *http.Request) {
	//title := r.FormValue("title")       //input поле
	//anons := r.FormValue("anona")       //anons поле
	//fulltext := r.FormValue("fulltext") //input поле

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.:8080/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//установка в базу данных
	insert, err := db.Query("INSERT INTO `articles`(`title`,`anons`, `fulltext`) VALUES (`Ukraine`,1)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", homePage) // отслеживание гдавной страницы home_page- какой-то метод, который будет вызываться при переходе на главную страницу
	http.HandleFunc("/purchase", purchasePage)
	http.HandleFunc("/saveArticle", saveArticle)

	err := http.ListenAndServe(":8080", nil) // первый параметр порт какой-то
	if err != nil {
		log.Fatalln("error to launch server:", err)
	}
}

func main() {
	handleRequest()
}
