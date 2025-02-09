package routes

import (
	"platform-go-challenge/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterFavoritesRoutes(router *gin.RouterGroup, controller *controllers.FavoritesController) {
	router.GET("/favorites", controller.GetUserFavorites)
	router.POST("/favorites/:asset_id", controller.AddUserFavorite)
	router.DELETE("/favorites/:asset_id", controller.DeleteUserFavorite)
}
