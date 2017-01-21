package website


import (

"fmt"

"net/http"

"../database"

"html/template"

"github.com/qiniu/log"

)



var db *database.Database






var templates = template.Must(template.ParseFiles("templates/base.html", "templates/item.html", "templates/home.html", "templates/test1.html" ,"templates/item.html"))



//main function

func StartWebsite(){

	//load configuration



	//connect to database defined in config

	db = database.ConnectToDatabase()

	defer db.CloseDatabase()



	startHttpServer()

}



//starts the httpserver

func startHttpServer(){



	fmt.Println("Starting web server...")



	//register web handlers
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//api

	http.HandleFunc("/api/search", searchApiHandler)
	http.HandleFunc("/api/item/", itemApiHandler)
	http.HandleFunc("/api/", apiHandler)



	//main
	http.HandleFunc("/test1.html/", testHandler)
	http.HandleFunc("/item/", itemHandler)

	http.HandleFunc("/", indexHandler)



	//listen on port 8080

	err := http.ListenAndServe(":8080", nil)



	if err != nil{

		log.Fatal(err)

	}

}

