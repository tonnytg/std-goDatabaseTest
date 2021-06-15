package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(IP:3306)/mysql")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Printf("Possivelmente o SQL está inalcansável ou a porta está bloqueada.")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Create a database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS test")
	if err != nil {
		panic(err)
	}
	db.Close()
	fmt.Println("Database test created!")

	// // Create a table
	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS test (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, valor TEXT NOT NULL)")
	// if err != nil {
	//     panic(err)
	// }

	// // perform a db.Query insert
	// insert, err := db.Query("INSERT INTO test VALUES ( 'TEST' )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()

}
