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
func userHandler(w http.ResponseWriter, r *http.Request){

	c := db.Session.DB("codejam").C("users")

	var users []User

	var err error
	//This query finds every user with a name that starts with F
	err = c.Find(bson.M{"name":bson.M{"$regex":"^F"}}).All(&users)

	if err != nil {
		log.Fatal(err)
	}
	jsonText, err := json.Marshal(users)




	fmt.Println(users)

	fmt.Fprint(w,string(jsonText ))
}

