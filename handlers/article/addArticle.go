package article

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-blog/database"
	"net/http"
	"os"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Form)
	code := r.FormValue("code")
	content := r.FormValue("Article")
	title := r.FormValue("Title")
	description := r.FormValue("Description")
	//id := r.FormValue("id")
	//lastDb := database.GetFromDb()
	lastDb, err := database.GetAllFiles()
	id := fmt.Sprint(lastDb[len(lastDb)-1].ID + 1)
	if code != os.Getenv("CODE") {
		http.Error(w, "Invalid code", http.StatusForbidden)
		return
	}
	fmt.Println("Title: ", title)
	fmt.Println("Description: ", description)
	fmt.Println("Content: ", content)
	fmt.Println("Code: ", code)
	//ttl := strings.Map(func(r rune) rune {
	//	switch {
	//	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
	//		return r
	//	case r == ' ':
	//		return '_'
	//	default:
	//		return -1
	//	}
	//}, title)
	err = os.WriteFile("./blog/"+id+".md", []byte(content), 0644)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.AddToDb(&database.File{
		Title:       title,
		Description: description,
		FilePath:    "./blog/" + id + ".md",
	})
}
