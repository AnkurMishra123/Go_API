package models

type Car struct {
	ID     int    `json:"id"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Color  string `json:"color"`
	Status string `json:"status"`
	// Other relevant fields
}
