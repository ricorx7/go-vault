package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"
	"github.com/go-zoo/bone"
)

// AdcpDataUpdate holds the ADCP data.
type AdcpDataUpdate struct {
	Adcp   Adcp
	Token  string
	Errors struct {
		SerialNumber string
	}
}

// Update the ADCP data
func adcpUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		val := bone.GetValue(r, "id")
		adcp := getAdcp(val)

		// Generate a token
		token := genToken()

		//var adcpData = {"SerialNumber": "1000056", "Desc": "1200 kHz", "token":token}
		//adcp := &Adcp{SerialNumber: "100001212", Frequency: "1200 kHz"}
		adcpData := &AdcpDataUpdate{Adcp: *adcp, Token: token}

		t, _ := template.ParseFiles("header.html", "adcp_update.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpData)
	} else {
		// Parse the form
		formData, err := forms.Parse(r)
		CheckError(err)

		fmt.Println("Update the ADCP")

		// Check token
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
			fmt.Println("Good token")
		} else {
			// give error if no token
			fmt.Println("Bad token")
		}

		// Use data to create a user object
		adcp := getAdcp(formData.Get("SerialNumber"))
		adcp.Customer = formData.Get("Customer")
		adcp.OrderNumber = formData.Get("OrderNumber")
		adcp.Application = formData.Get("Application")
		adcp.ConnectorType = formData.Get("ConnectorType")
		adcp.DepthRating = formData.Get("DepthRating")
		adcp.Firmware = formData.Get("Firmware")
		adcp.Hardware = formData.Get("Hardware")
		adcp.HeadType = formData.Get("HeadType")
		adcp.Modified = time.Now()
		adcp.PressureSensorPresent = formData.GetBool("PressureSensorPresent")
		adcp.PressureSensorRating = formData.Get("PressureSensorRating")
		adcp.RecorderFormated = formData.GetBool("RecorderFormated")
		adcp.RecorderSize = formData.Get("RecorderSize")
		adcp.Software = formData.Get("Software")
		adcp.SystemType = formData.Get("SystemType")
		adcp.TemperaturePresent = formData.GetBool("TemperaturePresent")

		fmt.Println("Serial: ", adcp.SerialNumber)
		fmt.Println("Customer: ", adcp.Customer)
		fmt.Println("OrderNumber: ", adcp.OrderNumber)

		// Save the data to the DB
		updateAdcp(adcp)

		// Go to the list of ADCP
		http.Redirect(w, r, "/adcp", http.StatusFound)
	}
}
