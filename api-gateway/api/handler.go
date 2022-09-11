package api

import (
	"context"
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
		"Message": "data is successfuly downloaded",
		"Success": true,
	})
}
