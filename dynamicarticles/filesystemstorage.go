package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TYPEASSERTIONSTART OMIT
var _ ArticleStorage = FilesystemStorage{}

type FilesystemStorage struct{}

// TYPEASSERTIONEND OMIT

// LISTARTICLESSTART OMIT
func (fs FilesystemStorage) Articles() ([]Article, error) {
	fileInfos, err := ioutil.ReadDir("articles")
	if err != nil {
		return nil, fmt.Errorf("read articles directory: %w", err)
	}
	var articles []Article
	for _, fileInfo := range fileInfos {
		filename := fileInfo.Name()
		extension := filepath.Ext(fmt.Sprintf("articles/%s", filename))
		if extension != ".html" {
			fmt.Printf("[filename=%s] invalid extension: %v\n", filename, extension)
			continue
		}
		filenameWithoutExtension := strings.TrimSuffix(filename, extension)
		articleID, err := strconv.Atoi(filenameWithoutExtension)
		if err != nil {
			return nil, fmt.Errorf("[filename=%s] convert to articleID: %w", filename, err)
		}
		article, err := fs.Article(articleID) // HL
		if err != nil {
			return nil, fmt.Errorf("[articleID=%v] get article: %w", articleID, err)
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// LISTARTICLESEND OMIT

// GETARTICLESTART OMIT
func (fs FilesystemStorage) Article(id int) (Article, error) {
	filename := fmt.Sprintf("%v.html", id)
	filepath := fmt.Sprintf("articles/%s", filename)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return Article{}, fmt.Errorf("[filepath=%s] read file: %w", filepath, err)
	}
	article, err := fs.constructArticle(id, b) // HL
	if err != nil {
		return Article{}, fmt.Errorf("[filepath=%s] read file: %w", filepath, err)
	}
	return article, nil
}

// GETARTICLEEND OMIT

// CONSTRUCTARTICLESTART OMIT
func (fs FilesystemStorage) constructArticle(id int, content []byte) (Article, error) {
	stringContent := string(content)

	return Article{
		ID:      id,
		Title:   stringBetween(stringContent, "<h1>", "</h1>"),
		Lead:    stringBetween(stringContent, "<p>", "</p>"),
		Content: stringContent,
	}, nil
}

// CONSTRUCTARTICLEEND OMIT

// STRINGBETWEENSTART OMIT
func stringBetween(content, start, end string) string {
	startIndex := strings.Index(content, start)
	if startIndex != -1 {
		startIndex += len(start)
	}
	endIndex := strings.Index(content, end)
	if startIndex == -1 || endIndex == -1 || endIndex < startIndex {
		return ""
	}
	return content[startIndex:endIndex]
}

// STRINGBETWEENEND OMIT
