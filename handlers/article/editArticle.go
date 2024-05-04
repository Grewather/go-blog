package article

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-blog/database"
	"net/http"
	"os"
)

func EditArticle(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error4143123: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)

	code := r.FormValue("code")
	content := r.FormValue("Article")
	title := r.FormValue("Title")
	description := r.FormValue("Description")

	if code != os.Getenv("CODE") {
		http.Error(w, "Invalid code", http.StatusForbidden)
		return
	}

	id := r.FormValue("id")

	file, err := os.OpenFile("./blog/"+id+".md", os.O_RDWR, 0644)
	if err != nil {
		fmt.Print("2Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Print("3Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.WriteFile("./blog/"+id+".md", []byte(content), 0644)
	if err != nil {
		fmt.Print("4Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.EditDb(&database.File{
		Title:       title,
		Description: description,
		FilePath:    "./blog/" + id + ".md",
	})
}
