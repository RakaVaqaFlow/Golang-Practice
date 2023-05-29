package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	// Многоуровневый путь
	http.HandleFunc("/", rootHandler)

	// В случае повторной регистрации - паника
	// http.HandleFunc("/", rootHandler)

	// Фиксированный путь
	http.HandleFunc("/home", homeHandler)

	// Многоуровневый путь
	http.HandleFunc("/article/", articleHandler)

	// Фиксированный путь
	http.HandleFunc("/article/hello", articleHelloHandler)

	// Классический роут пакета http не поддерживает регулярные выражения
	http.HandleFunc("/article/*/hello", articleHelloHandler)

	// Обработчик редиректа
	http.HandleFunc("/redirect", http.RedirectHandler("http://localhost:9000/home", http.StatusPermanentRedirect).ServeHTTP)
	http.HandleFunc("/redirect/2", redirectArticleHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// ------------------------------------------------------------------------------------------------------------
// Обработчики
// ------------------------------------------------------------------------------------------------------------

func rootHandler(_ http.ResponseWriter, req *http.Request) {
	// if req.URL.Path != "/" {
	// 	fmt.Println("unsupported path")
	// }

	fmt.Println("root")
}

func homeHandler(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("home")
}

func articleHandler(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("article")
}

func articleHelloHandler(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("article hello")
}

func redirectArticleHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:9000/article", http.StatusPermanentRedirect)
}
