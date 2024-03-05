package middlewares

import (
	"fmt"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model_views.ErrorResponse{Erro: "Cabeçalho de autorização não fornecido"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model_views.ErrorResponse{Erro: "Formato de autorização inválido"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		chave := libs.GetEnv("JWT_TOKEN", "desafio_go")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}

			return []byte(chave), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, model_views.ErrorResponse{Erro: "Token inválido"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, idOk := claims["id"].(string)
			nome, nomeOk := claims["nome"].(string)
			email, emailOk := claims["email"].(string)
			super, superOk := claims["super"].(bool)

			if !nomeOk || !emailOk || !idOk || !superOk {
				c.JSON(http.StatusUnauthorized, model_views.ErrorResponse{Erro: "Falha ao processar claims do token"})
				c.Abort()
				return
			}

			currentRoute := c.FullPath()
			if !super && strings.Contains(currentRoute, "administradores") {
				c.JSON(http.StatusForbidden, model_views.ErrorResponse{Erro: "Usuário sem acesso a esta área"})
				c.Abort()
				return
			}

			adm := model_views.AdmView{
				Id:    id,
				Nome:  nome,
				Email: email,
				Super: super,
			}

			c.Set("adm", adm)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, model_views.ErrorResponse{Erro: "Token JWT inválido ou expirado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
