package entities

type Software struct {
	Name  string `json:"name"`
	Description string `json:"description"`
	Score int `json:"score"`
	Author *Author `json:"author"`
	Host Hosts `json:"host"`
}