package main

import (
	"fmt"
	"net/http"

	"html/template"
	"database/sql"
	"os"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name, err := os.Getwd()
	checkErr(err)
	fmt.Fprintf(w, "HOSTNAME : %s\n", name)
}


func firstNames(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	db, err := sql.Open("sqlite3", "../../sqlite/tmm.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM years_served where \"Seeking Re-election 2021.04.06\" = 'TRUE'")

	var precinct, first, middle, last, streetnum, streetname, unitnum, yearsserved, reelection string
	
	var data []string
	

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&precinct, &first, &middle, &last, &streetnum, &streetname, &unitnum, &yearsserved, &reelection)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(precinct, first, middle, last, streetnum, streetname, unitnum, yearsserved, reelection)
		data = append(data, first)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

        tmpl.Execute(w, data)
    
}

func main() {


	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/firstnames", firstNames)
	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
