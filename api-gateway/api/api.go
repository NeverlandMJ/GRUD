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

	return router

}