package models

type AuthorModel struct {
	UserName      string  		`json:"username"`
	Name          string        `json:"name"`
	AvatarUrl     string        `json:"avatarUrl"`
}