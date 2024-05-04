package article

import (
	"fmt"
	"go-blog/database"
	"net/http"
	"os"
	"strconv"
)

func RemoveArticle(w http.ResponseWriter, r *http.Request) {
	idstr := r.FormValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	if code != os.Getenv("CODE") {
		http.Error(w, "Invalid code", http.StatusForbidden)
		return
	}
	res, err := database.GetFromDb(id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("43")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.Remove(res.FilePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("22")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = database.RemoveFromDb(id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("1")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
