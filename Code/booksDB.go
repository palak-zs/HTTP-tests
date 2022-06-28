package Code

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")
	if err != nil {
		fmt.Println("1")
		panic(err.Error())
	}

	fmt.Println(db.Ping())

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("2")
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	return db
}

//func CreateBooksTable() {
//	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")
//
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println(db.Ping())
//	_, err = db.Exec("USE testhttp")
//
//	stmt, err := db.Prepare("CREATE Table books(book_id varchar(50) NOT NULL , author_id varchar(50), Title varchar(50), Author varchar(50), Publication varchar(50), published_date varchar(50), primary key (book_id));")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println("DB selected successfully..")
//	}
//
//	_, err = stmt.Exec()
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println("Table created successfully..")
//	}
//}

func InsertBook(db *sql.DB, book Book) (Book, error) {

	db = Connection()

	//CreateBooksTable()
	book1 := book
	book1.ID = "-1"

	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")

	if err != nil {
		return book1, errors.New("Invalid key!")
	}

	_, err = db.Exec("INSERT INTO books(book_id, author_id, title, publication, published_date) values (?,?,?,?,?)",
		book.ID, book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		return book1, errors.New("Key already present.")
	}

	//book.ID, err = res.LastInsertId()
	//if err != nil {
	//	return book1, errors.New("Key inserted.")
	//}

	//fmt.Printf("The last inserted row ID: %d\n", lastId)
	return book, nil
}

func FetchAll() ([]Book, error) {
	db := Connection()
	defer db.Close()

	fmt.Println("before scan")
	var books []Book

	row, err := db.Query("SELECT * FROM books")
	defer row.Close()
	if err != nil {
		return books, errors.New("key not find in table")
	}

	fmt.Println(row)

	for row.Next() {
		var temp Book
		err := row.Scan(&temp.ID, &temp.AuthorID, &temp.Title, &temp.Publication, &temp.PublishedDate)
		if err != nil {
			return books, errors.New("some eror happen in recieving value")
		}
		_, Author := FetchAuthorID(temp.AuthorID)
		fmt.Println(Author)
		temp.Author = &Author
		books = append(books, temp)
		//row.NextResultSet()
	}

	//fmt.Println(books)

	return books, nil
}

//func FetchBook(db *sql.DB, ID int) (Book, error) {
//	book := Book{}
//
//	row := db.QueryRow("SELECT * FROM employee WHERE emp_id = ?", ID)
//
//	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Publication, &book.PublishedDate); err != nil {
//		return book, fmt.Errorf("GetId %d: %v", ID, err)
//	}
//	return book, nil
//}

func FetchAuthorID(id int) (int, Author) {
	db := Connection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM authors where author_id=?", id)

	var author Author

	if err := row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		fmt.Errorf("Failed %d: %v", id, err)
	}
	return author.AuthorID, author

}
