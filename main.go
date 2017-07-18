package main

import (
	"fmt"
	"net/http"
	"html/template"
	"crypto/rand"
	"./models"
)

var posts map[string]*models.Post

func generateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "write", post)
}

func saveHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	title := r.FormValue("title")
	desc := r.FormValue("desc")

	var post *models.Post
	if id !="" {
		post = posts[id]
		post.Title = title
		post.Desc = desc
	} else {
		id = generateId()
		post := models.NewPost(id, title, desc)
		posts[id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}


func main() {
	fmt.Println("Listening on port: 3000")
	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.ListenAndServe(":3000", nil)
}