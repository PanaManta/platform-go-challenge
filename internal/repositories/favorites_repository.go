package repositories

import (
	"platform-go-challenge/internal/models"
	"sync"
	"time"
)

type FavoritesRepository interface {
	GetUserFavorites(userId string, searchAfter string, limit int) []models.Favorite
	AddUserFavorite(userId string, assetId string) models.Favorite
	DeleteUserFavorite(userId string, assetId string) models.Favorite
}

type DummyFavoritesRepository struct {
	data map[string][]models.Favorite
	lock sync.RWMutex
}

var once sync.Once
var instance *DummyFavoritesRepository

func NewDummyFavoritesRepository() FavoritesRepository {
	once.Do(func() {
		instance = &DummyFavoritesRepository{
			data: make(map[string][]models.Favorite, 100),
		}
	})
	return instance
}

func (r *DummyFavoritesRepository) GetUserFavorites(userId string, searchAfter string, limit int) []models.Favorite {
	r.lock.RLock()
	defer r.lock.RUnlock()

	// Todo: add pagination
	return r.data[userId]
}

func (r *DummyFavoritesRepository) AddUserFavorite(userId string, assetId string) models.Favorite {
	r.lock.Lock()
	defer r.lock.Unlock()

	favoriteToAdd := models.Favorite{
		UserId:    userId,
		AssetId:   assetId,
		CreatedAt: time.Now(),
	}

	r.data[userId] = append(r.data[userId], favoriteToAdd)
	return favoriteToAdd
}

func (r *DummyFavoritesRepository) DeleteUserFavorite(userId string, assetId string) models.Favorite {
	r.lock.Lock()
	defer r.lock.Unlock()

	userFavorites := r.data[userId]
	for i, favorite := range userFavorites {
		if favorite.AssetId != assetId {
			continue
		}

		r.data[userId] = append(userFavorites[:i], userFavorites[i+1:]...)
		return favorite
	}

	return models.Favorite{}
}
