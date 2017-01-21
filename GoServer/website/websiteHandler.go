package website

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/qiniu/log"
)


//handles index landing page
func indexHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "home.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}

//displays a list of movie titles using tempalates
func moviesHandler(w http.ResponseWriter, r *http.Request){
	//create database connection
	c := db.Session.DB("codejam").C("movies")
	var results []interface{}

	//get all the results in database and map it to results
	err := c.Find(bson.M{}).All(&results)

	if err != nil{
		log.Fatal(err)
	}

	//load template with movie results
	err = templates.ExecuteTemplate(w, "movies.html", &results)

	if err != nil{
		log.Fatal(err)
	}
}