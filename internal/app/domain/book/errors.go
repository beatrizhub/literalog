package book

import (
	"net/http"

	"github.com/literalog/library/internal/app/domain/cerror"
)

var (
	ErrEmptyId            = cerror.New("empty id", http.StatusBadRequest)
	ErrInvalidTitle       = cerror.New("invalid title", http.StatusBadRequest)
	ErrEmptyTitle         = cerror.New("empty title", http.StatusBadRequest)
	ErrInvalidTitleLength = cerror.New("title must be between x and y", http.StatusBadRequest)
)
