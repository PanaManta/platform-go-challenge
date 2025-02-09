package services

import (
	"math/rand"
	"platform-go-challenge/internal/models"
	"platform-go-challenge/internal/repositories"
)

type FavoritesService interface {
	GetUserFavorites(userId string) []PopulatedUserFavorite
	AddUserFavorite(userId string, assetId string) models.Favorite
	DeleteUserFavorite(userId string, assetId string) models.Favorite
}

type DefaultFavoritesService struct {
	favoritesRepository repositories.FavoritesRepository
	assetRepository     repositories.AssetRepository
}

type PopulatedUserFavorite struct {
	Favorite models.Favorite
	Asset    models.Asset
}

func (s *DefaultFavoritesService) GetUserFavorites(userId string) []PopulatedUserFavorite {
	// TODO: pagination
	userFavorites := s.favoritesRepository.GetUserFavorites(userId, "0", 10)

	availableAssetTypes := []string{"chart", "audience", "insight"}
	result := make([]PopulatedUserFavorite, len(userFavorites))
	for i, favorite := range userFavorites {
		randomAsset := s.assetRepository.GetAsset(favorite.AssetId, availableAssetTypes[rand.Intn(2)])
		result[i] = PopulatedUserFavorite{Favorite: favorite, Asset: randomAsset}
	}

	return result
}

func (s *DefaultFavoritesService) AddUserFavorite(userId string, assetId string) models.Favorite {
	added := s.favoritesRepository.AddUserFavorite(userId, assetId)

	return added
}

func (s *DefaultFavoritesService) DeleteUserFavorite(userId string, assetId string) models.Favorite {
	deleted := s.favoritesRepository.DeleteUserFavorite(userId, assetId)

	return deleted
}

func NewDefaultFavoritesService(favoritesRepo repositories.FavoritesRepository, assetRepo repositories.AssetRepository) FavoritesService {
	return &DefaultFavoritesService{
		favoritesRepository: favoritesRepo,
		assetRepository:     assetRepo,
	}
}
