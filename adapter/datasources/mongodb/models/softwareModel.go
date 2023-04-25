package models

type SoftwareModel struct {
	Name  string `json:"name"`
	Description string `json:"description"`
	Score int `json:"score"`
	Host int `json:"host"`
}