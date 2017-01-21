package website

import (
"net/http"
"fmt"
"encoding/json"
"github.com/qiniu/log"
	"gopkg.in/mgo.v2/bson"
)

//handles the api landing page
func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "THIS IS THE API")
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