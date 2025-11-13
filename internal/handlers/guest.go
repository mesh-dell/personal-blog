package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/mesh-dell/personal-blog/internal/article"
)

var templHome = template.Must(template.ParseFiles("web/templates/home.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// read articles
	articles, err := article.GetArticles()

	type articleView struct {
		Id             int
		Title          string
		PublishingDate string
	}

	var articleViews []articleView
	for _, article := range articles {
		formattedDate := article.PublishingDate.Format("January 2, 2006")
		articleViews = append(articleViews, articleView{
			Id:             article.Id,
			Title:          article.Title,
			PublishingDate: formattedDate,
		})
	}

	if err != nil {
		http.Error(w, "Failed to read articles", http.StatusInternalServerError)
		log.Println("Failed to read articles")
		return
	}

	data := map[string]any{
		"Articles": articleViews,
	}

	if err := templHome.Execute(w, data); err != nil {
		log.Println("Template execution error", err)
	}
}

func ArticleViewHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("web/templates/article.html")

	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Template parse error:", err)
		return
	}

	articleId := r.PathValue("id")
	articleIdInt, err := strconv.Atoi(articleId)

	if err != nil {
		http.Error(w, "Invalid id type use int", http.StatusBadRequest)
		log.Println("Invalid id type:", err)
		return
	}

	articleData, err := article.FindArticle(articleIdInt)

	if err != nil && err.Error() == "article not found" {
		http.Error(w, "Article not found", http.StatusNotFound)
		log.Println("Article not found:", err)
	}

	formattedDate := articleData.PublishingDate.Format("January 2, 2006")
	articleView := struct {
		Id             int
		Title          string
		Content        template.HTML
		PublishingDate string
	}{
		Id:             articleData.Id,
		Title:          articleData.Title,
		Content:        template.HTML(articleData.Content),
		PublishingDate: formattedDate,
	}

	data := articleView

	if err := templ.Execute(w, data); err != nil {
		log.Println("Template execution error", err)
	}
}
