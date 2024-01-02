package author

import (
	"net/http"

	"github.com/literalog/library/internal/app/domain/cerror"
)

var (
	ErrEmptyId = cerror.New("empty id", http.StatusBadRequest)
)
