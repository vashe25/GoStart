package models

type Post struct {
	Id string
	Title string
	Desc string
}

func NewPost(id, title, desc string) *Post {
	return &Post{id, title, desc}
}