package AuthorBooksHTTP

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func CreateTable() {
	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db.Ping())
	_, err = db.Exec("USE test")

	stmt, err := db.Prepare("CREATE Table books(book_id int NOT NULL AUTO_INCREMENT, title varchar(50), author varchar(50), publication varchar(50), published_date varchar(50), primary key (book_id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}
}

func InsertData(book Book) (Book, error) {

	CreateTable()

	book1 := book
	book1.id = -1

	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")

	if err != nil {
		return book1, errors.New("Invalid key!")
	}

	sql := "INSERT INTO books(book_id, title, author, publication, published_date) values (?,?,?,?,?)"

	res, err := db.Exec(sql, book.id, book.title, book.author, book.publication, book.publishedDate)
	if err != nil {
		return book1, errors.New("Key already present.")
	}

	book.id, err = res.LastInsertId()
	if err != nil {
		return book1, errors.New("Key inserted.")
	}

	//fmt.Printf("The last inserted row id: %d\n", lastId)
	return book, nil
}

func FetchData(id int) (Book, error) {
	book := Book{}

	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db.Ping())

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	row := db.QueryRow("SELECT * FROM author WHERE book_id = ?", id)

	if err := row.Scan(&book.id, &book.title, &book.author, &book.publication, &book.publishedDate); err != nil {
		return book, fmt.Errorf("GetId %d: %v", id, err)
	}
	return book, nil
}
