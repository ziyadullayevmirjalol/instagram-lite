package models


type Comments struct {
	Id        int
	UserId    int
	PostId    int
	Text      string
	CreatedAt string
}