package app

import (
	"fmt"
	"go-practice/pkg/app/handlers"
	"go-practice/pkg/domain/services"
	"go-practice/pkg/infra/db"
	"go-practice/pkg/infra/repos"
	"net/http"
)

func Start() {
	db, err := db.NewDBContext("postgresql://postgres:Salmon123@localhost:5432/library?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	urepo := repos.NewUsersRepository(db)
	usvc := services.NewUsersService(urepo)
	uhndl := handlers.NewUsersHandler(usvc)

	brepo := repos.NewBooksRepository(db)
	bsvc := services.NewBooksService(brepo)
	bhndl := handlers.NewBooksHandler(bsvc)

	brrepo := repos.NewBorrowsRepository(db)
	brsvc := services.NewBorrowsService(brrepo)
	brhndl := handlers.NewBorrowsHandler(brsvc)

	http.HandleFunc("/users", uhndl.CreateUsersHandlerFunc)
	http.HandleFunc("/users/", uhndl.UsersHandlerFunc)
	http.HandleFunc("/books", bhndl.CreateBooksHandlerFunc)
	http.HandleFunc("/books/", bhndl.BooksHandlerFunc)
	http.HandleFunc("/borrows", brhndl.CreateBorrowHandlerFunc)
	http.HandleFunc("/borrows/", brhndl.BorrowHandlerFunc)
	// http.HandleFunc("/books", booksHandler)
	// http.HandleFunc("/borrows", borrowHandler)

	http.ListenAndServe(":8080", nil)
}
