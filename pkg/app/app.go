package app

import (
	"fmt"
	"go-practice/pkg/app/handlers"
	"go-practice/pkg/domain/services"
	"go-practice/pkg/infra/db"
	"go-practice/pkg/infra/repos"
	"go-practice/pkg/infra/rmq"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		panic("error loading .env file")
	}

	connString, ok := os.LookupEnv("DB_CONN_STRING")
	if !ok {
		panic("DB conn string not provided")
	}
	rmqConnString, ok := os.LookupEnv("RMQ_CONN_STRING")
	if !ok {
		panic("RMQ conn string not provided")
	}

	dB, err := db.NewDBContext(connString)
	if err != nil {
		fmt.Println(err)
		return
	}

	txHandler := db.NewTransactionHandler(dB)
	rmqConn := rmq.NewRMQConnection(rmqConnString)

	urepo := repos.NewUsersRepository(dB)
	brepo := repos.NewBooksRepository(dB)
	brrepo := repos.NewBorrowsRepository(dB)

	usvc := services.NewUsersService(urepo, brrepo, rmqConn)
	bsvc := services.NewBooksService(brepo, brrepo, rmqConn)
	brsvc := services.NewBorrowsService(urepo, brepo, brrepo, txHandler, rmqConn)

	uhndl := handlers.NewUsersHandler(usvc)
	bhndl := handlers.NewBooksHandler(bsvc)
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
