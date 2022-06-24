package AuthorBooksHTTP

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	testcases := []struct {
		desc   string
		input  string
		output []Book
	}{
		{"List of books with their details.", "/books", []Book{
			{1, "XYZ", &Author{"Palak", "Kejriwal", "07-10-1998", "cilios"}, "Pqrs", "06-11-1940"},
			{2, "XYZ1", &Author{"Alice", "Thomas", "12-11-1997", "ninja"}, "Pqrs1", "08-01-1950"}}},
	}

	for _, test := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", test.input, nil)

		GetAll(w, req)
		resp := w.Result()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		var allBooks []Book

		err = json.Unmarshal(data, &allBooks)
		if err != nil {
			return
		}

		assert.Equal(t, test.output, allBooks)

		err = resp.Body.Close()

	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		desc   string
		input  string
		output Book
	}{
		{"The details for book XYZ: ", "/books/xyz",
			Book{1, "XYZ", &Author{"Alice", "Thomas", "12-11-1997", "ninja"}, "Pqrs", "06-11-1940"}},
	}

	for _, test := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", test.input, nil)

		GetByID(w, req)
		resp := w.Result()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		var Book Book

		err = json.Unmarshal(data, &Book)
		if err != nil {
			return
		}

		assert.Equal(t, test.output, Book)

		err = resp.Body.Close()
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc       string
		book       Book
		statusCode int
	}{
		{"Success Case.", Book{1, "ABC", Author{"Palak", "Kejriwal", "07-10-1998", "cilios"}, "Arihanth", "18-08-2018"},
			http.StatusCreated},
		{"Blank book name.", Book{2, "", Author{"Palak", "Kejriwal", "07-10-1998", "cilios"}, "Oxford", "21-04-1985"},
			http.StatusBadRequest},
		{"Invalid author.", Book{3, "ABC", Author{"John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-1985"},
			http.StatusBadRequest,
		},
		{"Invalid publication date.", Book{4, "ABC", Author{"John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-1879"},
			http.StatusBadRequest,
		},
		{"Invalid publication date.", Book{5, "ABC", Author{"John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-2031"},
			http.StatusBadRequest,
		},
	}

	for _, test := range testcases {
		newData, _ := json.Marshal(test.book)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(newData))
		PostBook(w, req)
		resp := w.Result()

		assert.Equal(t, test.statusCode, resp.StatusCode)

	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc       string
		author     Author
		statusCode int
	}{
		{"Valid case.", Author{"Palak", "Kejriwal", "07-10-1998", "cilios"},
			http.StatusCreated},
		{"Blank first name.", Author{"", "Kejriwal", "07-10-1998", "cilios"},
			http.StatusBadRequest},
		//{"Blank last name.", Author{"Palak", "", "07-10-1998", "cilios"},
		//	http.StatusBadRequest},
		{"Blank name.", Author{"", "", "07-10-1998", "cilios"},
			http.StatusBadRequest},
		{"Blank penname.", Author{"Palak", "Kejriwal", "07-10-1998", ""},
			http.StatusCreated},
	}

	for _, test := range testcases {
		newData, _ := json.Marshal(test.author)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/author", bytes.NewBuffer(newData))
		PostBook(w, req)
		resp := w.Result()

		assert.Equal(t, test.statusCode, resp.StatusCode)

	}
}
