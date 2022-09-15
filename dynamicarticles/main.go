package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

// ARTICLESTART OMIT
type Article struct {
	ID      int
	Title   string
	Lead    string
	Content string
}

// ARTICLEEND OMIT

// INTERFACESTART OMIT
type ArticleGetter interface {
	Article(id int) (Article, error)
}

type ArticleLister interface {
	Articles() ([]Article, error)
}

type CombinedArticleStorage interface {
	ArticleGetter
	ArticleLister
}

type ExplicitArticleStorage interface {
	Article(id int) (Article, error)
	Articles() ([]Article, error)
}

// two interfaces above are equivalent
var _ ExplicitArticleStorage = (CombinedArticleStorage)(nil)

// type alias for less reading
type ArticleStorage CombinedArticleStorage

// INTERFACEEND OMIT

// MAINSTART OMIT
func main() {
	var articleStorage ArticleStorage
	articleStorage = FilesystemStorage{}
	articleStorage = WithCallLogging(articleStorage)
	r := chi.NewRouter()
	r.Get("/", ListArticlesHandler(articleStorage))                       // HL
	r.Get("/blog/posts/{articleID}/*", GetArticleHandler(articleStorage)) // HL
	http.ListenAndServe(":8080", r)
}

// MAINEND OMIT

// GETARTICLEHANDLERSTART OMIT
func GetArticleHandler(articleStorage ArticleStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "articleID") // HL
		id, err := strconv.Atoi(idParam)        // HL
		if err != nil {
			errMsg := fmt.Sprintf("[idParam=%v] GetArticle: invalid article id parameter: %v", idParam, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			fmt.Println(errMsg)
			return
		}
		article, err := articleStorage.Article(id) // HL
		if err != nil {
			errMsg := fmt.Sprintf("[id=%v] GetArticle: %v", id, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMsg))
			fmt.Println(errMsg)
			return
		}
		w.Write([]byte(article.Content))
	}
}

// GETARTICLEHANDLEREND OMIT

// TEMPLATESTART OMIT
const pageTemplate = `<!DOCTYPE html>
<html>
    <head>
        <title>daliga.pl</title>
    </head>
    <body>
        <div id="container">
            <div id="header">
                <a href="/"><span id="author">Kuba Daliga</span></a>
            </div>
            <div id="content">%s</div>
            </div>
        </div>
    </body>
</html>`

// TEMPLATEEND OMIT

// TEASERSTART OMIT
const teaserTemplate = `<div class="post">
<div class="title"><a href="%s">%s</a></div>
<div class="lead">%s</div>
<a href="%s"><div class="button cta">Read more...</div></a>
</div>`

// TEASEREND OMIT

// LISTARTICLESSTART OMIT
func ListArticlesHandler(articleStorage ArticleStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleStorage.Articles()
		if err != nil {
			errMsg := fmt.Sprintf("list articles: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			fmt.Println(errMsg)
			return
		}
		content := ""
		for _, article := range articles {
			url := articleURL(article)
			teaser := fmt.Sprintf(teaserTemplate,
				url, article.Title,
				article.Lead,
				url)
			content += teaser
		}
		landingContent := fmt.Sprintf(pageTemplate, content)
		w.Write([]byte(landingContent))
	}
}

// LISTARTICLESEND OMIT

// ARTICLEURLSTART OMIT
var matchSpecialCharacters = regexp.MustCompile("[^a-z0-9 ]+")

func articleURL(article Article) string {
	formattedTitle := strings.ToLower(article.Title)
	formattedTitle = matchSpecialCharacters.ReplaceAllString(formattedTitle, "")
	formattedTitle = strings.Replace(formattedTitle, " ", "-", -1)
	return fmt.Sprintf("/blog/posts/%d/%s", article.ID, formattedTitle)
}

// ARTICLEURLEND OMIT
