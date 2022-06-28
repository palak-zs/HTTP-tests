package Code

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
			{"1", 1, "XYZ", &Author{1, "Palak", "Kejriwal", "07-10-1998", "cilios"}, "Arihanth", "21-04-2001"},
			{"2", 3, "Dusk", &Author{3, "John", "Crook", "12-03-1978", "ninja"}, "Penguin", "31-05-1997"}}},
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
			Book{"1", 1, "XYZ", &Author{1, "Alice", "Thomas", "12-11-1997", "ninja"}, "Pqrs", "06-11-1940"}},
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
		{"Success Case.", Book{"1", 1, "ABC", &Author{1, "Palak", "Kejriwal", "07-10-1998", "cilios"}, "Arihanth", "18-08-2018"},
			http.StatusCreated},
		{"Blank book name.", Book{"2", 2, "", &Author{2, "Palak", "Kejriwal", "07-10-1998", "cilios"}, "Oxford", "21-04-1985"},
			http.StatusBadRequest},
		{"Invalid Author.", Book{"3", 3, "ABC", &Author{3, "John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-1985"},
			http.StatusBadRequest,
		},
		{"Invalid Publication date.", Book{"4", 4, "ABC", &Author{4, "John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-1879"},
			http.StatusBadRequest,
		},
		{"Invalid Publication date.", Book{"5", 5, "ABC", &Author{5, "John", "Fernandes", "19-12-1972", "cilios"}, "Oxford", "21-04-2031"},
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
		{"Valid case.", Author{1, "Palak", "Kejriwal", "07-10-1998", "cilios"},
			http.StatusCreated},
		{"Blank first name.", Author{2, "", "Kejriwal", "07-10-1998", "cilios"},
			http.StatusBadRequest},
		//{"Blank last name.", Author{"Palak", "", "07-10-1998", "cilios"},
		//	http.StatusBadRequest},
		{"Blank name.", Author{3, "", "", "07-10-1998", "cilios"},
			http.StatusBadRequest},
		{"Blank penname.", Author{4, "Palak", "Kejriwal", "07-10-1998", ""},
			http.StatusCreated},
	}

	for _, test := range testcases {
		newData, _ := json.Marshal(test.author)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/Author", bytes.NewBuffer(newData))
		PostBook(w, req)
		resp := w.Result()

		assert.Equal(t, test.statusCode, resp.StatusCode)

	}
}
