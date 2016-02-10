package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/albrow/forms"
	"github.com/go-zoo/bone"
)

// Update the Product.
func productUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		productNum := bone.GetValue(r, "id")
		product := getProduct(productNum)
		productData := &ProductUpdate{}
		productData.Product = *product

		displayProductUpdateTemplate(w, productData)
	} else {
		// Parse the form
		formData, err := forms.Parse(r)
		CheckError(err)

		// Check token
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
			fmt.Println("Good token")
		} else {
			// give error if no token
			fmt.Println("Bad token")
		}

		// Validate data
		val := formData.Validator()
		val.Require("PartNumber")

		// Use data to create a user object
		product := getProduct(formData.Get("PartNumber"))
		product.PartNumber = formData.Get("PartNumber")
		product.Desc = formData.Get("Desc")
		product.ReferenceDes = formData.Get("ReferenceDes")
		product.UnitPrice = formData.GetFloat("UnitPrice")
		product.ListPrice = formData.GetFloat("ListPrice")
		product.Type = formData.Get("Type")
		product.Qty = formData.GetInt("Qty")
		product.UnitOfMeasure = formData.Get("UnitOfMeasure")
		product.Cost = formData.GetFloat("Cost")
		product.Modified = time.Now().Local()

		fmt.Printf("Product Update: %s\n", product.PartNumber)

		// Update the Product in DB
		updateProduct(product)

		// Go to the list of Products
		http.Redirect(w, r, "/product", http.StatusFound)
	}
}

// Display the template
func displayProductUpdateTemplate(w http.ResponseWriter, data *ProductUpdate) {
	// Generate token
	data.Token = genToken()

	// Redirect back to the page with error message
	t, _ := template.ParseFiles("header.html", "product_update.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", data)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, data)
}

// Find the Product from the database
func getProduct(partNum string) *Product {
	fmt.Println("Get Product - PartNumber: ", partNum)

	var data Product
	err := Vault.Mongo.C("Products").Find(bson.M{"PartNumber": partNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find Product %v\n", err)
	}

	fmt.Println("Product: ", data.PartNumber)
	fmt.Println("ID:", data.ID)

	return &data
}

func updateProduct(product *Product) {
	fmt.Println("updateProduct - ID", product.ID)

	err := Vault.Mongo.C("Products").Update(bson.M{"_id": product.ID}, bson.M{"$set": bson.M{
		"PartNumber":    product.PartNumber,
		"Desc":          product.Desc,
		"ReferenceDes":  product.ReferenceDes,
		"UnitPrice":     product.UnitPrice,
		"ListPrice":     product.ListPrice,
		"Type":          product.Type,
		"Qty":           product.Qty,
		"UnitOfMeasure": product.UnitOfMeasure,
		"Cost":          product.Cost,
		"Modified":      product.Modified}})
	if err != nil {
		fmt.Printf("Can't update Product %v\n", err)
	}
}
