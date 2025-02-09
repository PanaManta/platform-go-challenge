package models

type Audience struct {
	Gender        string `json:"gender"`
	BirthCountry  string `json:"birth_country"`
	AgeGroup      string `json:"age_group"`
	HoursOnSocial int    `json:"hours_on_social"`
	NumPurchases  int    `json:"num_purchases"`
}
