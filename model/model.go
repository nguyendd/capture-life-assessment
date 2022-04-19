package model

import "time"

type BlogID struct {
	Value int64 `json:"value"`
}

type Blog struct {
	ID        BlogID    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"time"`
}

type UpdateBlogInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CommentID struct {
	Value int64 `json:"value"`
}

type Comment struct {
	ID              CommentID  `json:"id"`
	Author          string     `json:"author"`
	Content         string     `json:"content"`
	BlogID          *BlogID    `json:"blog_id"`
	ParentCommentID *CommentID `json:"parent_comment_id"`
}

type UpdateCommentInput struct {
	Content string `json:"content"`
}
