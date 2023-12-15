package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	IncidentApi IncidentApi
	IocApi      IocApi
	WorklogApi  WorklogApi
	UserApi     UserApi
}

type Store interface {
	createTable()
	createHandles(*ApiServer)
}

func setupDB(adress string, user string, pw string, s *ApiServer) *DB {
	fmt.Println("Connecting to DB")
	var DB = new(DB)
	db_adress := user + ":" + pw + "@tcp(" + adress + ")/"
	fmt.Println(db_adress)
	db, err := sql.Open("mysql", db_adress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS contentdb")
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	db, err = sql.Open("mysql", db_adress+"contentdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Stores")

	CreateIncidentApi(db, s)
	CreateIocApi(db, s)
	CreateUserApi(db, s)
	CreateWorklogApi(db, s)
	CreateTaskApi(db, s)

	fmt.Println("DB Connection established")
	return DB
}
