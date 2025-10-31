package models

type Post struct {
	POST_ID     int    `json:"post_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Author      int    `json:"author,omitempty"`
}

type Comment struct {
	Comment_id int    `json:"comment_id,omitempty"`
	Post_id    int    `json:"post_id,omitempty"`
	Author_id  int    `json:"author_id,omitempty"`
	Text       string `json:"text,omitempty"`
}
