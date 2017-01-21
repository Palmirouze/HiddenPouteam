package database

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/qiniu/log"
	"../config"
	"errors"
)

func (db Database) GetItemById(id bson.ObjectId) (*Item, error){
	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	var item Item

	err := c.Find(bson.M{"_id":id}).One(&item)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return &item, nil
}

func (db *Database) GetItemNameContains( term string) (*[]Item, error){


	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	var items []Item

	err := c.Find(bson.M{"name":bson.M{"$regex":term}}).All(&items)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return &items, nil
}
