package website

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/qiniu/log"
	"fmt"
)


//handles index landing page
func indexHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "home.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "test1.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}

//displays a list of movie titles using tempalates
func itemHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	itemId := bson.ObjectIdHex(r.URL.Path[len("/item/"):])

	item, err := db.GetItemById(itemId)

	if err != nil{
		log.Fatal(err.Error())
	}

	//load template with movie results
	err = templates.ExecuteTemplate(w, "item.html", &item)

	if err != nil{
		log.Fatal(err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "about.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}