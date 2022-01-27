create database go_course;

use go_course;

create table posts (id int auto_increment primary key, title varchar(50) not null, body text);

// Implementação antes de usar o gorilla mux
func main() {

	// stmt, err := db.Prepare("insert into posts(title,body) values (?, ?)")
	// checkError(err)

	// _, err = stmt.Exec("My first Post", "My first content")
	// checkError(err)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	post := Post{Id: 1, Title: "Unamed Post", Body: "No Content"}

	if title := r.FormValue("title"); title != "" {
		post.Title = title
	}
	t := template.Must(template.ParseFiles("templates/index.html"))

	if err := t.ExecuteTemplate(w, "index.html", items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	

	fmt.Println(http.ListenAndServe(":8080", r))
}