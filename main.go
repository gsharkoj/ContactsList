package main

import (
	"database/sql"
	"encoding/json"
	"html"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Contact struct {
	Name  string
	Url   string
	Color string
	Id    int
}

func main() {

	// статичные файлы: react + bootstrap
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// обработки
	http.HandleFunc("/data", data)
	http.HandleFunc("/save", save)
	http.HandleFunc("/del", del)
	http.HandleFunc("/", view)

	http.ListenAndServe(":8001", nil)
}

func view(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("public/index.html")
	tmpl.Execute(w, nil)
}

func data(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./contacts.db")
	defer db.Close()

	rows, _ := db.Query("SELECT Name, Phone, id as Id FROM clients order by Name")
	defer rows.Close()
	data := []Contact{}

	// чередующиеся виды цветов
	colors := [4]string{"success", "error", "info", "warning"}

	index := 0
	for rows.Next() {

		var st = Contact{}
		_ = rows.Scan(&st.Name, &st.Url, &st.Id)
		st.Color = colors[index]
		data = append(data, st)

		index++
		if index == len(colors) {
			index = 0
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func Add(name, url, color string) Contact {

	st := Contact{Name: name, Url: url, Color: color}
	return st
}

func save(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	name := html.EscapeString(r.FormValue("name"))
	phone := html.EscapeString(r.FormValue("phone"))

	if len(name) > 0 && len(phone) > 0 {
		db, _ := sql.Open("sqlite3", "./contacts.db")
		defer db.Close()
		stmt, _ := db.Prepare("insert into clients(Name, Phone) values(?, ?)")
		defer stmt.Close()
		stmt.Exec(name, phone)
	}
}

func del(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	id := html.EscapeString(r.FormValue("id"))

	if html.EscapeString(id) != "" {
		db, _ := sql.Open("sqlite3", "./contacts.db")
		defer db.Close()
		db.Exec("delete from clients where id = $1", html.EscapeString(id))
	}
}