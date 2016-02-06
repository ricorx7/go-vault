package main

import (
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2/bson"

	"github.com/albrow/forms"
	"github.com/go-zoo/bone"
)

// Add the SalesOrder.
func salesOrderUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		val := bone.GetValue(r, "id")
		so := getSalesOrder(val)
		soData := &SalesOrderUpdate{}
		soData.SO = *so

		displaySalesOrderUpdateTemplate(w, soData)
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
		val.Require("SalesOrderNumber")

		// Use data to create a user object
		so := &SalesOrder{
			IsSelected:  formData.GetBool("IsSelected"),
			UserName:    formData.Get("UserName"),
			SalesPerson: formData.Get("SalesPerson"),
			POReference: formData.Get("POReference"),
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
		}

		fmt.Printf("Sales Order Update: %s\n", so.SalesOrderNumber)

	}
}

// Display the template
func displaySalesOrderUpdateTemplate(w http.ResponseWriter, soData *SalesOrderUpdate) {
	// Generate token
	soData.Token = genToken()

	// Redirect back to the page with error message
	t, _ := template.ParseFiles("header.html", "salesorder_update.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", soData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, soData)
}

// Find the ADCP from the database
func getSalesOrder(salesOrderNum string) *SalesOrder {
	fmt.Println("Get ADCP - SerialNum: ", salesOrderNum)

	var data SalesOrder
	err := Vault.Mongo.C("SalesOrder").Find(bson.M{"SalesOrderNumber": salesOrderNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find SalesOrder %v\n", err)
	}

	fmt.Println("SalesOrder: ", data.SalesOrderNumber)
	fmt.Println("SerialNum: ", data.SerialNumber)
	//fmt.Println("CustomerID: ", data.CustomerID)
	fmt.Println("ID:", data.ID)

	return &data
}
