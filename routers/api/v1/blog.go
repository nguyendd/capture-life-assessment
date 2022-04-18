package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"capturelife.assessment.daniel/model"
	"github.com/gin-gonic/gin"
)

var currentBlogId = int64(4)

var blogs = []model.Blog{
	{ID: model.BlogID{Value: 1}, Title: "Eastern Conference NBA Playoffs", Author: "Bill Simmons", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", Timestamp: time.Now()},
	{ID: model.BlogID{Value: 2}, Title: "Western Conference NBA Playoffs", Author: "Zach Lowe", Content: "Mauris nec mattis est. Ut vel tincidunt nisi. Aenean consectetur sapien non bibendum viverra.", Timestamp: time.Now()},
	{ID: model.BlogID{Value: 3}, Title: "Monday", Author: "Daniel Nguyen", Content: "Dread it, run from it, Monday arrives all the same.", Timestamp: time.Now()},
}

func PostBlogs(c *gin.Context) {
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

func GetBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs)
}

func GetBlogByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	for _, b := range blogs {
		if b.ID.Value == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("could not find a blog with id %s", idStr))
}

func UpdateBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	var input model.UpdateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("an error occurred while parsing the blog update: %s", err))
		return
	}

	for i, _ := range blogs {
		if blogs[i].ID.Value == id {
			blogs[i].Title = input.Title
			blogs[i].Content = input.Content
			c.IndentedJSON(http.StatusOK, blogs[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, fmt.Sprintf("blog with id %s did not exist", idStr))
}

func DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("failed to parse id, see error: %s", err)
		c.IndentedJSON(http.StatusBadRequest, "the id cannot be parsed as a number")
		return
	}

	tempBlogs := blogs[:0]
	tempComments := comments[:0]

	for _, blog := range blogs {
		blogId := blog.ID.Value
		if blogId == id {
			//skip which is equivalent to deleting here
		} else {
			tempBlogs = append(tempBlogs, blog)
		}

		for _, comment := range comments {
			if comment.BlogID != nil && comment.BlogID.Value == id {
				//skip which is equivalent to deleting here
			} else {
				tempComments = append(tempComments, comment)
			}
		}
	}

	if len(blogs) == len(tempBlogs) {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("no-op, a blog with id %s does not exist", idStr))
		return
	}

	blogs = tempBlogs
	comments = tempComments
	c.IndentedJSON(http.StatusOK, "")
}
