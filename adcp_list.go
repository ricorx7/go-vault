package main

import (
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2/bson"
)

// AdcpData holds the ADCP data.
type AdcpData struct {
	Data []Adcp // ADCPs
}

// List all the ADCPs.
func adcpListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method adcpListHander:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		adcpData := &AdcpData{}

		// Get data form DB
		err := Vault.Mongo.C("adcps").Find(bson.M{}).Sort("-created").All(&adcpData.Data)
		CheckError(err)
		fmt.Println("Number of ADCPs: ", len(adcpData.Data))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "adcp_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpData)
	}
}
