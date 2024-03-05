package controllers

import (
	"http_gin/src/model_views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

// @Summary Home
// @Description Json de apresentação da API
// @Tags home
// @Accept  json
// @Produce  json
// @Success 200 {object} model_views.Home
// @Router / [get]
func (hc *HomeController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, model_views.Home{
		Mensagem: "Bem-vindo à API com Golang e Gin",
		Docs:     "/swagger/index.html",
	})
}
