package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mesh-dell/personal-blog/internal/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/add", handlers.AuthMiddleware(handlers.AddArticleHandler))
	http.HandleFunc("/admin", handlers.AuthMiddleware(handlers.AdminHandler))
	http.HandleFunc("/delete/{id}", handlers.AuthMiddleware(handlers.DeleteArticleHandler))
	http.HandleFunc("/edit/{id}", handlers.AuthMiddleware(handlers.UpdateArticleHandler))
	http.HandleFunc("/article/{id}", handlers.ArticleViewHandler)
	http.HandleFunc("/search", handlers.ArticleSearchHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
