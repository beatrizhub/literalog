package models

import (
	"strings"

	"github.com/google/uuid"
)

type Book struct {
	Id        string   `json:"id" bson:"_id"`
	Title     string   `json:"title" bson:"title" validate:"required"`
	AuthorId  string   `json:"author_id" bson:"author_id"`
	Isbn      string   `json:"isbn" bson:"isbn"`
	SeriesId  string   `json:"series_id" bson:"series_id"`
	SeriesNo  int      `json:"series_no" bson:"series_no"`
	Year      int      `json:"year" bson:"year"`
	Publisher string   `json:"publisher" bson:"publisher"`
	Language  string   `json:"language" bson:"language"`
	Format    Format   `json:"format" bson:"format"`
	PagesNo   int      `json:"pages_no" bson:"pages_no"`
	HoursNo   int      `json:"hours_no" bson:"hours_no"`
	Genre     []string `json:"genre" bson:"genre"`
	Blurb     string   `json:"blurb" bson:"blurb"`
	Cover     string   `json:"cover" bson:"cover"`
	IsBook    bool     `json:"is_book" bson:"is_book"`
}

type Format string

const (
	Hardcover Format = "Hardcover"
	Paperback Format = "Paperback"
	Digital   Format = "Digital"
	Audio     Format = "Audio"
)

func NewFormat(s string) Format {
	switch strings.ToLower(s) {
	case "hardcover":
		return Hardcover
	case "paperback":
		return Paperback
	case "digital":
		return Digital
	case "audio":
		return Audio
	default:
		return ""
	}
}

func NewBook(b Book) *Book {
	return &Book{
		Id:        uuid.NewString(),
		Title:     b.Title,
		AuthorId:  b.AuthorId,
		Isbn:      b.Isbn,
		SeriesId:  b.SeriesId,
		SeriesNo:  b.SeriesNo,
		Year:      b.Year,
		Publisher: b.Publisher,
		Language:  b.Language,
		Format:    b.Format,
		PagesNo:   b.PagesNo,
		HoursNo:   b.HoursNo,
		Genre:     b.Genre,
		Blurb:     b.Blurb,
		Cover:     b.Cover,
		IsBook:    b.IsBook,
	}
}
