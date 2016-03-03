package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/albrow/forms"
)

func adcpAddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method adcpAddHander:", r.Method) // get request method
	if r.Method == "GET" {

		// Init data
		adcp := &Adcp{}
		adcpData := &AdcpDataUpdate{Adcp: *adcp}

		displayTemplate(w, adcpData)
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
		val.Require("SerialNumber")
		val.LengthRange("SerialNumber", 16, 16)

		// Use data to create a user object
		adcp := &Adcp{
			SerialNumber:          formData.Get("SerialNumber"),
			Frequency:             formData.Get("Frequency"),
			Customer:              formData.Get("Customer"),
			OrderNumber:           formData.Get("OrderNumber"),
			Application:           formData.Get("Application"),
			ConnectorType:         formData.Get("ConnectorType"),
			DepthRating:           formData.Get("DepthRating"),
			Firmware:              formData.Get("Firmware"),
			Hardware:              formData.Get("Hardware"),
			HeadType:              formData.Get("HeadType"),
			Modified:              time.Now(),
			Created:               time.Now(),
			PressureSensorPresent: formData.GetBool("PressureSensorPresent"),
			PressureSensorRating:  formData.Get("PressureSensorRating"),
			RecorderFormatted:     formData.GetBool("RecorderFormatted"),
			RecorderSize:          formData.Get("RecorderSize"),
			Software:              formData.Get("Software"),
			SystemType:            formData.Get("SystemType"),
			TemperaturePresent:    formData.GetBool("TemperaturePresent"),
		}
		adcpData := &AdcpDataUpdate{
			Adcp: *adcp,
		}

		// Check for errors and print the error message
		if val.HasErrors() {
			fmt.Println("Error with value:")
			for err := range val.ErrorMap() {
				fmt.Print(err)
				fmt.Print(":")
				fmt.Println(val.ErrorMap()[err])
			}

			// Set the error values
			adcpData.Errors.SerialNumber = val.ErrorMap()["SerialNumber"][0]

			// Display the template
			displayTemplate(w, adcpData)
		} else if checkSerialNumber(adcp.SerialNumber) {
			adcpData.Errors.SerialNumber = "Serial Number already used."
			// Display template
			displayTemplate(w, adcpData)

		} else {
			fmt.Println("SerialNumber:", adcpData.Adcp.SerialNumber)
			fmt.Println("Customer:", adcpData.Adcp.Customer)

			// Add the ADCP to the DB
			err := Vault.Mongo.C("adcps").Insert(adcpData.Adcp)
			CheckError(err)

			// Add the ADCP checklist to the DB
			checklist := &AdcpChecklist{
				SerialNumber: adcpData.Adcp.SerialNumber,
				Modified:     time.Now(),
				Created:      time.Now(),
			}
			err = Vault.Mongo.C("AdcpChecklists").Insert(checklist)
			CheckError(err)

			// Go to the list of ADCP
			http.Redirect(w, r, "/adcp", http.StatusFound)
		}
	}
}

// Check if the serial number is already used.
// If it is, send true, if it is a new serial number, send false.
func checkSerialNumber(serialNum string) bool {
	var data []Adcp
	err := Vault.Mongo.C("adcps").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	CheckError(err)

	if len(data) > 0 {
		return true
	}

	return false
}

// Display the template
func displayTemplate(w http.ResponseWriter, adcpData *AdcpDataUpdate) {
	// Generate token
	adcpData.Token = genToken()

	// Redirect back to the page with error message
	t, _ := template.ParseFiles("header.html", "adcp_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", adcpData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, adcpData)
}
