package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start")
	StartServer()
}
func StartServer() {

	fmt.Println("Сервер запущен на порту 3000...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/htmlTemplate/assets"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/htmlTemplate/assets/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/htmlTemplate/assets/js"))))
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("web/htmlTemplate/images"))))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/algorithm.html", algorithm)
	mux.HandleFunc("/automata_theory.html", automataTheory)

	mux.HandleFunc("/algorithm_lab1.html", algorithm_lab1)
	mux.HandleFunc("/automata_theory_lab1.html", automata_theory_lab1)

	http.ListenAndServe(":"+port, mux)

}

// Обработчик страницы теория алгоритмов.
func algorithm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/algorithm.html" {
		http.NotFound(w, r)

		return
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"web/htmlTemplate/main.html",
		"web/htmlTemplate/algorithm.html",
		"web/htmlTemplate/footer.html",
		"web/htmlTemplate/header.html",
		"web/htmlTemplate/nav.html",
		"web/htmlTemplate/banner.html",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func automataTheory(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"web/htmlTemplate/main.html",
		"web/htmlTemplate/automata_theory.html",
		"web/htmlTemplate/footer.html",
		"web/htmlTemplate/header.html",
		"web/htmlTemplate/nav.html",
		"web/htmlTemplate/banner.html",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"web/htmlTemplate/main.html",
		"web/htmlTemplate/index.html",
		"web/htmlTemplate/footer.html",
		"web/htmlTemplate/header.html",
		"web/htmlTemplate/banner.html",
		"web/htmlTemplate/nav.html",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func algorithm_lab1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/algorithm_lab1.html" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("nen")

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"web/htmlTemplate/main.html",
		"web/htmlTemplate/algorithm_lab1.html",
		"web/htmlTemplate/footer.html",
		"web/htmlTemplate/header.html",
		"web/htmlTemplate/banner.html",
		"web/htmlTemplate/nav.html",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func automata_theory_lab1(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/automata_theory_lab1.html" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("nen")

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"web/htmlTemplate/main.html",
		"web/htmlTemplate/automata_theory_lab1.html",
		"web/htmlTemplate/footer.html",
		"web/htmlTemplate/header.html",
		"web/htmlTemplate/banner.html",
		"web/htmlTemplate/nav.html",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
