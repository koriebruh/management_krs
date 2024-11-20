package dto

type LoginRes struct {
	ID       int    `json:"id"`
	NIM      string `json:"nim"`
	UserName string `json:"username"`
	Token    string `json:"token"`
}
