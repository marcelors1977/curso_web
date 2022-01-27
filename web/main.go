package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, _ = sql.Open("mysql", "root:root@(localhost:33006)/go_course?charset=utf8")

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/{id}/view", ViewHandler)
	r.HandleFunc("/", HomeHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func GetPostsById(id string) Post {

	row := db.QueryRow("select * from posts where id = ?", id)

	posts := Post{}

	row.Scan(&posts.Id, &posts.Title, &posts.Body)

	return posts

}

func ListPosts() []Post {

	rows, err := db.Query("select * from posts")
	checkError(err)

	items := []Post{} // slice de posts

	for rows.Next() {

		posts := Post{}
		rows.Scan(&posts.Id, &posts.Title, &posts.Body)
		items = append(items, posts)
	}

	return items

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))

	if err := t.ExecuteTemplate(w, "layout.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/view.html"))

	if err := t.ExecuteTemplate(w, "layout.html", GetPostsById(vars["id"])); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
