package repositories

import (
	"math/rand"
	"platform-go-challenge/internal/models"
)

type AssetRepository interface {
	GetAsset(assetId string, assetType string) models.Asset
}

type DummyAssetRepository struct{}

func NewDummyAssetRepository() AssetRepository {
	return &DummyAssetRepository{}
}

func (r *DummyAssetRepository) GetAsset(assetId string, assetType string) models.Asset {
	switch assetType {
	case "chart":
		return chart(assetId)
	case "audience":
		return audience(assetId)
	case "insight":
		return insight((assetId))
	default:
		return models.Asset{}
	}
}

func chart(assetId string) models.Asset {
	dataPoints := make([]models.DataPoint, 5)
	for i := 0; i < 5; i++ {
		dataPoints[i] = models.DataPoint{
			X: float64(i),
			Y: rand.Float64() * 100,
		}
	}

	chartData := models.Chart{
		Title: "Sample Sales Chart",
		XAxis: "Months",
		YAxis: "Sales",
		Data:  dataPoints,
	}

	return models.Asset{
		Id:             assetId,
		Type:           "chart",
		Description:    "Sample description for chart",
		StructuredData: chartData,
	}
}

func audience(assetId string) models.Asset {
	return models.Asset{
		Id:          assetId,
		Type:        "audience",
		Description: "Audience Characteristics",
		StructuredData: models.Audience{
			Gender:        "Male",
			AgeGroup:      "18-24",
			HoursOnSocial: rand.Intn(10) * 5,
			NumPurchases:  rand.Intn(5),
		},
	}
}

func insight(assetId string) models.Asset {
	insights := []string{
		"40% of users spend 3+ hours on social media.",
		"Sales increased by 25% in Q4.",
		"User engagement spiked during holidays.",
	}

	return models.Asset{
		Id:          assetId,
		Type:        "insight",
		Description: "User Insight",
		StructuredData: models.Insight{
			Text: insights[rand.Intn(2)],
		},
	}
}
