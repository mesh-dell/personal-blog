package article

import (
	"fmt"
	"strings"
	"time"
)

type Article struct {
	Id             int
	Title          string
	PublishingDate time.Time
	Content        string
}

func NewArticle(id int, title, content string, publishingDate time.Time) *Article {
	newArticle := Article{
		Id:             id,
		Title:          title,
		PublishingDate: publishingDate,
		Content:        content,
	}
	return &newArticle
}

func GetArticles() ([]Article, error) {
	articles, err := ReadArticles()
	if err != nil {
		return []Article{}, err
	}
	return articles, nil
}

func CreateArticle(title, content string, publishingDate time.Time) error {
	// read all articles
	articles, err := ReadArticles()
	if err != nil {
		return err
	}

	var newArticleId int
	if len(articles) == 0 {
		newArticleId = 1
	} else {
		newArticleId = articles[len(articles)-1].Id + 1
	}

	newArticle := NewArticle(newArticleId, title, content, publishingDate)
	return SaveArticle(*newArticle)
}

func UpdateArticle(id int, title, content string, publishingDate time.Time) error {
	articles, err := ReadArticles()
	if err != nil {
		return err
	}

	found := false
	for _, article := range articles {
		if article.Id == id {
			found = true
			article.Title = title
			article.Content = content
			article.PublishingDate = publishingDate
			err := WriteArticle(article)
			if err != nil {
				return err
			}
		}
	}

	if !found {
		return fmt.Errorf("article not found")
	}

	fmt.Println("Article updated successfully")
	return nil
}

func DeleteArticle(id int) error {
	articles, err := ReadArticles()
	if err != nil {
		return err
	}

	found := false
	for _, article := range articles {
		if article.Id == id {
			found = true
			err := RemoveArticle(id)
			if err != nil {
				fmt.Println("Error deleting article")
				return err
			}
		}
	}

	if !found {
		return fmt.Errorf("article not found")
	}

	fmt.Println("Article deleted successfully")
	return nil
}

func FindArticle(id int) (Article, error) {
	articles, err := ReadArticles()
	if err != nil {
		return Article{}, err
	}

	for _, article := range articles {
		if article.Id == id {
			return article, nil
		}
	}
	return Article{}, fmt.Errorf("article not found")
}

func SearchArticle(titleQuery string) ([]Article, error) {
	articles, err := ReadArticles()
	if err != nil {
		return nil, err
	}

	var filteredArticles []Article
	// filter by title
	for _, article := range articles {
		if strings.Contains(article.Title, titleQuery) {
			filteredArticles = append(filteredArticles, article)
		}
	}

	return filteredArticles, nil
}
