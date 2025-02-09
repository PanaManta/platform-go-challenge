package app

import (
	"log"
	"platform-go-challenge/config"
	"platform-go-challenge/internal/controllers"
	"platform-go-challenge/internal/middlewares"
	"platform-go-challenge/internal/repositories"
	"platform-go-challenge/internal/routes"
	"platform-go-challenge/internal/services"

	_ "platform-go-challenge/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Router              *gin.Engine
	FavoritesController *controllers.FavoritesController
}

func (app *App) Init() {
	config.LoadConfig()

	app.Router = gin.Default()
	app.registerSwagger()
	app.setupDependencies()
	app.registerApiRoutes()
}

func (app *App) setupDependencies() {
	favoritesRepository := repositories.NewDummyFavoritesRepository()
	assetRepository := repositories.NewDummyAssetRepository()
	favoritesService := services.NewDefaultFavoritesService(favoritesRepository, assetRepository)
	app.FavoritesController = controllers.NewFavoritesController(favoritesService)
}

func (app *App) registerApiRoutes() {
	api := app.Router.Group("/api")
	api.Use(middlewares.AuthUserMiddleware)
	routes.RegisterFavoritesRoutes(api, app.FavoritesController)
}

func (app *App) registerSwagger() {
	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (app *App) Start() {
	log.Printf("Starting server on port %s...", config.Config.Port)
	if err := app.Router.Run(":" + config.Config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
