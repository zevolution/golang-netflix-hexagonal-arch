package entities

type Author struct {
	UserName      string  		`json:"userName"`
	Name          string        `json:"name"`
	AvatarUrl     string        `json:"avatarUrl"`
}