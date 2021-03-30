package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var Stack []string

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/push", push).Methods("POST")
	router.HandleFunc("/pop", pop).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
func push(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringvalue := string(reqBody)
	fmt.Println(stringvalue)

	Stack = append(Stack, stringvalue)
	fmt.Println(Stack)
	//Need to push the stack into database

}
func pop(w http.ResponseWriter, r *http.Request) {
	Stack = []string{"10", "20", "30"} //Hardcoded as of now
	//*****NEED to get stack from DATABASE*****//
	if len(Stack) > 0 {
		n := len(Stack) - 1
		fmt.Println(Stack[n])
		Stack = Stack[:n]
		fmt.Println(Stack)
	} else {
		fmt.Println("STACK IS EMPTY", Stack)
	}
	// stmt, err := db.Prepare("delete from posts where id=?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// CREATE TABLE stack
// (
//     id SERIAL,
//     stack TEXT NOT NULL,
//)
