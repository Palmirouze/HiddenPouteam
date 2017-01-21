package website


import (

"fmt"

"net/http"

"../config"

"../database"

"html/template"

"github.com/qiniu/log"

)



var db *database.Database



var configuration *config.Configuration



var templates = template.Must(template.ParseFiles("templates/base.html", "templates/movies.html", "templates/home.html"))



//main function

func StartWebsite(){

	//load configuration

	configuration = config.GetConfig("config.json")



	//connect to database defined in config

	db = database.ConnectToDatabase(configuration.Database.Url)

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

	http.HandleFunc("/api/users/", userHandler)

	http.HandleFunc("/api/", apiHandler)



	//main

	http.HandleFunc("/movies/", moviesHandler)

	http.HandleFunc("/", indexHandler)



	//listen on port 8080

	err := http.ListenAndServe(":8080", nil)



	if err != nil{

		log.Fatal(err)

	}

}

