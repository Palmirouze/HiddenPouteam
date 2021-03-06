package database



import (

	"gopkg.in/mgo.v2"

	"fmt"
	"../config"
	"gopkg.in/mgo.v2/bson"
)

//hold database information
type Database struct{

	Session *mgo.Session

}

//Item struct that matches database listing
type Item struct{
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string
	Price float64
	Brand string
	Brandmodel string
	Source string
	Url string
}

//connects to database and return Database

func ConnectToDatabase() *Database{

	dbUrl := config.MainConfig.Database.Url

	fmt.Println("Connection to database on url "+dbUrl)



	session, err := mgo.Dial(dbUrl)

	if err != nil{

		panic(err)

	}



	database := Database{session}



	return &database

}

//close database session

func (db *Database) CloseDatabase(){

	db.Session.Close()

}