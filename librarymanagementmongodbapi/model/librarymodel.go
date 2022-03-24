package librarymodel

type LibraryModel struct {
	BookName   string `json:"bookname"`
	BookAuthor string `json:"bookauthor"`
	Location   string `json:"location"`
	Uniqueid   string `json:"uniqueid"`
}
