package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/download/pages/:limit", h.SavePosts)
	router.GET("/posts/:userID", h.GetPostsByUserID)
	router.GET("/post/:postID", h.GetPostByID)
	router.PUT("/update/title", h.UpdateTitle)
	router.PUT("/update/body", h.UpdateBody)
	router.DELETE("/delete/:postID", h.DeletePost)
	router.GET("/get/posts/", h.ListPosts)

	return router
}
