package main

import (
	"http_gin/src/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startWebApp() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true // permitir qualquer origem
		},
		MaxAge: 12 * time.Hour,
	}))

	config.Routes(router)

	router.Run(":5000") // Por padrão, escuta na porta 5000
}

// @title Desafio de Golang
// @description Este é um app construido junto com os alunos do torne-se um programador, com objetivo em aprender a linguagem Golang e seus componentes
// @version 1.0
// @BasePath /
// @contact.name Danilo Aparecido
// @contact.url https://www.torneseumprogramador.com.br/cursos/desafio_go_lang
// @contact.email suporte@torneseumprogramador.com.br
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	startWebApp()
}
