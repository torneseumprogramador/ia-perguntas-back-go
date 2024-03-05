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

func petRepo() *repositorios.PetRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.PetRepositorioMySql{DB: db}
}

type PetsController struct{}

// @Summary Lista de pets
// @Description Retorna uma lista de todos os pets
// @Tags Pets
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Pet
// @Failure 400 {object} model_views.ErrorResponse
// @Router /pets [get]
// @Security Bearer
func (pc *PetsController) Index(c *gin.Context) {
	servico := servicos.NovoPetServico(petRepo())
	pets, erro := servico.ListaPetView()

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
	}

	c.JSON(http.StatusOK, pets)
}

// @Summary Cadastrar pet
// @Description Cadastra um novo pet
// @Tags Pets
// @Accept  json
// @Produce  json
// @Param   pet body    DTOs.PetDTO true  "Dados do Pet"
// @Success 201 {object} models.Pet
// @Failure 400 {object} model_views.ErrorResponse
// @Router /pets [post]
// @Security Bearer
func (pc *PetsController) Cadastrar(c *gin.Context) {
	var petDTO DTOs.PetDTO

	if err := c.BindJSON(&petDTO); err != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: err.Error(),
		})
		return
	}

	pet := models.Pet{
		Id:     "",
		Nome:   petDTO.Nome,
		DonoId: petDTO.DonoId,
		Tipo:   petDTO.Tipo,
	}

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	id, erro := servico.Repo.Adicionar(pet)
	pet.Id = id

	if erro == nil {
		c.JSON(http.StatusCreated, pet)
		return
	}

	c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
		Erro: erro.Error(),
	})
}

// @Summary Mostrar pet
// @Description Retorna os detalhes de um pet específico pelo ID
// @Tags Pets
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Pet"
// @Success 200 {object} models.Pet
// @Failure 400 {object} model_views.ErrorResponse "Pet não encontrado ou erro na busca"
// @Router /pets/{id} [get]
// @Security Bearer
func (pc *PetsController) Mostrar(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	petDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
		return
	}

	if petDb == nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: "pet não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, petDb)
}

// @Summary Excluir pet
// @Description Exclui um pet pelo ID
// @Tags Pets
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Pet"
// @Success 204
// @Failure 400 {object} model_views.ErrorResponse "Erro ao excluir o pet"
// @Router /pets/{id} [delete]
// @Security Bearer
func (pc *PetsController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	servico.Repo.Excluir(id)

	c.JSON(http.StatusNoContent, model_views.ErrorResponse{})
}

// @Summary Alterar pet
// @Description Altera os dados de um pet pelo ID
// @Tags Pets
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "ID do Pet"
// @Param   pet body    DTOs.PetDTO true  "Dados atualizados do Pet"
// @Success 200 {object} models.Pet
// @Failure 400 {object} model_views.ErrorResponse "Pet não encontrado ou erro na alteração"
// @Router /pets/{id} [put]
// @Security Bearer
func (pc *PetsController) Alterar(c *gin.Context) {
	id := c.Param("id")
	var petDTO DTOs.PetDTO

	if err := c.BindJSON(&petDTO); err != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: err.Error(),
		})
		return
	}

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	petDb, erro := servico.Repo.BuscarPorId(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
		return
	}

	if petDb == nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: "pet não encontrado",
		})
		return
	}

	petDb.Nome = petDTO.Nome
	petDb.DonoId = petDTO.DonoId
	petDb.Tipo = petDTO.Tipo

	erroAlterar := servico.Repo.Alterar(*petDb)

	if erroAlterar != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erroAlterar.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, petDb)
}
