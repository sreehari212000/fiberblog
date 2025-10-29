package models

type Post struct {
	POST_ID     int    `json:"post_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Author      int    `json:"author,omitempty"`
}
