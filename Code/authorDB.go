package Code

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

//func CreateAuthorTable() {
//	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")
//
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println(db.Ping())
//	_, err = db.Exec("USE testhttp")
//
//	stmt, err := db.Prepare("CREATE Table Author(auth_id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(50), dob varchar(50), pen_name varchar(50), primary key (auth_id));")
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

func InsertAuthor(db *sql.DB, auth Author) (Author, error) {

	//CreateAuthorTable()
	db = Connection()
	auth1 := auth
	//auth1.ID = -1

	db, err := sql.Open("mysql", "root:Abc!@#07@/testhttp")

	if err != nil {
		return auth1, errors.New("Invalid key!")
	}

	sql := "INSERT INTO Author(auth_id, first_name, last_name,dob,pen_name) values (?,?,?,?,?)"

	_, err = db.Exec(sql, auth.AuthorID, auth.FirstName, auth.LastName, auth.DOB, auth.PenName)
	if err != nil {
		return auth1, errors.New("Key already present.")
	}

	//auth.AuthorID, err = res.LastInsertId()
	//if err != nil {
	//	return auth1, errors.New("Key inserted.")
	//}

	//fmt.Printf("The last inserted row ID: %d\n", lastId)
	return auth, nil
}

func FetchData(id int) (Author, error) {
	auth := Author{}

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
	row := db.QueryRow("SELECT * FROM Author WHERE auth_id = ?", id)

	if err := row.Scan(&auth.AuthorID, &auth.FirstName, &auth.LastName, &auth.DOB, &auth.PenName); err != nil {
		return auth, fmt.Errorf("GetId %d: %v", id, err)
	}
	return auth, nil
}
