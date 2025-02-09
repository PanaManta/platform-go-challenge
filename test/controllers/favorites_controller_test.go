package tests

import (
	"net/http"
	"net/http/httptest"
	"platform-go-challenge/internal/controllers"
	"platform-go-challenge/internal/models"
	"platform-go-challenge/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFavoritesService struct {
	mock.Mock
}

var (
	GetUserFavoritesMockedData = []services.PopulatedUserFavorite{
		{
			Favorite: models.Favorite{
				UserId:  "user-123",
				AssetId: "asset123",
			},
			Asset: models.Asset{
				Id:          "asset123",
				Type:        "insight",
				Description: "Description",
				StructuredData: models.Insight{
					Text: "mocked insight",
				},
			},
		},
	}

	AddUserFavoriteMockedData = models.Favorite{
		UserId:  "user-123",
		AssetId: "asset123",
	}

	DeleteUserFavoriteMockedData = models.Favorite{
		UserId:  "user-123",
		AssetId: "asset123",
	}
)

func (m *MockFavoritesService) GetUserFavorites(userId string) []services.PopulatedUserFavorite {
	m.Called(userId)
	return GetUserFavoritesMockedData
}

func (m *MockFavoritesService) AddUserFavorite(userId string, assetId string) models.Favorite {
	m.Called(userId, assetId)
	return AddUserFavoriteMockedData
}

func (m *MockFavoritesService) DeleteUserFavorite(userId string, assetId string) models.Favorite {
	m.Called(userId, assetId)
	return DeleteUserFavoriteMockedData
}

func TestGetUserFavorites(t *testing.T) {
	mockService := new(MockFavoritesService)
	controller := controllers.NewFavoritesController(mockService)

	mockService.On("GetUserFavorites", "user-123").Return(GetUserFavoritesMockedData)

	responceRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responceRecorder)
	ctx.Set("user_id", "user-123")

	controller.GetUserFavorites(ctx)

	mockService.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responceRecorder.Code)

	snaps.MatchSnapshot(t, string(responceRecorder.Body.Bytes()))
}

func TestAddUserFavorite(t *testing.T) {
	mockService := new(MockFavoritesService)
	controller := controllers.NewFavoritesController(mockService)

	mockService.On("AddUserFavorite", "user-123", "asset123").Return(AddUserFavoriteMockedData)

	responceRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responceRecorder)
	ctx.Set("user_id", "user-123")
	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "asset_id",
		Value: "asset123",
	})

	controller.AddUserFavorite(ctx)

	mockService.AssertExpectations(t)

	assert.Equal(t, http.StatusCreated, responceRecorder.Code)

	snaps.MatchSnapshot(t, string(responceRecorder.Body.Bytes()))
}

func TestDeleteUserFavorite(t *testing.T) {
	mockService := new(MockFavoritesService)
	controller := controllers.NewFavoritesController(mockService)

	mockService.On("DeleteUserFavorite", "user-123", "asset123").Return(DeleteUserFavoriteMockedData)

	responceRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responceRecorder)
	ctx.Set("user_id", "user-123")
	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "asset_id",
		Value: "asset123",
	})

	controller.DeleteUserFavorite(ctx)

	mockService.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, responceRecorder.Code)

	snaps.MatchSnapshot(t, string(responceRecorder.Body.Bytes()))
}
