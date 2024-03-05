package controllers

import (
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TradutorController struct{}

func (pc *TradutorController) Index(c *gin.Context) {
	servico := servicos.IAServico{}
	palavras := servico.BuscarPalavras()
	c.JSON(http.StatusOK, palavras)
}
