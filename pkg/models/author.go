package models

import "github.com/google/uuid"

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateAuthorRequest struct {
	Name string `json:"name"`
}

func NewAuthor(req CreateAuthorRequest) *Author {
	return &Author{
		Id:   uuid.NewString(),
		Name: req.Name,
	}
}
