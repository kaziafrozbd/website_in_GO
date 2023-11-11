package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error
func init(){
	
    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err = sql.Open("mysql", "root:hosting123@tcp(127.0.0.1:3306)/webhosting_db")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    // defer db.Close()

	
    
	fmt.Println("db connected..")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	// http.Handle("./", http.FileServer(http.Dir("./assets")))
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8888", nil)
}
func home(w http.ResponseWriter, r *http.Request){
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)

	// fmt.Fprintf(w,`welcome`)
}
func features(w http.ResponseWriter, r *http.Request){
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)

	// fmt.Fprintf(w,`welcome`)
}
func docs(w http.ResponseWriter, r *http.Request){
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err!=nil{
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)

	// fmt.Fprintf(w,`welcome`)
}
func request(w http.ResponseWriter, r *http.Request){
	
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	// fmt.Println(name, company, email)
	// fmt.Fprintf(w,"received \n %s \n %s \n %s", name, company, email)
	// r.ParseForm()
	// for key, val := range r.Form{
	// 	fmt.Println(key, val)
	// }

	// perform a db.Query insert
	sq := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1')"
	sql := fmt.Sprintf(sq, name, company, email)

	// fmt.Println(sql)
    insert, err := db.Query(sql)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
	fmt.Fprintf(w,`ok`)
}