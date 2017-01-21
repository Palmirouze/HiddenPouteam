package website

import (
"net/http"
"fmt"
"gopkg.in/mgo.v2/bson"
"encoding/json"
"github.com/qiniu/log"
)
type User struct{
	Name string
	Genres []string
}
//handles the api landing page
func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "THIS IS THE API")
}

//Displays a list of users in json format
func searchApiHandler(w http.ResponseWriter, r *http.Request){

	c := db.Session.DB(configuration.Database.Name).C("Items")

	searchTerm := r.URL.Query()["q"][0]
	fmt.Println(searchTerm)
	var items []interface{}

	var err error
	//This query finds every user with a name that starts with F
	err = c.Find(bson.M{"name":bson.M{"$regex":searchTerm}}).All(&items)

	if err != nil {
		log.Fatal(err)
	}
	jsonText, err := json.Marshal(items)




	fmt.Println(items)

	fmt.Fprint(w,string(jsonText ))
}

