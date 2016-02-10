package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"
	"gopkg.in/mgo.v2/bson"
)

// Product will keep track of the product information.
type Product struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Desc          string        `bson:"Desc" json:"Desc"`                   // Description
	PartNumber    string        `bson:"PartNumber" json:"PartNumber"`       // Partnumber
	ReferenceDes  string        `bson:"ReferenceDes" json:"ReferenceDes"`   // Reference Designator
	UnitPrice     float64       `bson:"UnitPrice" json:"UnitPrice"`         // Unit Price
	ListPrice     float64       `bson:"ListPrice" json:"ListPrice"`         // List Price
	Type          string        `bson:"Type" json:"Type"`                   // Type of assembly
	Qty           int           `bson:"Qty" json:"Qty"`                     // Quantity of parts
	UnitOfMeasure string        `bson:"UnitOfMeasure" json:"UnitOfMeasure"` // Unit of Measure
	Cost          float64       `bson:"Cost" json:"Cost"`                   // Cost of product
	Created       time.Time     `bson:"Created" json:"Created"`             // Created
	Modified      time.Time     `bson:"Modified" json:"Modified"`           // Modified
}

// ProductUpdate will contain the RMA data.
type ProductUpdate struct {
	Product Product
	Token   string
}

// Add the Product.
func productAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Display blank form
		productData := &ProductUpdate{}
		displayProductTemplate(w, productData)
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
		val.Require("Desc")

		// Use data to create a user object
		product := &Product{
			PartNumber:    formData.Get("PartNumber"),
			Desc:          formData.Get("Desc"),
			ReferenceDes:  formData.Get("ReferenceDes"),
			UnitPrice:     formData.GetFloat("UnitPrice"),
			ListPrice:     formData.GetFloat("ListPrice"),
			Type:          formData.Get("Type"),
			Qty:           formData.GetInt("Qty"),
			UnitOfMeasure: formData.Get("UnitOfMeasure"),
			Cost:          formData.GetFloat("Cost"),
			Created:       time.Now().Local(),
		}

		// Add the Products to the DB
		err = Vault.Mongo.C("Products").Insert(product)
		CheckError(err)

		// Go to the list of products
		http.Redirect(w, r, "/product", http.StatusFound)
	}
}

// Display the template
func displayProductTemplate(w http.ResponseWriter, data *ProductUpdate) {
	// Generate token
	data.Token = genToken()

	t, _ := template.ParseFiles("header.html", "product_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", data)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, data)
}
