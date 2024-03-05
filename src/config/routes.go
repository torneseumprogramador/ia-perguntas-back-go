package config

import (
	"http_gin/src/controllers"
	"http_gin/src/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.Static("/docs/", "./docs")
	url := ginSwagger.URL("/docs/swagger.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	homeController := controllers.HomeController{}
	router.GET("/", homeController.Index)

	loginController := controllers.LoginController{}
	router.POST("/login", loginController.Login)

	tradutorController := controllers.TradutorController{}
	router.GET("/tradutor", tradutorController.Index)

	protectedRoutes := router.Group("/").Use(middlewares.AuthRequired())
	{
		protectedRoutes.HEAD("/login/validar", loginController.ValidarLogin)

		petsController := controllers.PetsController{}
		protectedRoutes.GET("/pets", petsController.Index)
		protectedRoutes.POST("/pets", petsController.Cadastrar)
		protectedRoutes.GET("/pets/:id", petsController.Mostrar)
		protectedRoutes.DELETE("/pets/:id", petsController.Excluir)
		protectedRoutes.PUT("/pets/:id", petsController.Alterar)

		donosController := controllers.DonosController{}
		protectedRoutes.GET("/donos", donosController.Index)
		protectedRoutes.POST("/donos", donosController.Cadastrar)
		protectedRoutes.DELETE("/donos/:id", donosController.Excluir)
		protectedRoutes.PUT("/donos/:id", donosController.Alterar)
		protectedRoutes.GET("/donos/:id", donosController.Mostrar)

		administradoresController := controllers.AdministradoresController{}
		protectedRoutes.GET("/administradores", administradoresController.Index)
		protectedRoutes.POST("/administradores", administradoresController.Cadastrar)
		protectedRoutes.DELETE("/administradores/:id", administradoresController.Excluir)
		protectedRoutes.PUT("/administradores/:id", administradoresController.Alterar)
		protectedRoutes.GET("/administradores/:id", administradoresController.Mostrar)
	}
}
