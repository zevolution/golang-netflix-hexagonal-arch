package models

type SoftwaresByAuthorModel struct {
	Author *AuthorModel `json:"author"`
	Softwares *[]SoftwareModel `json:"softwares"`
}