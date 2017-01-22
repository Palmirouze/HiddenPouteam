package website

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/qiniu/log"
	"../products"
	"../database"
)


//handles index landing page
func indexHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "home.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "test1.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}

//displays a list of movie titles using tempalates
func itemHandler(w http.ResponseWriter, r *http.Request){
	itemId := bson.ObjectIdHex(r.URL.Path[len("/item/"):])

	item, err := db.GetItemById(itemId)

	if err != nil{
		log.Fatal(err.Error())
	}

	//load template with movie results
	err = templates.ExecuteTemplate(w, "item.html", &item)

	if err != nil{
		log.Fatal(err) 
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "about.html", nil)

	if err != nil{
		log.Fatal(err)
	}
}
func viewProduct(w http.ResponseWriter, r *http.Request) {
	fullName := r.URL.Path[len("/view/"):]
	
	items, _ := db.GetItemBrandModelFull(fullName)
	
	product := products.SearchProducts(fullName)[0]
	
	type displayInterface struct{
		Items []database.Item
		Name string
		Product products.Product
	}
	err := templates.ExecuteTemplate(w, "productDisplay.html", displayInterface{items, fullName,product})

	if err != nil{
		log.Fatal(err)
	}
}
func productResultHandler(w http.ResponseWriter, r *http.Request) {
	var searchTerm string

	params := r.URL.Query()

	if val, ok := params["q"]; ok{
		searchTerm = val[0]
	}else{
		searchTerm = ""
	}
	products := products.SearchProducts(searchTerm)

	err := templates.ExecuteTemplate(w, "productSearchResult.html", products)
	
	if err != nil{
		log.Fatal(err)
	}
}

func searchResultHandler(w http.ResponseWriter, r *http.Request){
	var searchTerm string

	params := r.URL.Query()

	if val, ok := params["q"]; ok{
		searchTerm = val[0]
	}else{
		searchTerm = ""
	}
	
	items, err := db.GetItemNameContains(searchTerm)


	if err != nil{
		log.Fatal(err)
	}

	err = templates.ExecuteTemplate(w, "searchResult.html", items)



	if err != nil{
		log.Fatal(err)
	}
}

func brandStatHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "brandStats.html", products.BrandList)

	if err != nil{
		log.Fatal(err)
	}
}