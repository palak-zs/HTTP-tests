package main

import (
	"AuthorBooksHTTP/Code"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"AuthorBooksHTTP/Database"
)

func main() {
	r := mux.NewRouter()

	//h := func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "hello")
	//}
	//r.HandleFunc("/books", h).Methods(http.MethodGet)
	r.HandleFunc("/books", Code.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", Code.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/books", Code.PostBook).Methods(http.MethodPost)
	r.HandleFunc("/authors", Code.PostAuthor).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", Code.PutBook).Methods(http.MethodPut)
	r.HandleFunc("/authors/{id}", Code.PutAuthor).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", Code.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/authors/{id}", Code.DeleteAuthor).Methods(http.MethodDelete)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("server started on 8080")
}
