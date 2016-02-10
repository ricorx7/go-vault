package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// RmaData holds the ADCP data.
type RmaData struct {
	RMA    []RMA  // RMAs
	Filter string // Filter
}

// List all the RMAs.
func rmaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method rmaListHandler:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		rmaData := &RmaData{}

		// Get data form DB
		err := Vault.Mongo.C("RMAs").Find(bson.M{}).Sort("-Modified").All(&rmaData.RMA)
		CheckError(err)
		fmt.Println("Number of RMAs: ", len(rmaData.RMA))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "rma_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", rmaData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, rmaData)
	} else {

		// Init
		rmaData := &RmaData{}

		// Get the partial serial number
		formData, err := forms.Parse(r)
		CheckError(err)
		var partialRMA = formData.Get("PartialRmaNumber")
		rmaData.RMA = *getRmaContain(partialRMA)
		fmt.Println("Number of RMAs: ", len(rmaData.RMA))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "rma_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", rmaData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, rmaData)
	}
}

// Find all the RMAs with the RMA number partial given.  This will filter the RMAs down.
func getRmaContain(rmaPartial string) *[]RMA {
	var rma []RMA
	err := Vault.Mongo.C("RMAs").Find(bson.M{"RmaNumber": bson.M{"$regex": rmaPartial}}).Sort("-created").All(&rma)
	if err != nil {
		fmt.Printf("Can't find RMA Partials %v\n", err)
	}

	return &rma
}
