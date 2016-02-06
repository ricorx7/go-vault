package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// SalesOrderData holds the ADCP data.
type SalesOrderData struct {
	SO     []SalesOrder // ADCPs
	Filter string       // Sales Order filter
}

// List all the ADCPs.
func salesOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method salesorderListHandler:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		soData := &SalesOrderData{}

		// Get data form DB
		err := Vault.Mongo.C("SalesOrders").Find(bson.M{}).Sort("-created").All(&soData.SO)
		CheckError(err)
		fmt.Println("Number of SOs: ", len(soData.SO))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "salesorder_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", soData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, soData)
	} else {

		// Init
		soData := &SalesOrderData{}

		// Get the partial serial number
		formData, err := forms.Parse(r)
		CheckError(err)
		var partialSO = formData.Get("PartialSalesOrderNumber")
		soData.SO = *getSalesOrderContain(partialSO)
		fmt.Println("Number of SOs: ", len(soData.SO))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "salesorder_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", soData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, soData)
	}
}

// Find all the ADCPs with the serial number partial given.  This will filter the ADCPs down.
func getSalesOrderContain(soPartial string) *[]SalesOrder {
	var so []SalesOrder
	err := Vault.Mongo.C("SalesOrders").Find(bson.M{"SalesOrderNumber": bson.M{"$regex": soPartial}}).Sort("-created").All(&so)
	if err != nil {
		fmt.Printf("Can't find SalesOrder Partials %v\n", err)
	}

	return &so
}
