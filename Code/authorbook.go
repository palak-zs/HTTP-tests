package Code

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetAll(w http.ResponseWriter, r *http.Request) {

	b, err := FetchAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

	bk, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error in marshalling.")
	}

	bytes.NewBuffer(bk)
	_, err = w.Write(bk)
	if err != nil {
		fmt.Errorf("Error: %d", err)
	}
}

func GetByID(w http.ResponseWriter, r *http.Request) {

	//var db *sql.DB
	//bk, _ := FetchAll()

	pathParam := mux.Vars(r)

	//for _, item := range bk {
	//	if item.ID == pathParam["ID"] {
	//		data, err := json.Marshal(item)
	//		if err != nil {
	//			fmt.Errorf("error: %d", err)
	//		}
	//
	//		bytes.NewBuffer(data)
	//		_, err = w.Write(data)
	//		if err != nil {
	//			fmt.Errorf("error: %d", err)
	//		}
	//	}
	//}
	Db := Connection()
	Row := Db.QueryRow("select* from books where book_id=?", pathParam["id"])
	var b Book
	err := Row.Scan(&b.ID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err != nil {
		fmt.Println("sdfasf")
	}
	_, Author := FetchAuthorID(b.AuthorID)
	b.Author = &Author

	data, _ := json.Marshal(b)
	w.Write(data)
}

func checkDOB(dob string) bool {

	DOB := strings.Split(dob, "-")
	day, _ := strconv.Atoi(DOB[0])
	month, _ := strconv.Atoi(DOB[1])
	year, _ := strconv.Atoi(DOB[2])

	switch {
	case day < 0 || day > 31:
		return false
	case month < 0 || month > 12:
		return false
	case year > 2015:
		return false
	}
	return true
}

func checkPublication(pub string) bool {
	strings.ToLower(pub)

	if pub == "Scholastic" || pub == "Arihanth" || pub == "Penguin" {
		return true
	}
	return false
}

func checkPublishedDate(pubDate string) bool {
	pd := strings.Split(pubDate, "-")
	day, _ := strconv.Atoi(pd[0])
	month, _ := strconv.Atoi(pd[1])
	year, _ := strconv.Atoi(pd[2])

	switch {
	case day < 0 || day > 31:
		return false
	case month < 0 || day > 12:
		return false
	case year < 1880 || year > 2022:
		return false
	}
	return true
}

func PostAuthor(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Errorf("Error: %d", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var author Author
	json.Unmarshal(data, &author)

	auth, _ := FetchAuthorID(author.AuthorID)
	if auth == author.AuthorID || author.FirstName == "" {
		fmt.Errorf("Error: %d", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var db *sql.DB
	Connection()
	InsertAuthor(db, author)
	if err != nil {
		fmt.Errorf("Error while inserting Author: %v", err)
	}

	w.Write(data)
	if err != nil {
		fmt.Errorf("Error while writing data of Author: %v", err)
	}
}

func PostBook(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error: %d", err)
	}

	var book Book
	json.Unmarshal(body, &book)
	if book.ID == "" || book.AuthorID <= 0 || book.Author.FirstName == "" || book.Title == "" {
		fmt.Println("Invalid credentials in book")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkDOB(book.Author.DOB) && !checkPublishedDate(book.PublishedDate) && !checkPublication(book.Publication) {
		fmt.Println("Invalid credentials in author")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var db *sql.DB
	//db := Connection()
	book, _ = InsertBook(db, book)

	fmt.Fprintf(w, "successfully inserted at id=%v", book.ID)

	w.Write(body)
	w.WriteHeader(http.StatusCreated)

}

func PutBook(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	param := mux.Vars(r)
	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return
	}

	var book Book
	json.Unmarshal(data, &book)

	id, author := FetchAuthorID(book.AuthorID)

	if id != book.AuthorID {
		fmt.Errorf("Author does not exist.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book.Author = &author

	db := Connection()

	if !checkPublishedDate(book.PublishedDate) && !checkPublication(book.Publication) || book.Title == "" && !checkDOB(book.Author.DOB) {
		fmt.Println("Invalid credentials.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var temp Book
	row := db.QueryRow("SELECT * from books where book_id=?", param["id"])
	if err = row.Scan(&temp.ID, &temp.AuthorID, &temp.Title, &temp.Publication, &temp.PublishedDate); err == nil {
		_, err = db.Exec("UPDATE books SET book_id=?, author_id=?, title=?, publication=?, published_date=? WHERE book_id=?",
			book.ID, book.AuthorID, book.Title, book.Publication, book.PublishedDate, param["id"])
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		_, err = db.Exec("INSERT into books(book_id, author_id, title, publication, published_date)values (?,?,?,?,?)",
			book.ID, book.AuthorID, book.Title, book.Publication, book.PublishedDate)

		fmt.Fprintf(w, "Successfully inserted id = %v\n", param["id"])
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func PutAuthor(w http.ResponseWriter, r *http.Request) {
	ReqData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("failed:%v\n", err)
		return
	}
	var author Author
	json.Unmarshal(ReqData, &author)

	params := mux.Vars(r)
	Db := Connection()

	if !checkDOB(author.DOB) {
		fmt.Println("no valid DOB")
		w.WriteHeader(http.StatusBadRequest)
	}

	id, _ := strconv.Atoi(params["id"])
	var checkExistingAuthor Author

	row := Db.QueryRow("select * from authors where author_id=?", id)
	if err = row.Scan(&checkExistingAuthor.AuthorID, &checkExistingAuthor.FirstName, &checkExistingAuthor.LastName,
		&checkExistingAuthor.DOB, &checkExistingAuthor.PenName); err == nil {
		_, err = Db.Exec("update authors set author_id=?,first_name=?,last_name=?,DOB=?,pen_name=? where author_id=?",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, id)

		fmt.Fprintf(w, "successfull updated id =%v\n", params["id"])
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		_, err = Db.Exec("insert into authors(author_id,first_name,last_name,DOB, pen_name)values(?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)

		fmt.Fprintf(w, "successfull inserted id =%v\n", params["id"])
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Errorf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id <= 0 {
		fmt.Errorf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := Connection()
	_ = db.QueryRow("delete from books where book_id=?", id)
	w.WriteHeader(http.StatusNoContent)

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Errorf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Author id")
		return
	}

	if id <= 0 {
		fmt.Errorf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Author")
		return
	}

	db := Connection()
	_ = db.QueryRow("delete from authors where author_id=?", id)
	w.WriteHeader(http.StatusNoContent)

}
