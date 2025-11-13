package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mesh-dell/personal-blog/internal/article"
)

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("web/templates/admin/add.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Template parse error:", err)
		return
	}

	if r.Method == http.MethodPost {
		title := r.PostFormValue("title")
		publishingDate := r.PostFormValue("publishingDate")
		content := r.PostFormValue("content")

		layout := "2006-01-02T15:04"
		publishingDateTime, err := time.Parse(layout, publishingDate)

		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			log.Println("Date parse error:", err)
			return
		}

		err = article.CreateArticle(title, content, publishingDateTime)
		if err != nil {
			http.Error(w, "Failed to create article file", http.StatusInternalServerError)
			log.Println("Create article error", err)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	if err := templ.Execute(w, nil); err != nil {
		log.Println("Template execution error", err)
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("web/templates/admin/dashboard.html")

	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Template parse error:", err)
		return
	}

	data := map[string]any{
		"Articles": []article.Article{},
	}

	articles, err := article.GetArticles()

	if err != nil {
		http.Error(w, "Failed to read articles", http.StatusInternalServerError)
		log.Println("Failed to read articles")
		return
	}

	data["Articles"] = articles
	templ.Execute(w, data)
}

func UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("web/templates/admin/edit.html")

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
		return
	}

	data := struct {
		Id             int
		Title          string
		Content        string
		PublishingDate string
	}{
		Id:             articleData.Id,
		Title:          articleData.Title,
		Content:        articleData.Content,
		PublishingDate: articleData.PublishingDate.Format("2006-01-02T15:04"),
	}
	// get the article
	if r.Method == http.MethodPost {
		title := r.PostFormValue("title")
		publishingDate := r.PostFormValue("publishingDate")
		content := r.PostFormValue("content")

		layout := "2006-01-02T15:04"
		publishingDateTime, err := time.Parse(layout, publishingDate)

		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			log.Println("Date parse error:", err)
			return
		}

		err = article.UpdateArticle(articleData.Id, title, content, publishingDateTime)
		if err != nil {
			http.Error(w, "Unable to update article", http.StatusBadRequest)
			log.Println("unable to update article:", err)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	if err := templ.Execute(w, data); err != nil {
		log.Println("Template execution error", err)
	}
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		articleId := r.PathValue("id")
		articleIdInt, err := strconv.Atoi(articleId)

		if err != nil {
			http.Error(w, "Invalid id type use int", http.StatusBadRequest)
			log.Println("Invalid id type:", err)
			return
		}

		err = article.DeleteArticle(articleIdInt)

		if err != nil {
			http.Error(w, "Unable to delete article", http.StatusBadRequest)
			log.Println("unable to delete article:", err)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
}
