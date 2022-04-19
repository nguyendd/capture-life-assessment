package routers

import (
	v1 "capturelife.assessment.daniel/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/blogs", v1.GetBlogs)
	r.GET("/blogs/:id", v1.GetBlogByID)
	r.POST("/blogs", v1.PostBlogs)
	r.PATCH("/blogs/:id", v1.UpdateBlog)
	r.DELETE("/blogs/:id", v1.DeleteBlog)
	r.GET("/comments", v1.GetComments)
	r.GET("/comments/:id", v1.GetCommentsByBlogID)
	r.POST("/comments", v1.PostComments)
	r.PATCH("/comments/:id", v1.UpdateComment)
	r.DELETE("/comments/:id", v1.DeleteComment)
	return r
}
