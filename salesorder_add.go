package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// SalesOrder will keep track of the sales order information.
type SalesOrder struct {
	ID               bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IsSelected       bool          `bson:"IsSelected" json:"IsSelected"`
	Created          time.Time     `bson:"Created" json:"Created"`
	UserName         string        `bson:"UserName" json:"UserName"`
	SalesOrderNumber string        `bson:"SalesOrderNumber" json:"SalesOrderNumber"`
	//CustomerID               bson.ObjectId
	SalesPerson              string    `bson:"SalesPerson" json:"SalesPerson"`
	POReference              string    `bson:"POReference" json:"POReference"`
	PODateIssued             time.Time `bson:"PODateIssued" json:"PODateIssued"`
	PODateReceived           time.Time `bson:"PODateReceived" json:"PODateReceived"`
	DueDate                  time.Time `bson:"DueDate" json:"DueDate"`
	EstShipDate              time.Time `bson:"EstShipDate" json:"EstShipDate"`
	Qty                      int       `bson:"Qty" json:"Qty"`
	UnitPrice                string    `bson:"UnitPrice" json:"UnitPrice"`
	TotalPrice               string    `bson:"TotalPrice" json:"TotalPrice"`
	EUS                      string    `bson:"EUS" json:"EUS"`
	BuildStatusPurchase      string    `bson:"BuildStatusPurchase" json:"BuildStatusPurchase"`
	BuildStatusPurchaseNotes string    `bson:"BuildStatusPurchaseNotes" json:"BuildStatusPurchaseNotes"`
	BuildStatusXdcr          string    `bson:"BuildStatusXdcr" json:"BuildStatusXdcr"`
	BuildStatusXdcrNotes     string    `bson:"BuildStatusXdcrNotes" json:"BuildStatusXdcrNotes"`
	BuildStatusAssembly      string    `bson:"BuildStatusAssembly" json:"BuildStatusAssembly"`
	BuildStatusAssemblyNotes string    `bson:"BuildStatusAssemblyNotes" json:"BuildStatusAssemblyNotes"`
	BuildStatusTesting       string    `bson:"BuildStatusTesting" json:"BuildStatusTesting"`
	BuildStatusTestingNotes  string    `bson:"BuildStatusTestingNotes" json:"BuildStatusTestingNotes"`
	DateShiped               time.Time `bson:"DateShiped" json:"DateShiped"`
	SerialNumber             string    `bson:"SerialNumber" json:"SerialNumber"`
	DeliveryDate             time.Time `bson:"DeliveryDate" json:"DeliveryDate"`
	Region                   string    `bson:"Region" json:"Region"`
	BeamAngle                string    `bson:"BeamAngle" json:"BeamAngle"`
	ElectronicsType          string    `bson:"ElectronicsType" json:"ElectronicsType"`
	BatteryType              string    `bson:"BatteryType" json:"BatteryType"`
	NumBattery               int       `bson:"NumBattery" json:"NumBattery"`
	ExternalBatteryPack      string    `bson:"ExternalBatteryPack" json:"ExternalBatteryPack"`
	CableLength              string    `bson:"CableLength" json:"CableLength"`
	InternalMemory           string    `bson:"InternalMemory" json:"InternalMemory"`
	InternalCompass          string    `bson:"InternalCompass" json:"InternalCompass"`
	DeckBox                  string    `bson:"DeckBox" json:"DeckBox"`
	PulseVersion             string    `bson:"PulseVersion" json:"PulseVersion"`
	UserVersion              string    `bson:"UserVersion" json:"UserVersion"`
	ThirdParty               string    `bson:"ThirdParty" json:"ThirdParty"`
	Notes                    string    `bson:"Notes" json:"Notes"`
}

// SalesOrderUpdate will contain the sales order data.
type SalesOrderUpdate struct {
	SO    SalesOrder
	Token string
}

// Add the SalesOrder.
func salesOrderAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		soData := &SalesOrderUpdate{}

		displaySalesOrderTemplate(w, soData)
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

		// Add the SalesOrders to the DB
		err = Vault.Mongo.C("SalesOrder").Insert(so)
		CheckError(err)

		// Go to the list of SalesOrders
		http.Redirect(w, r, "/so", http.StatusFound)
	}
}

// Display the template
func displaySalesOrderTemplate(w http.ResponseWriter, soData *SalesOrderUpdate) {
	// Generate token
	soData.Token = genToken()

	t, _ := template.ParseFiles("header.html", "salesorder_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", soData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, soData)
}
