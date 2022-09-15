package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// LANDINGSTART OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("articles/landing.html")
		if err != nil {
			fmt.Println("os.ReadFile(\"articles/landing.html\"):", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})
	// LANDINGEND OMIT
	// 1START OMIT
	http.HandleFunc("/blog/posts/1/try-writing-a-website-by-hand", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("articles/1.html")
		if err != nil {
			fmt.Println("os.ReadFile(\"articles/1.html\"):", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})
	// 1END OMIT
	// 2START OMIT
	http.HandleFunc("/blog/posts/2/when-to-use-third-party-packages", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("articles/2.html")
		if err != nil {
			fmt.Println("os.ReadFile(\"articles/2.html\"):", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})
	// 2END OMIT
	// 3START OMIT
	http.HandleFunc("/blog/posts/3/you-dont-need-an-operations-team-to-deploy-your-code", func(w http.ResponseWriter, r *http.Request) { // HLchanges
		content, err := os.ReadFile("articles/3.html") // HLchanges
		if err != nil {
			fmt.Println("os.ReadFile(\"articles/3.html\"):", err) // HLchanges
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})
	// 3END OMIT
	http.ListenAndServe(":8080", nil)
}
