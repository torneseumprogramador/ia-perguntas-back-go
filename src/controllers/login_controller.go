package controllers

import (
	"fmt"
	"http_gin/src/DTOs"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/servicos"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginController struct{}

// @Summary Login
// @Description Autentica um administrador e retorna um token JWT junto com informações do administrador
// @Tags Login
// @Accept  json
// @Produce  json
// @Param   loginDTO body    DTOs.LoginDTO true  "Credenciais de Login"
// @Success 200 {object} model_views.AdmTokenView "Retorna o token JWT e informações do administrador"
// @Failure 400 {object} model_views.ErrorResponse "Email ou senha inválido ou erro ao gerar token"
// @Router /login [post]
func (hc *LoginController) Login(c *gin.Context) {
	var loginDTO DTOs.LoginDTO

	if err := c.BindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: err.Error(),
		})
		return
	}

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())

	credenciais := make(map[string]string)
	credenciais["email"] = loginDTO.Email

	adms, erro := servico.Repo.Where(credenciais)

	if erro != nil {
		c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
			Erro: erro.Error(),
		})
		return
	}

	if len(adms) > 0 && libs.CryptoEq(loginDTO.Senha, adms[0].Senha) {
		adm := adms[0]
		token, erro := tokenJwt(c, adm)
		if erro != nil {
			c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
				Erro: erro.Error(),
			})
		}

		c.JSON(http.StatusOK, model_views.AdmTokenView{
			Token: token,
			Id:    adm.Id,
			Nome:  adm.Nome,
			Email: adm.Email,
			Super: adm.Super,
		})

		return
	}

	c.JSON(http.StatusBadRequest, model_views.ErrorResponse{
		Erro: "Email ou senha inválido",
	})
}

// @Summary ValidarLogin
// @Description Verifica se token é valido
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 401 {object} model_views.ErrorResponse "Token inválido"
// @Router /login [post]
func (hc *LoginController) ValidarLogin(c *gin.Context) {
	c.JSON(http.StatusNoContent, model_views.ErrorResponse{})
}

func tokenJwt(c *gin.Context, adm models.Administrador) (string, error) {
	tempoExpiracao := time.Now().Add(time.Hour * 1)

	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = adm.Id
	claims["nome"] = adm.Nome
	claims["email"] = adm.Email
	claims["super"] = adm.Super
	claims["exp"] = tempoExpiracao.Unix()

	chave := libs.GetEnv("JWT_TOKEN", "desafio_go")
	tokenString, err := token.SignedString([]byte(chave))
	if err != nil {
		return "", fmt.Errorf("Login ou senha inválido")
	}

	return tokenString, nil
}
