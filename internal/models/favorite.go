package models

import "time"

type Favorite struct {
	UserId    string    `json:"user_id"`
	AssetId   string    `json:"asset_id"`
	CreatedAt time.Time `json:"created_at"`
}
