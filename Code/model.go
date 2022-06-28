package Code

type Book struct {
	ID            string  `json:"ID"`
	AuthorID      int     `json:"AuthorID"`
	Title         string  `json:"Title"`
	Author        *Author `json:"Author"`
	Publication   string  `json:"Publication"`
	PublishedDate string  `json:"PublishedDate"`
}

type Author struct {
	AuthorID  int    `json:"AuthorID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	DOB       string `json:"DOB"`
	PenName   string `json:"PenName"`
}
