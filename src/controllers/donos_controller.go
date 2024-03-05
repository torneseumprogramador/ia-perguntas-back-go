package controllers

import (
	"http_gin/src/DTOs"
	"http_gin/src/database"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func donoRepo() *repositorios.DonoRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.DonoRepositorioMySql{DB: db}
}

type DonosController struct{}

// @Summary Lista de donos
// @Description Retorna uma lista de todos os donos
// @Tags Donos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Dono
// @Failure 400 {object} model_views.ErrorResponse
// @Router /donos [get]
// @Security Bearer
func (pc *DonosController) Index(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donos, erro := servico.Repo.Lista()

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
	}

	c.JSON(http.StatusOK, donos)
}

// @Summary Mostrar dono
// @Description Retorna os detalhes de um dono específico pelo ID
// @Tags Donos
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Dono"
// @Success 200 {object} models.Dono
// @Failure 400 {object} model_views.ErrorResponse
// @Router /donos/{id} [get]
// @Security Bearer
func (pc *DonosController) Mostrar(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donoDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
		return
	}

	if donoDb == nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: "pet não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, donoDb)
}

// @Summary Cadastrar dono
// @Description Cadastra um novo dono
// @Tags Donos
// @Accept  json
// @Produce  json
// @Param   dono body    DTOs.DonoDTO true  "Dados do Dono"
// @Success 201 {object} models.Dono
// @Failure 400 {object} model_views.ErrorResponse
// @Router /donos [post]
// @Security Bearer
func (pc *DonosController) Cadastrar(c *gin.Context) {
	var donoDTO DTOs.DonoDTO

	if err := c.BindJSON(&donoDTO); err != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: err.Error(),
		})
		return
	}

	dono := models.Dono{
		Id:       "",
		Nome:     donoDTO.Nome,
		Telefone: donoDTO.Telefone,
	}

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	id, erro := servico.Repo.Adicionar(dono)
	dono.Id = id

	if erro == nil {
		c.JSON(http.StatusCreated, dono)
		return
	}

	c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
		Erro: erro.Error(),
	})
}

// @Summary Excluir dono
// @Description Exclui um dono pelo ID
// @Tags Donos
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Dono"
// @Success 204
// @Router /donos/{id} [delete]
// @Security Bearer
func (pc *DonosController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	servico.Repo.Excluir(id)

	c.JSON(http.StatusNoContent, model_views.ErrorResponse{})
}

// @Summary Alterar dono
// @Description Altera os dados de um dono pelo ID
// @Tags Donos
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Dono"
// @Param   dono body    DTOs.DonoDTO true  "Dados do Dono"
// @Success 200 {object} models.Dono
// @Failure 400 {object} model_views.ErrorResponse
// @Router /donos/{id} [put]
// @Security Bearer
func (pc *DonosController) Alterar(c *gin.Context) {
	id := c.Param("id")
	var donoDTO DTOs.DonoDTO

	if err := c.BindJSON(&donoDTO); err != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: err.Error(),
		})
		return
	}

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donoDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
		return
	}

	if donoDb == nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: "pet não encontrado",
		})
		return
	}

	donoDb.Nome = donoDTO.Nome
	donoDb.Telefone = donoDTO.Telefone

	erroAlterar := servico.Repo.Alterar(*donoDb)

	if erroAlterar != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erroAlterar.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, donoDb)
}
