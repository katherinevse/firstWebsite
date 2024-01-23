package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/homePage.html", "tmpl/header.html", "tmpl/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "homePage", nil) // блок html файла
}

func purchasePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/purchasePage.html", "tmpl/header.html", "tmpl/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "purchasePage", nil) // блок зтмл файла
}
func errorForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/errorForm.html", "tmpl/header.html", "tmpl/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "errorForm", nil) // блок зтмл файла
}

// обработка данных,отображение шаблона не нужно
// получение данных из формы
func saveArticle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")       //input поле
	surname := r.FormValue("surname") //anons поле
	age := r.FormValue("age")         //input поле

	if name == "" || surname == "" || age == "" {
		errorForm(w, r)
	} else {
		ageNumber, err := strconv.Atoi(age) //приведение к инту
		// ошибка со стороны пользователя, поэтому BadRequest
		if err != nil {
			fmt.Println("Ошиба пользователя, ввел неправильно что-то")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/site")
		if err != nil {
			fmt.Println("Ошибка при открытии соединения с базой данных")
			return
		}
		if err = db.Ping(); err != nil {
			fmt.Println("Ошибка при открытии соединения с базой данных PING")
			return
		}
		fmt.Println("Успешное подключение к базе данных")
		defer db.Close()

		q := fmt.Sprintf("INSERT INTO site(name, surname, age) VALUES ('%s', '%s', '%d');", name, surname, ageNumber)

		// Исправлен SQL-запрос
		insert, err := db.Query(q)
		if err != nil {
			fmt.Println("Ошибка выполнения скрипта:", err)
			return
		}
		defer insert.Close()

		//переадрессация на другую старницу после успешного заполнения формы
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
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
