package viewmodels

import (
	"platform-go-challenge/internal/models"
	"platform-go-challenge/internal/services"
)

type FavoriteView struct {
	AssetId     string `json:"asset_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Data        interface{}
}

type FavoriteViewActionResponse struct {
	AssetId string `json:"asset_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func MapToFavoriteViews(populatedFavorites []services.PopulatedUserFavorite) []FavoriteView {
	result := make([]FavoriteView, len(populatedFavorites))
	for i, populatedFavorite := range populatedFavorites {
		result[i] = ToFavoriteView(populatedFavorite.Favorite, populatedFavorite.Asset)
	}

	return result
}

func ToFavoriteView(userFavorite models.Favorite, asset models.Asset) FavoriteView {
	return FavoriteView{
		AssetId:     userFavorite.AssetId,
		Type:        asset.Type,
		Description: asset.Description,
		Data:        asset.StructuredData,
	}
}

func AddedFavoriteSuccess(favorite models.Favorite) FavoriteViewActionResponse {
	return FavoriteViewActionResponse{
		AssetId: favorite.AssetId,
		Status:  "success",
		Message: "Favorite added successfully",
	}
}

func DeletedFavoriteSuccess(favorite models.Favorite) FavoriteViewActionResponse {
	return FavoriteViewActionResponse{
		AssetId: favorite.AssetId,
		Status:  "success",
		Message: "Favorite deleted successfully",
	}
}
