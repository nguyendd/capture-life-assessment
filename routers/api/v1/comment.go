package v1

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetComments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, comments)
}

func GetCommentsByBlogID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	blogComments := make([]model.Comment, 0)

	for _, comment := range comments {
		if comment.ID.Value == id {
			blogComments = append(blogComments, comment)
		}
	}

	c.IndentedJSON(http.StatusOK, blogComments)
}

func UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	var input model.UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("an error occurred while parsing the blog update: %s", err))
		return
	}

	for i, _ := range comments {
		if comments[i].ID.Value == id {
			comments[i].Content = input.Content
			c.IndentedJSON(http.StatusOK, comments[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, fmt.Sprintf("comment with id %s did not exist", idStr))
}

func getChildren(id int64) []model.Comment {
	children := make([]model.Comment, 0)
	for _, comment := range comments {
		if comment.ParentCommentID != nil && comment.ParentCommentID.Value == id {
			children = append(children, comment)
			children = append(children, getChildren(comment.ID.Value)...)
		}
	}

	return children
}

func DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	allChildren := getChildren(id)
	tempComments := comments[:0]

	for _, comment := range comments {
		if comment.ID.Value == id {
			continue
		}
		skip := false
		for _, child := range allChildren {
			if comment.ID.Value == child.ID.Value {
				break
			}
		}

		if !skip {
			tempComments = append(tempComments, comment)
		}
	}

	if len(tempComments) == len(comments) {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("no-op, a comment with id %s does not exist", idStr))
		return
	}

	comments = tempComments
	c.IndentedJSON(http.StatusOK, "")
}
