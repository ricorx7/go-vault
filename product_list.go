package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// ProductData holds the ADCP data.
type ProductData struct {
	Products []Product // Products
	Filter   string    // Serial Number filter
}

// List all the Products.
func productListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method productListHander:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		data := &ProductData{}

		// Get data form DB
		err := Vault.Mongo.C("Products").Find(bson.M{}).Sort("-created").All(&data.Products)
		CheckError(err)
		fmt.Println("Number of Products: ", len(data.Products))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "product_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", data)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, data)
	} else {

		// Init
		data := &ProductData{}

		// Get the partial serial number
		formData, err := forms.Parse(r)
		CheckError(err)
		var partial = formData.Get("PartialProduct")
		data.Products = *getProductContain(partial)
		fmt.Println("Number of Products: ", len(data.Products))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "product_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", data)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, data)
	}
}

// Find all the Products with the Products number partial given.  This will filter the Products down.
func getProductContain(partial string) *[]Product {
	var data []Product
	err := Vault.Mongo.C("Products").Find(bson.M{"PartNumber": bson.M{"$regex": partial}}).Sort("-created").All(&data)
	if err != nil {
		fmt.Printf("Can't find Product Partials %v\n", err)
	}

	return &data
}

// Get a list of all the products.
func getProductList() *[]Product {
	var data []Product
	// Get data form DB
	err := Vault.Mongo.C("Products").Find(bson.M{}).Sort("-created").All(&data)
	if err != nil {
		fmt.Printf("Can't find Product Partials %v\n", err)
	}

	return &data
}

// Find a product.
func getProductPartNumber(partNum string) *Product {
	var data Product
	// Get data form DB
	//err := Vault.Mongo.C("Products").Find(bson.M{"_id": id.Hex()}).One(data)
	err := Vault.Mongo.C("Products").Find(bson.M{"PartNumber": partNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find Product by Partnumber %v\n", err)
		return nil
	}

	return &data
}
