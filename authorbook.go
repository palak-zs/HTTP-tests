package AuthorBooksHTTP

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	id            int    `json:"message"`
	title         string `json:"message"`
	author        *Author `json:"message"`
	publication   string `json:"message"`
	publishedDate string `json:"message"`
}

type Author struct {
	firstName string `json:"message"`
	lastName  string `json:"message"`
	DOB       string `json:"message"`
	penName   string `json:"message"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8000/books")
	if err != nil {
		_ = fmt.Errorf("Error. error: %s", err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		_ = fmt.Errorf("Failed to fetch the response body. Got error : %v", err)
	}

	var b Book

	err = json.Unmarshal(body, &b)
	if err != nil {
		_ = fmt.Errorf("Failed to unmarshal the response body. Got error : %v", err)
	}

	_, _ = fmt.Print(b)

}

func GetByID(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8000/books/id")
	if err != nil {
		_ = fmt.Errorf("Error. error: %s", err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		_ = fmt.Errorf("Failed to fetch the response body. Got error : %v", err)
	}

	var b Book

	err = json.Unmarshal(body, &b)
	if err != nil {
		_ = fmt.Errorf("Failed to unmarshal the response body. Got error : %v", err)
	}

	_, _ = fmt.Print(b)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	if strconv.strings.Split(b.publishedDate,"-")
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {

		FetchData(b.id)

	})

}

func PostAuthor(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/author", func(w http.ResponseWriter, r *http.Request) {
		var a Author

	})

}
