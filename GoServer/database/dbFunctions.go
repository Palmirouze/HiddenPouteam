package database

import (
	"gopkg.in/mgo.v2/bson"
	"../config"
	"errors"
	"strings"
	"log"
)
//Get an item by it's mongo db _id
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

//Get array of items that match the full model string exactly. Ex: apple iphone 6
func (db *Database) GetItemBrandModelFull(full string) ([]Item, error){
	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	var items []Item

	err := c.Find(bson.M{"brandmodel":bson.M{"$regex":bson.RegEx{"^"+full+"$", "i"}}}).All(&items)


	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return items, nil

}
//Get array from brand and model strings
func (db *Database) GetItemBrandModel(brand string, model string) ([]Item, error){
	
	return db.GetItemBrandModelFull(brand+" "+model)
	
}
//Get array of items where the title loselly matches the term string
func (db *Database) GetItemNameContains( term string) ([]Item, error){
	c := db.Session.DB(config.MainConfig.Database.Name).C(config.MainConfig.Database.Tables.Items)

	allTerms := strings.Split(term, " ")

	regex := ""

	//construct regex
	for _, t := range allTerms{
		regex+="(?=.*"+t+")"
	}
	var items []Item

	err := c.Find(bson.M{"title":bson.M{"$regex":bson.RegEx{regex, "i"}}}).All(&items)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not do query for item anme contains.")
	}

	return items, nil
}
//Get the average price for a name contains query
func (db *Database) GetAverageNameContains(term string) float64 {
	items, err := db.GetItemNameContains(term)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return GetAverageItemPrice(items)
}
//Get average price of an array of items
func GetAverageItemPrice(items []Item) float64{
	var average float64
	for _, item := range items{
		average += float64(item.Price)
	}
	return average / float64(len(items))
}