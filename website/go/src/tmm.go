package main

import (
//	"fmt"
	"net/http"
//	"html"
	"html/template"
	"database/sql"
//	"os"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

type Candidate struct {
	Precinct, First, Middle, Last, Streetnum, Streetname, Unitnum, Yearsserved, Reelection string
}

type BallotPrecinctData struct {
	PrecinctNumber string
	Candidates []Candidate
}

type FakeData struct {
	Stuff string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("index_layout.html"))

	data := FakeData{
		Stuff: "",
	}

    tmpl.Execute(w, data)
}

func AllPrecinctHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("all_precinct_layout.html"))

	db, err := sql.Open("sqlite3", "../../sqlite/tmm.db")
	checkErr(err)

	precinct_num := vars["id"]
	rows, err := db.Query("SELECT * FROM years_served where \"Precinct\" =" + precinct_num)

	var precinct, first, middle, last, streetnum, streetname, unitnum, yearsserved, reelection string
	

	candidate := Candidate{}
	candidates := []Candidate{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&precinct, &first, &middle, &last, &streetnum, &streetname, &unitnum, &yearsserved, &reelection)
		if err != nil {
			log.Fatal(err)
		}



		candidate.Precinct = precinct
		candidate.First = first
		candidate.Middle = middle
		candidate.Last = last
		candidate.Streetnum = streetnum
		candidate.Streetname = streetname
		candidate.Unitnum = unitnum
		candidate.Yearsserved = yearsserved
		candidate.Reelection = reelection
		candidates = append(candidates, candidate)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	data := BallotPrecinctData{
		PrecinctNumber: precinct_num,
		Candidates: candidates,
	}

    tmpl.Execute(w, data)
}

func BallotPrecinctHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("ballot_precinct_layout.html"))

	db, err := sql.Open("sqlite3", "../../sqlite/tmm.db")
	checkErr(err)

	precinct_num := vars["id"]
	rows, err := db.Query("SELECT * FROM years_served where \"Seeking Re-election 2021.04.06\" = 'TRUE' and \"Precinct\" =" + precinct_num)

	var precinct, first, middle, last, streetnum, streetname, unitnum, yearsserved, reelection string
	

	candidate := Candidate{}
	candidates := []Candidate{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&precinct, &first, &middle, &last, &streetnum, &streetname, &unitnum, &yearsserved, &reelection)
		if err != nil {
			log.Fatal(err)
		}



		candidate.Precinct = precinct
		candidate.First = first
		candidate.Middle = middle
		candidate.Last = last
		candidate.Streetnum = streetnum
		candidate.Streetname = streetname
		candidate.Unitnum = unitnum
		candidate.Yearsserved = yearsserved
		candidate.Reelection = reelection
		candidates = append(candidates, candidate)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	data := BallotPrecinctData{
		PrecinctNumber: precinct_num,
		Candidates: candidates,
	}

    tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/all/precinct/{id:[1-8]}", AllPrecinctHandler)
	r.HandleFunc("/ballot/precinct/{id:[1-8]}", BallotPrecinctHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
