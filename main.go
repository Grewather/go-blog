package main

import (
	"go-blog/database"
	"go-blog/route"
	"net/http"
)

func main() {
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	database.Init()
	r := route.Routes()
	http.ListenAndServe(":3000", r)
}
