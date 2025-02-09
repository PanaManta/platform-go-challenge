package controllers

import (
	"net/http"
	"platform-go-challenge/internal/services"
	"platform-go-challenge/internal/viewmodels"

	"github.com/gin-gonic/gin"
)

type FavoritesController struct {
	UserFavoritesService services.FavoritesService
}

func NewFavoritesController(service services.FavoritesService) *FavoritesController {
	return &FavoritesController{
		UserFavoritesService: service,
	}
}

// GetUserFavorites godoc
// @Summary Get User Favorites
// @Description Retrieves all the favorites for the user. Requires a valid user JWT in the Authorization header.
// @Tags favorites
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with Bearer token"
// @Success 200 {array} viewmodels.FavoriteView "List of user favorites"
// @Router /api/favorites [get]
func (c *FavoritesController) GetUserFavorites(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	userPopulatedFavorites := c.UserFavoritesService.GetUserFavorites(userId.(string))

	response := make([]viewmodels.FavoriteView, len(userPopulatedFavorites))
	for i, populatedFavorite := range userPopulatedFavorites {
		response[i] = viewmodels.ToFavoriteView(populatedFavorite.Favorite, populatedFavorite.Asset)
	}

	ctx.JSON(http.StatusOK, response)
}

// AddUserFavorite godoc
// @Summary Add User Favorite in the temporary repository
// @Description Adds a user favorite in the repository
// @Tags favorites
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header with Bearer token"
// @Param asset_id path string true "Asset ID to be added as a favorite"
// @Success 201 {object} viewmodels.FavoriteViewActionResponse "Added entry"
// @Router /api/favorites/{asset_id} [post]
func (c *FavoritesController) AddUserFavorite(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	assetId := ctx.Param("asset_id")

	added := c.UserFavoritesService.AddUserFavorite(userId.(string), assetId)

	ctx.JSON(http.StatusCreated, viewmodels.AddedFavoriteSuccess(added))
}

// DeleteUserFavorite godoc
// @Summary Delete User Favorite in the temporary repository
// @Description Deletes a user favorite in the repository
// @Tags favorites
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header with Bearer token"
// @Param asset_id path string true "Asset ID to be added as a favorite"
// @Success 200 {object} viewmodels.FavoriteViewActionResponse "Deleted entry"
// @Router /api/favorites/{asset_id} [post]
func (c *FavoritesController) DeleteUserFavorite(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	assetId := ctx.Param("asset_id")

	deleted := c.UserFavoritesService.DeleteUserFavorite(userId.(string), assetId)

	ctx.JSON(http.StatusOK, viewmodels.DeletedFavoriteSuccess(deleted))
}
