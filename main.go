package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/a-h/templ"
	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
)

type Post struct {
	Title string    `yaml:"title"`
	Slug  string    `yaml:"slug"`
	Date  time.Time `yaml:"date"`
	Draft bool      `yaml:boolean`
	Body  string
}

func main() {

	router := http.NewServeMux()

	router.Handle("/", templ.Handler(indexPage(getFile("about-me.md"))))
	router.HandleFunc("/posts/{slug}", routePosts)
	router.Handle("/posts/", templ.Handler(blogPage(getBlogPosts())))
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/feed/", serveFeed)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func routePosts(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	var posts = getBlogPosts()
	for _, post := range posts {
		if post.Slug == slug && post.Draft != true {
			contentPage(post).Render(r.Context(), w)
		}
	}
}

func getBlogPosts() []Post {
	ext := ".md"
	var posts []Post

	filepath.WalkDir("./content/blog", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ext {
			var post = Post{}

			content, err := os.ReadFile(path)

			rest, err := frontmatter.Parse(bytes.NewReader(content), &post)

			if err != nil {
				log.Println("Could not read file!")
			}

			post.Body = mdToHtml(rest)
			posts = append(posts, post)
		}
		return nil
	})

	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Date.Before(posts[i].Date)
	})

	for _, post := range posts {
		log.Println(post.Date)
	}

	return posts
}

func getFile(file string) string {
	var output string

	filepath.WalkDir("./content", func(path string, d fs.DirEntry, err error) error {
		if matched, err := filepath.Match(file, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			content, err := os.ReadFile(path)

			if err != nil {
				log.Println("Could not read file!")
			}

			output = mdToHtml(content)
		}
		return nil
	})
	return output
}

func mdToHtml(content []byte) string {
	var htmlOutput bytes.Buffer

	if err := goldmark.Convert(content, &htmlOutput); err != nil {
		log.Println("Converting from markdown to html failed!")
	}

	return htmlOutput.String()
}

func serveFeed(w http.ResponseWriter, r *http.Request) {
	rss, err := createFeed(getBlogPosts()).ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, rss)
}
