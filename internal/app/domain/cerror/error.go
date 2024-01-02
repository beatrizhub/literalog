package cerror

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status int    `json:"status"`
	Err    string `json:"message"`
}

func New(err string, status int) Error {
	return Error{
		Status: status,
		Err:    err,
	}
}

func (e Error) Error() string {
	return e.Err
}

func (e Error) Render(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	return json.NewEncoder(w).Encode(e)
}

func HandleError(e error, w http.ResponseWriter) {
	switch e := e.(type) {
	case Error:
		e.Render(w)
	default:
		Error{
			Status: http.StatusInternalServerError,
			Err:    "Internal Server Error",
		}.Render(w)
	}
}
