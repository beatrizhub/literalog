package book

type Book struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Genre   []string `json:"genre"`
}

type ReadBook struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	BookID   int    `json:"book_id"`
	ReadDate string `json:"read_date"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}

type ToBeReadBook struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
