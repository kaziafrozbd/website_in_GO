package main

import (
	"fmt"
	"net/url"
	"os"
	"github.com/mateors/mcb"
)
var db *mcb.DB

type requestTable struct {
	ID int `json:"aid"`
	Name       string   `json:"name"`
	Company    string   `json:"company"`
	Email string   `json:"email"`
	Type       string   `json:"type"`
	Status int `json:"status"`
}

func init() {

	db = mcb.Connect("localhost", "kazi", "kazi123", false)

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main(){


        //How to insert into couchbase bucket
	var myData requestTable

	form := make(url.Values, 0)
	form.Add("bucket", "test") //bucket and collection-> namespace:bucket.scope.collection
	form.Add("aid", "request::100") //document ID
	form.Add("name", "Mostain Billah")
	form.Add("company", "36@12")
	form.Add("email", "Developer@gmail.com")
	
    	form.Add("type", "request") //what type of data or table name in general (SQL)

	p := db.Insert(form, &myData) //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

    	//How to retrieve from couchbase bucket (selected fields only)

    // 	pres := db.Query("SELECT aid,name,age,profession FROM master_erp WHERE type='participant'")
	// rows := pres.GetRows()

	// fmt.Println("Total Rows:", len(rows))
	// fmt.Println(rows)

    	//How to retrieve from couchbase bucket (All fields using *)

    // 	pres := db.Query("SELECT * FROM master_erp WHERE type='participant'")
	// rows := pres.GetBucketRows("master_erp") //bucketName as argument

	// fmt.Println("Total Rows:", len(rows))
	// fmt.Println(rows)

}