package main

import (
	"fmt"
	"net/http"

	//	"html/template"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name, err := os.Getwd()
	checkErr(err)
	fmt.Fprintf(w, "HOSTNAME : %s\n", name)
}

func main() {
	db, err := sql.Open("sqlite3", "../../sqlite/tmm.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM years_served")
	checkErr(err)

	fmt.Println(rows)

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
