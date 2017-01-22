package products

import (
	"os"
	"encoding/json"
	"fmt"
	"../database"
	"strings"
	"sort"
)

//config entries

type ProductList struct{
	Apple []string
	Samsung []string
	Lg []string
	Motorola []string
	Google []string
	Htc []string
	Huawai []string
}

type Product struct{
	FullName string
	Brand string
	Model string
	Num int
	AveragePrice float64
	LowestPrice float64
	HighestPrice float64
	AveragePriceStr string

}

func(p *Product) PriceStr(price float64) (string){
	return fmt.Sprintf("$%.2f", price)
}

func(p *Product) PriceStrFInt(price int) (string){
	return fmt.Sprintf("$%.2f", float64(price)/100.0)
}
type Products []Product

func (slice Products) Len() int{
	return len(slice)
}
func (slice Products) Less(i,j int) bool{
	return slice[i].Num > slice[j].Num
}
func (slice Products) Swap(i,j int){
	slice[i], slice[j] = slice[j], slice[i]
}

type Brand struct {
	Name string
	Num int
	AveragePrice float64
	LowestPrice float64
	HighestPrice float64

	AveragePriceStr string
}

var ProductStatList Products
var BrandList []Brand




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
	ProductStatList = append(ProductStatList, createItemStats(db, "motorola", products.Motorola)...)
	ProductStatList = append(ProductStatList, createItemStats(db, "google", products.Google)...)
	ProductStatList = append(ProductStatList, createItemStats(db, "htc", products.Htc)...)
	ProductStatList = append(ProductStatList, createItemStats(db, "huawai", products.Huawai)...)
	
	sort.Sort(ProductStatList)
	
	fmt.Println(ProductStatList)
	
	if err != nil{
		panic(err)
	}
	fmt.Println(products)

}

func createItemStats(db *database.Database, brand string, models []string) []Product{
	var pList []Product
	
	var brandNum int = 0;
	var brandAverage float64 = -1
	var brandLowest float64 = -1
	var brandHighest float64 = -1
	
	for _, model := range models{
		
		items, _ := db.GetItemBrandModel(brand, model)
		
		var lowestPrice float64 = -1
		var highestPrice float64 = -1
		
		var averagePrice = 0.0
		
		for _, item := range items{
			averagePrice += item.Price
			if(highestPrice < 0 || item.Price > highestPrice){
				highestPrice = item.Price
			}
			if(lowestPrice <0 || item.Price < lowestPrice){
				lowestPrice = item.Price
			}
		}
		
		if(lowestPrice > 0){
			if(lowestPrice < brandLowest){
				brandLowest = lowestPrice
			}	
		}
		
		if(highestPrice > 0){
			if(highestPrice > brandHighest){
				brandHighest = highestPrice
			}
		}
		
		
		itemNum := len(items)
		
		brandNum += itemNum
		brandAverage += averagePrice
		
		if(itemNum > 0) {
			averagePrice /= float64(len(items))
			pList = append(pList, Product{brand+" "+model, brand, model, itemNum,ConvertToRealPrice(averagePrice), ConvertToRealPrice(lowestPrice), ConvertToRealPrice(highestPrice), fmt.Sprintf("%.2f",(ConvertToRealPrice(averagePrice )))})
		}
		
	}
	if(brandNum > 0) {
		brandAverage /= float64(brandNum)
		BrandList = append(BrandList, Brand{brand, brandNum, ConvertToRealPrice(brandAverage), ConvertToRealPrice(brandLowest), ConvertToRealPrice(brandHighest), fmt.Sprintf("%.2f",(ConvertToRealPrice(brandAverage )))})
	}
	
	
	
	return pList
}

func ConvertToRealPrice(price float64) float64{
	return price/100.0
}