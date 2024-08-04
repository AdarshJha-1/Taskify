package model

type SignIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
