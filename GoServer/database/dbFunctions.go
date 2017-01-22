package database

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/qiniu/log"
	"../config"
	"errors"
	"strings"
	"fmt"
)

func (db *Database) GetItemById(id bson.ObjectId) (*Item, error){
	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	var item Item

	err := c.Find(bson.M{"_id":id}).One(&item)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return &item, nil
}

func (db *Database) GetItemNameContains( term string) ([]Item, error){
	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	allTerms := strings.Split(term, " ")

	regex := ""

	//construct regex
	for _, t := range allTerms{
		regex+="(?=.*"+t+")"
	}
	fmt.Println(regex)
	var items []Item

	err := c.Find(bson.M{"title":bson.M{"$regex":bson.RegEx{regex, "i"}}}).All(&items)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return items, nil
}

func (db *Database) GetAverageNameContains(term string) float64 {
	items, err := db.GetItemNameContains(term)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return GetAverageItemPrice(items)
}

func GetAverageItemPrice(items []Item) float64{
	var average float64
	for _, item := range items{
		average += float64(item.Price)
	}
	return average / float64(len(items))
}