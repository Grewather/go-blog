package route

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go-blog/handlers/article"
	"html/template"
	"net/http"
	"os"
)

func Routes() http.Handler {
	godotenv.Load(".env")
	r := chi.NewRouter()
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		title := os.Getenv("BLOG_TITLE")
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		err := tmpl.Execute(w, title)
		if err != nil {
			panic(err)
		}
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		type About struct {
			Title   string
			Content string
		}
		about := About{}
		title := os.Getenv("BLOG_TITLE")
		about.Title = title
		content, err := os.ReadFile("./about/about.md")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		about.Content = string(content)
		tmpl := template.Must(template.ParseFiles("./templates/about.html"))
		err = tmpl.Execute(w, about)
		if err != nil {
			panic(err)
		}
	})

	r.Route("/admin", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			title := os.Getenv("BLOG_TITLE")
			tmpl := template.Must(template.ParseFiles("./templates/admin.html"))
			err := tmpl.Execute(w, title)
			if err != nil {
				panic(err)
			}
		})
		r.Get("/add", func(w http.ResponseWriter, r *http.Request) {
			title := os.Getenv("BLOG_TITLE")
			tmpl := template.Must(template.ParseFiles("./templates/addArt.html"))
			err := tmpl.Execute(w, title)
			if err != nil {
				panic(err)
			}
		})
		r.Get("/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("DDD")
			title := os.Getenv("BLOG_TITLE")
			tmpl := template.Must(template.ParseFiles("./templates/editArt.html"))
			err := tmpl.Execute(w, title)
			if err != nil {
				panic(err)
			}
		})
	})

	r.Get("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		type Article struct {
			Title     string `json:"name"`
			Desc      string `json:"description"`
			Content   string `json:"content"`
			NewField  string `json:"newField"`
			BlogTitle string
		}
		resp, err := http.Get("http://127.0.0.1:3000/api/v1/article/" + chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		var articl Article
		err = json.NewDecoder(resp.Body).Decode(&articl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		articl.BlogTitle = os.Getenv("BLOG_TITLE")

		tmpl := template.Must(template.ParseFiles("./templates/article.html"))
		err = tmpl.Execute(w, articl)
		if err != nil {
			panic(err)
		}

	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/article/{id}", article.GetArticle)
		r.Post("/add", article.AddArticle)
		r.Post("/delete", article.RemoveArticle)
		r.Post("/edit", article.EditArticle)
	})

	return r
}
