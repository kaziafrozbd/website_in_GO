package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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
	
	// name := r.FormValue("name")
	// comapny := r.FormValue("company")
	// email := r.FormValue("email")

	// fmt.Println(name, comapny, email)
	// fmt.Fprintf(w,"received \n %s \n %s \n %s", name, comapny, email)
	r.ParseForm()
	for key, val := range r.Form{
		fmt.Println(key, val)
	}
	fmt.Fprintf(w,`ok`)
}