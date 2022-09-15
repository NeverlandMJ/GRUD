package api

import (
	"context"
	"fmt"
	"github.com/NeverlnadMJ/GRUD/api-gateway/pkg/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"

	"github.com/NeverlnadMJ/GRUD/api-gateway/service"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
)

type Handler struct {
	provider service.Service
}

func NewHandler(prov service.Service) Handler {
	return Handler{
		provider: prov,
	}
}

func (h Handler) SavePosts(c *gin.Context) {
	numberOfPages, ok := c.Params.Get("limit")
	if !ok {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "number of pages is not provided",
			"Success": false,
		})
		return
	}

	n, err := strconv.Atoi(numberOfPages)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "parameter is not a number",
			"Success": false,
		})
		return
	}

	err = h.provider.Collecter.SavePosts(context.Background(), n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}

	c.JSON(http.StatusOK, render.M{
		"Message": "data is successfully downloaded",
		"Success": true,
	})
}

func (h Handler) GetPostsByUserID(c *gin.Context) {
	userID, ok := c.Params.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "number of pages is not provided",
			"Success": false,
		})
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "parameter is not a number",
			"Success": false,
		})
		return
	}

	posts, err := h.provider.Grud.GetUserPosts(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h Handler) GetPostByID(c *gin.Context) {
	postID, ok := c.Params.Get("postID")
	if !ok {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "number of pages is not provided",
			"Success": false,
		})
		return
	}

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "parameter is not a number",
			"Success": false,
		})
		return
	}

	post, err := h.provider.Grud.GetPostByID(context.Background(), id)
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, render.M{
					"Message": err.Error(),
					"Success": false,
				})
			default:
				c.JSON(http.StatusInternalServerError, render.M{
					"Message": err.Error(),
					"Success": false,
				})
			}
			return
		}
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h Handler) DeletePost(c *gin.Context) {
	postID, ok := c.Params.Get("postID")
	if !ok {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "number of pages is not provided",
			"Success": false,
		})
		return
	}

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "parameter is not a number",
			"Success": false,
		})
		return
	}
	err = h.provider.Grud.DeletePost(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}
	c.JSON(http.StatusOK, render.M{
		"Message": "post successfully deleted",
		"Success": true,
	})
}

func (h Handler) UpdateTitle(c *gin.Context) {
	var req request.UpdatePostRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}
	resp, err := h.provider.Grud.UpdateTitle(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h Handler) UpdateBody(c *gin.Context) {
	var req request.UpdatePostRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}
	resp, err := h.provider.Grud.UpdateBody(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h Handler) ListPosts(c *gin.Context) {
	pageS := c.DefaultQuery("page", "1")
	limitS := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageS)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "page query is not a number",
			"Success": false,
		})
		return
	}

	limit, err := strconv.Atoi(limitS)
	if err != nil {
		c.JSON(http.StatusBadRequest, render.M{
			"Message": "limit query is not a number",
			"Success": false,
		})
		return
	}

	resp, all, err := h.provider.Grud.ListPosts(context.Background(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, render.M{
			"Message": err.Error(),
			"Success": false,
		})
		return
	}

	c.JSON(http.StatusOK, render.M{
		"Message": fmt.Sprintf("%v/%v posts", page*limit, all),
		"Posts":   resp,
		"Success": true,
	})
}
