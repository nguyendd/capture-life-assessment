package v1

import (
	"fmt"
	"net/http"

	"capturelife.assessment.daniel/model"
	"github.com/gin-gonic/gin"
)

var comments = []model.Comment{
	{ID: model.CommentID{Value: 1}, Author: "Elon Musk", Content: "This is true.", BlogID: &model.BlogID{Value: 3}},
	{ID: model.CommentID{Value: 2}, Author: "Daniel Nguyen", Content: "Huh?", BlogID: &model.BlogID{Value: 1}},
	{ID: model.CommentID{Value: 3}, Author: "Thanos", Content: "Is this how it goes?", ParentCommentID: &model.CommentID{Value: 1}},
}

func PostComments(c *gin.Context) {
	var newBlog model.Blog
	if err := c.BindJSON(&newBlog); err != nil {
		fmt.Printf("Unable to add blog. See error: %s", err)
		return
	}
	newBlog.ID = model.BlogID{
		Value: currentBlogId,
	}
	currentBlogId++
	blogs = append(blogs, newBlog)
	c.IndentedJSON(http.StatusCreated, newBlog)
}
