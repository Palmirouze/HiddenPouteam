package website


import (
"fmt"
"net/http"
"../database"
	"../products"
"html/template"
	"log"
)



var db *database.Database






var templates = template.Must(template.ParseFiles("templates/base.html",
	"templates/item.html", "templates/home.html",
	"templates/item.html", "templates/about.html", "templates/searchResult.html",
	"templates/productSearchResult.html", "templates/productDisplay.html", "templates/brandStats.html"))



//main function

func StartWebsite(){




	//connect to database defined in config

	db = database.ConnectToDatabase()

	defer db.CloseDatabase()

	products.SetupProducts(db)
	
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
	http.HandleFunc("/api/stats/brands/", brandStatApiHandler)
	http.HandleFunc("/api/", apiHandler)



	//main

	http.HandleFunc("/item/", itemHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/searchResult", searchResultHandler)
	http.HandleFunc("/productSearchResult", productResultHandler)
	http.HandleFunc("/view/", viewProduct)
	http.HandleFunc("/stats/", brandStatHandler)
	http.HandleFunc("/", indexHandler)



	//listen on port 8080

	err := http.ListenAndServe(":8080", nil)



	if err != nil{

		log.Fatal(err)

	}

}

