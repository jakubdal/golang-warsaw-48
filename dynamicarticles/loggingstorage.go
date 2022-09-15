package main

import "fmt"

var _ ArticleStorage = LoggingStorage{}

type LoggingStorage struct {
	// if you only write ArticleStorage you will "embed" the interface
	// and our interface implementation highlight will not work anymore
	storage ArticleStorage
}

func WithCallLogging(articleStorage ArticleStorage) ArticleStorage {
	return LoggingStorage{storage: articleStorage}
}

func (ls LoggingStorage) Article(id int) (Article, error) {
	fmt.Printf("Article(%v)\n", id)
	return ls.storage.Article(id)
}

func (ls LoggingStorage) Articles() ([]Article, error) {
	fmt.Printf("Articles()\n")
	return ls.storage.Articles()
}
