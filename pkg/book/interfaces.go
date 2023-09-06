package book

type Handler interface {
	CreateBook(book Book) (Book, error)
	CreateReadBook(readBook ReadBook) (ReadBook, error)
	CreateToBeReadBook(toBeReadBook ToBeReadBook) (ToBeReadBook, error)

	UpdateBook(id int, book Book) error
	UpdateReadBook(id int, readBook ReadBook) error

	DeleteBook(id int) error
	DeleteReadBook(id int) error
	DeleteToBeReadBook(id int) error

	GetBooks() ([]Book, error)
	GetReadBooks() ([]ReadBook, error)
	GetToBeReadBooks() ([]ToBeReadBook, error)

	GetBookByID(id int) (Book, error)
	GetReadBookByID(id int) (ReadBook, error)

	GetBooksByGenre(genre string) ([]Book, error)
	GetBooksByAuthor(author string) ([]Book, error)
	GetBooksByTitle(title string) ([]Book, error)

	GetReadBooksByUserID(userID int) ([]ReadBook, error)
	GetToBeReadBooksByUserID(userID int) ([]ToBeReadBook, error)

	GetBooksRecommendations(userID int) ([]Book, error)
}
