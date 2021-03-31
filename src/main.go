package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/push", push).Methods("POST")
	router.HandleFunc("/pop", pop).Methods("DELETE")
	http.ListenAndServe(":8090", router)
}
func push(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stringValue := string(reqBody)
	//fmt.Println(stringValue)
	db := dbConn()
	_, er := db.Query("INSERT INTO stack (elements) VALUES (?)", stringValue)
	if err != nil {
		log.Fatal(er)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Element has been Pushed Succesfully"}`))
}
func pop(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	isElementPresent, e := db.Query("SELECT COUNT(*) FROM stack ")
	if e != nil {
		log.Fatal()
	}
	for isElementPresent.Next() {
		var rows int
		err := isElementPresent.Scan(&rows)
		if err != nil {
			log.Panic(err.Error())
		}
		if rows == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Sorry There are no elements in the stack to POP"}`))
		} else {
			_, err := db.Query("DELETE FROM stack ORDER BY Id DESC LIMIT 1")
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Element has been Poped Succesfully"}`))
		}
	}
}
func dbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:4b3@tcp(127.0.0.1:3306)/stackdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Database Schema
// CREATE TABLE stack (
//     Id int NOT NULL AUTO_INCREMENT,
//     elements varchar(255),
//		PRIMARY KEY(Id)
//);
