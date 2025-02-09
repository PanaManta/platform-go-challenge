package models

type Asset struct {
	Id             string      `json:"id"`
	Type           string      `json:"type"`
	Description    string      `json:"description"`
	StructuredData interface{} `json:"structured_data"`
}
