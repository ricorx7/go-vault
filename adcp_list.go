package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// AdcpData holds the ADCP data.
type AdcpData struct {
	Adcps  []Adcp // ADCPs
	Filter string // Serial Number filter
}

// List all the ADCPs.
func adcpListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method adcpListHander:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		adcpData := &AdcpData{}

		// Get data form DB
		err := Vault.Mongo.C("adcps").Find(bson.M{}).Sort("-created").All(&adcpData.Adcps)
		CheckError(err)
		fmt.Println("Number of ADCPs: ", len(adcpData.Adcps))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "adcp_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpData)
	} else {

		// Init
		adcpData := &AdcpData{}

		// Get the partial serial number
		formData, err := forms.Parse(r)
		CheckError(err)
		var partialSerial = formData.Get("PartialSerialNumber")
		adcpData.Adcps = *getAdcpContain(partialSerial)
		fmt.Println("Number of ADCPs: ", len(adcpData.Adcps))

		// Display data to page
		t, _ := template.ParseFiles("header.html", "adcp_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpData)
	}
}
