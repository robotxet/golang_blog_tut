package models

//Post struct for blog post
type Post struct {
	ID              string
	Title           string
	ContentHTML     string
	ContentMarkdown string
}

//NewPost creates new blog post
func NewPost(id, title, contentHTML, contentMarkdown string) *Post {
	return &Post{id, title, contentHTML, contentMarkdown}
}
