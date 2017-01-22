package website

import (
"net/http"
"fmt"
"encoding/json"
	"log"
	"gopkg.in/mgo.v2/bson"
	"../products"
)

//handles the api landing page
func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "api")
}

//Displays a list of users in json format
func searchApiHandler(w http.ResponseWriter, r *http.Request){

	var searchTerm string

	params := r.URL.Query()

	if val, ok := params["q"]; ok{
		searchTerm = val[0]
	}else{
		searchTerm = ""
	}

	items, err := db.GetItemNameContains(searchTerm)


	if err != nil {
		log.Fatal(err)
	}

	jsonText, err := json.Marshal(items)

	fmt.Fprint(w,string(jsonText ))
}


func itemApiHandler(w http.ResponseWriter, r *http.Request){
	itemId := bson.ObjectIdHex(r.URL.Path[len("/api/item/"):])

	item, err := db.GetItemById(itemId)

	if err != nil{
		log.Fatal(err.Error())
	}
	jsonText, err := json.Marshal(item)

	fmt.Fprint(w, string(jsonText))

}

func brandStatApiHandler(w http.ResponseWriter, r *http.Request){


	jsonText, err := json.Marshal(products.BrandList)
	
	if err != nil{
		log.Fatal(err.Error())
	}
	
	fmt.Fprint(w, string(jsonText))

}