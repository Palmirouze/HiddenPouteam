package products

import (
	"os"
	"encoding/json"
	"fmt"
	"../database"
	"strings"
)

//config entries

type ProductList struct{
	Apple []string
	Samsung []string
	Lg []string
}

type Product struct{
	FullName string
	Brand string
	Model string
	Num int
	AveragePrice float64
	LowestPrice float64
	HighestPrice float64
}

var ProductStatList []Product

//gets the config from json file
func SearchProducts(term string) []Product{
	terms := strings.Split(term, " ")
	
	var list []Product
	
	for _, product := range ProductStatList{
		shouldAdd := true
		for _, t := range terms {
			if (!strings.Contains(strings.ToLower(product.FullName), strings.ToLower(t))){
				shouldAdd = false
			}
		}
		if(shouldAdd) {
			list = append(list, product)
		}
	}
	return list
}

func SetupProducts(db *database.Database){

	file, _ := os.Open("products.json")

	decoder := json.NewDecoder(file)

	products := ProductList{}
	
	err := decoder.Decode(&products)

	ProductStatList = createItemStats(db, "samsung", products.Samsung)
	ProductStatList = append(ProductStatList, createItemStats(db, "apple", products.Apple)...)
	ProductStatList = append(ProductStatList, createItemStats(db, "lg", products.Lg)...)
	fmt.Println(ProductStatList)
	
	if err != nil{
		panic(err)
	}
	fmt.Println(products)

}

func createItemStats(db *database.Database, brand string, models []string) []Product{
	var pList []Product
	
	for _, model := range models{
		
		items, _ := db.GetItemBrandModel(brand, model)
		
		var lowestPrice int = -1
		var highestPrice int = -1
		
		var averagePrice = 0
		
		for _, item := range items{
			averagePrice += item.Price
			if(highestPrice < 0 || item.Price > highestPrice){
				highestPrice = item.Price
			}
			if(lowestPrice <0 || item.Price < lowestPrice){
				lowestPrice = item.Price
			}
		}
		itemNum := len(items)
		if(itemNum > 0) {
			averagePrice /= len(items)
		}
		pList = append(pList, Product{brand+" "+model, brand, model, itemNum,float64(averagePrice)/100.0, float64(lowestPrice)/100.0, float64(highestPrice)/100.0})
	}
	
	return pList
}
