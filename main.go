package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"wallet-server/config"
	"wallet-server/helpers"
	"wallet-server/internal/client"
	"wallet-server/internal/models"
	"wallet-server/routes"
)

func main() {

	user := helpers.CheckMainStringVars(os.Getenv("DB_USER"), "root")
	pwd := helpers.CheckMainStringVars(os.Getenv("DB_PASSWORD"), "password")
	dbName := helpers.CheckMainStringVars(os.Getenv("DB_NAME"), "bcs")

	addr := flag.String("addr", ":4000", "Server Network Address")

	dsn := flag.String("dsn",
		fmt.Sprintf("%s:%s@/%s?parseTime=true",
			user,
			pwd,
			dbName), "Data Source Name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	db, err := client.DbConn(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			errorLog.Fatal(err)
		}
	}(db)

	app := &config.Application{
		ErrorLog:    errorLog,
		InfoLog:     infoLog,
		Wallet:      &models.WalletModel{DB: db},
		Transaction: &models.TransactionModel{DB: db},
		User:        &models.UserModel{DB: db},
	}

	server := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  routes.Router(app),
	}

	err = server.ListenAndServe()

	if err != nil {
		errorLog.Fatalf("error running server: %s", err)
	}
	infoLog.Printf("server listening on %s", *addr)

}
