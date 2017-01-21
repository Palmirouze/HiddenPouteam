package database



import (

	"gopkg.in/mgo.v2"

	"fmt"

)

//hold database information

type Database struct{

	Session *mgo.Session

}



//connects to database and return Database

func ConnectToDatabase(url string) *Database{

	fmt.Println("Connection to database on url "+url)



	session, err := mgo.Dial(url)

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