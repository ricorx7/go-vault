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

// Add the SalesOrder.
func salesOrderUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		salesorderNum := bone.GetValue(r, "id")
		so := getSalesOrder(salesorderNum)
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
		so := getSalesOrder(formData.Get("SalesOrderNumber"))
		so.IsSelected = formData.GetBool("IsSelected")
		so.UserName = formData.Get("UserName")
		so.SalesPerson = formData.Get("SalesPerson")
		so.POReference = formData.Get("POReference")
		//PODateIssued:             formData.get("PODateIssued"),
		//PODateReceived:           formData.Get("PODateReceived"),
		//DueDate:                  formData.Get("DueDate"),
		//EstShipDate:              formData.Get("EstShipDate"),
		so.Qty = formData.GetInt("Qty")
		so.UnitPrice = formData.Get("UnitPrice")
		so.TotalPrice = formData.Get("TotalPrice")
		so.EUS = formData.Get("EUS")
		so.BuildStatusPurchase = formData.Get("BuildStatusPurchase")
		so.BuildStatusPurchaseNotes = formData.Get("BuildStatusPurchaseNotes")
		so.BuildStatusXdcr = formData.Get("BuildStatusXdcr")
		so.BuildStatusXdcrNotes = formData.Get("BuildStatusXdcrNotes")
		so.BuildStatusAssembly = formData.Get("BuildStatusAssembly")
		so.BuildStatusAssemblyNotes = formData.Get("BuildStatusAssemblyNotes")
		so.BuildStatusTesting = formData.Get("BuildStatusTesting")
		so.BuildStatusTestingNotes = formData.Get("BuildStatusTestingNotes")
		//DateShiped:               formData.Get("DateShiped"),
		so.SerialNumber = formData.Get("SerialNumber")
		//DeliveryDate:             formData.Get("DeliveryDate"),
		so.Region = formData.Get("Region")
		so.BeamAngle = formData.Get("BeamAngle")
		so.ElectronicsType = formData.Get("ElectronicsType")
		so.BatteryType = formData.Get("BatteryType")
		so.NumBattery = formData.GetInt("NumBattery")
		so.ExternalBatteryPack = formData.Get("ExternalBatteryPack")
		so.CableLength = formData.Get("CableLength")
		so.InternalMemory = formData.Get("InternalMemory")
		so.InternalCompass = formData.Get("InternalCompass")
		so.DeckBox = formData.Get("DeckBox")
		so.PulseVersion = formData.Get("PulseVersion")
		so.UserVersion = formData.Get("UserVersion")
		so.ThirdParty = formData.Get("ThirdParty")
		so.Notes = formData.Get("Notes")
		so.Modified = time.Now().Local()

		fmt.Printf("Sales Order Update: %s\n", so.SalesOrderNumber)

		// Update the sales order in DB
		updateSalesOrder(so)

		// Go to the list of ADCP
		http.Redirect(w, r, "/so", http.StatusFound)
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
	fmt.Println("Get SalesOrder - SalesOrderNumber: ", salesOrderNum)

	var data SalesOrder
	err := Vault.Mongo.C("SalesOrders").Find(bson.M{"SalesOrderNumber": salesOrderNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find SalesOrder %v\n", err)
	}

	fmt.Println("SalesOrder: ", data.SalesOrderNumber)
	fmt.Println("SerialNum: ", data.SerialNumber)
	//fmt.Println("CustomerID: ", data.CustomerID)
	fmt.Println("ID:", data.ID)

	return &data
}

func updateSalesOrder(so *SalesOrder) {
	fmt.Println("updateSalesOrder - ID", so.ID)

	//err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp._id}, bson.M{"$inc": bson.M{"Customer": adcp.Customer}})
	err := Vault.Mongo.C("SalesOrders").Update(bson.M{"_id": so.ID}, bson.M{"$set": bson.M{
		"IsSelected":  so.IsSelected,
		"UserName":    so.UserName,
		"SalesPerson": so.SalesPerson,
		"POReference": so.POReference,
		//PODateIssued:             formData.get("PODateIssued"),
		//PODateReceived:           formData.Get("PODateReceived"),
		//DueDate:                  formData.Get("DueDate"),
		//EstShipDate:              formData.Get("EstShipDate"),
		"Qty":        so.Qty,
		"UnitPrice":  so.UnitPrice,
		"TotalPrice": so.TotalPrice,
		"EUS":        so.EUS,
		"BuildStatusPurchase":      so.BuildStatusPurchase,
		"BuildStatusPurchaseNotes": so.BuildStatusPurchaseNotes,
		"BuildStatusXdcr":          so.BuildStatusXdcr,
		"BuildStatusXdcrNotes":     so.BuildStatusXdcrNotes,
		"BuildStatusAssembly":      so.BuildStatusAssembly,
		"BuildStatusAssemblyNotes": so.BuildStatusAssemblyNotes,
		"BuildStatusTesting":       so.BuildStatusTesting,
		"BuildStatusTestingNotes":  so.BuildStatusTestingNotes,
		//DateShiped:               formData.Get("DateShiped"),
		"SerialNumber": so.SerialNumber,
		//DeliveryDate:             formData.Get("DeliveryDate"),
		"Region":              so.Region,
		"BeamAngle":           so.BeamAngle,
		"ElectronicsType":     so.ElectronicsType,
		"BatteryType":         so.BatteryType,
		"NumBattery":          so.NumBattery,
		"ExternalBatteryPack": so.ExternalBatteryPack,
		"CableLength":         so.CableLength,
		"InternalMemory":      so.InternalMemory,
		"InternalCompass":     so.InternalCompass,
		"DeckBox":             so.DeckBox,
		"PulseVersion":        so.PulseVersion,
		"UserVersion":         so.UserVersion,
		"ThirdParty":          so.ThirdParty,
		"Notes":               so.Notes,
		"Modified":            so.Modified}})
	if err != nil {
		fmt.Printf("Can't update SalesOrder %v\n", err)
	}
}
