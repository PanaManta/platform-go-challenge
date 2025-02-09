package models

type Chart struct {
	Title string      `json:"title"`
	XAxis string      `json:"x_axis"`
	YAxis string      `json:"y_axis"`
	Data  []DataPoint `json:"data"`
}

type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
