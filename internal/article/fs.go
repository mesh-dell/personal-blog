package article

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func SaveArticle(article Article) error {
	fileName := fmt.Sprintf("article-%d.json", article.Id)
	fileDir := GetFilepath()
	filePath := filepath.Join(fileDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file")
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(article)
	if err != nil {
		fmt.Println("Error encoding json")
		return err
	}
	return nil
}

func ReadArticles() ([]Article, error) {
	filePath := GetFilepath()
	files, err := os.ReadDir(filePath)

	if err != nil {
		return []Article{}, err
	}

	var articles []Article

	for _, file := range files {
		var article Article
		fileName := path.Join(filePath, file.Name())
		f, err := os.Open(fileName)

		if err != nil {
			fmt.Println("Error opening:", fileName)
			continue
		}

		err = json.NewDecoder(f).Decode(&article)
		f.Close()
		if err != nil {
			fmt.Println("Error decoding:", fileName)
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func WriteArticle(article Article) error {
	fileDir := GetFilepath()
	fileName := fmt.Sprintf("article-%d.json", article.Id)
	filePath := filepath.Join(fileDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(article)
	if err != nil {
		fmt.Println("Error encoding data")
		return err
	}
	return nil
}

func RemoveArticle(id int) error {
	articleName := fmt.Sprintf("article-%d.json", id)
	filepath := path.Join(GetFilepath(), articleName)
	return os.Remove(filepath)
}
func GetFilepath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	filePath := path.Join(cwd, "data")
	return filePath
}
