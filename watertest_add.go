package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"
)

// WaterTestUpdate will contain the Water Test data.
type WaterTestUpdate struct {
	WaterTest WaterTestResults
	Token     string
	IsSelected string
	IsSelectedList                []OptionItem
}

// Add the Water Test.
func watertestAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		wtData := &WaterTestUpdate{}

		displayWaterTestTemplate(w, wtData)
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

		// Use data to create a user object
		wt := &SalesOrder{
			SalesOrderNumber: formData.Get("SalesOrderNumber"),
			IsSelected:       formData.GetBool("IsSelected"),
			UserName:         formData.Get("UserName"),
			SalesPerson:      formData.Get("SalesPerson"),
			POReference:      formData.Get("POReference"),
			//PODateIssued:             formData.get("PODateIssued"),
			//PODateReceived:           formData.Get("PODateReceived"),
			//DueDate:                  formData.Get("DueDate"),
			//EstShipDate:              formData.Get("EstShipDate"),
			Qty:                      formData.GetInt("Qty"),
			UnitPrice:                formData.Get("UnitPrice"),
			TotalPrice:               formData.Get("TotalPrice"),
			EUS:                      formData.Get("EUS"),
			BuildStatusPurchase:      formData.Get("BuildStatusPurchase"),
			BuildStatusPurchaseNotes: formData.Get("BuildStatusPurchaseNotes"),
			BuildStatusXdcr:          formData.Get("BuildStatusXdcr"),
			BuildStatusXdcrNotes:     formData.Get("BuildStatusXdcrNotes"),
			BuildStatusAssembly:      formData.Get("BuildStatusAssembly"),
			BuildStatusAssemblyNotes: formData.Get("BuildStatusAssemblyNotes"),
			BuildStatusTesting:       formData.Get("BuildStatusTesting"),
			BuildStatusTestingNotes:  formData.Get("BuildStatusTestingNotes"),
			//DateShiped:               formData.Get("DateShiped"),
			SerialNumber: formData.Get("SerialNumber"),
			//DeliveryDate:             formData.Get("DeliveryDate"),
			Region:              formData.Get("Region"),
			BeamAngle:           formData.Get("BeamAngle"),
			ElectronicsType:     formData.Get("ElectronicsType"),
			BatteryType:         formData.Get("BatteryType"),
			NumBattery:          formData.GetInt("NumBattery"),
			ExternalBatteryPack: formData.Get("ExternalBatteryPack"),
			CableLength:         formData.Get("CableLength"),
			InternalMemory:      formData.Get("InternalMemory"),
			InternalCompass:     formData.Get("InternalCompass"),
			DeckBox:             formData.Get("DeckBox"),
			PulseVersion:        formData.Get("PulseVersion"),
			UserVersion:         formData.Get("UserVersion"),
			ThirdParty:          formData.Get("ThirdParty"),
			Notes:               formData.Get("Notes"),
			Created:             time.Now().Local(),
		}

		// Add the SalesOrders to the DB
		err = Vault.Mongo.C("WaterTestResults").Insert(wt)
		CheckError(err)

		// Go to the list of WaterTests
		http.Redirect(w, r, "/watertest", http.StatusFound)
	}
}

// Display the template
func displayWaterTestTemplate(w http.ResponseWriter, wtData *WaterTestUpdate) {
	// Generate token
	wtData.Token = genToken()

	t, _ := template.ParseFiles("header.html", "watertest_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", wtData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, wtData)
}
