package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// RMA will keep track of the sales order information.
type RMA struct {
	ID               bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	SalesOrderID     bson.ObjectId   `bson:"SalesOrderID,omitempty" json:"SalesOrderID"`
	RmaDate          time.Time       `bson:"RmaDate" json:"RmaDate"`
	RmaNumber        string          `bson:"RmaNumber" json:"RmaNumber"`
	Company          string          `bson:"Company" json:"Company"`
	ContactName      string          `bson:"ContactName" json:"ContactName"`
	ContactEmail     string          `bson:"ContactEmail" json:"ContactEmail"`
	ContactPhone     string          `bson:"ContactPhone" json:"ContactPhone"`
	ProductDesc      string          `bson:"ProductDesc" json:"ProductDesc"`
	ProductID        bson.ObjectId   `bson:"ProductID,omitempty" json:"ProductID"`
	SerialNumber     string          `bson:"SerialNumber" json:"SerialNumber"`
	ReasonReturn     string          `bson:"ReasonReturn" json:"ReasonReturn"`
	ReceiveDate      time.Time       `bson:"ReceiveDate" json:"ReceiveDate"`
	ReceiveInfo      string          `bson:"ReceiveInfo" json:"ReceiveInfo"`
	ReceiveUser      string          `bson:"ReceiveUser" json:"ReceiveUser"`
	InspectionDate   time.Time       `bson:"InspectionDate" json:"InspectionDate"`
	InspectionInfo   string          `bson:"InspectionInfo" json:"InspectionInfo"`
	InspectionUser   string          `bson:"InspectionUser" json:"InspectionUser"`
	RepairDate       time.Time       `bson:"RepairDate" json:"RepairDate"`
	RepairInfo       string          `bson:"RepairInfo" json:"RepairInfo"`
	RepairUser       string          `bson:"RepairUser" json:"RepairUser"`
	RepairEstHours   int             `bson:"RepairEstHours" json:"RepairEstHours"`
	RepairMaterial   []string        `bson:"RepairMaterial" json:"RepairMaterial"`
	RepairMaterialID []bson.ObjectId `bson:"RepairMaterialID,omitempty" json:"RepairMaterialID"`
	Billable         string          `bson:"Billable" json:"Billable"`
	QuoteNum         string          `bson:"QuoteNum" json:"QuoteNum"`
	OriginalRmaNum   string          `bson:"OriginalRmaNum" json:"OriginalRmaNum"`
	Notes            string          `bson:"Notes" json:"Notes"`
	Created          time.Time       `bson:"Created" json:"Created"`
}

// RmaUpdate will contain the RMA data.
type RmaUpdate struct {
	RMA   RMA
	Token string
}

// Add the RMA.
func rmaAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Display blank form
		rmaData := &RmaUpdate{}
		displayRmaTemplate(w, rmaData)
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
		val.Require("Company")
		val.Require("ContactName")
		val.Require("ContactEmail")
		val.Require("ContactPhone")
		val.Require("SerialNumber")
		val.Require("ReasonReturn")

		// Use data to create a user object
		rma := &RMA{
			RmaNumber:      formData.Get("RmaNumber"),
			Company:        formData.Get("Company"),
			ContactName:    formData.Get("ContactName"),
			ContactEmail:   formData.Get("ContactEmail"),
			ContactPhone:   formData.Get("ContactPhone"),
			ProductDesc:    formData.Get("ProductDesc"),
			ReasonReturn:   formData.Get("ReasonReturn"),
			ReceiveInfo:    formData.Get("ReceiveInfo"),
			ReceiveUser:    formData.Get("ReceiveUser"),
			InspectionInfo: formData.Get("InspectionInfo"),
			InspectionUser: formData.Get("InspectionUser"),
			RepairInfo:     formData.Get("RepairInfo"),
			RepairUser:     formData.Get("RepairUser"),
			RepairEstHours: formData.GetInt("RepairEstHours"),
			Billable:       formData.Get("Billable"),
			QuoteNum:       formData.Get("QuoteNum"),
			OriginalRmaNum: formData.Get("OriginalRmaNum"),
			SerialNumber:   formData.Get("SerialNumber"),
			Notes:          formData.Get("Notes"),
			RmaDate:        time.Now().Local(),
			Created:        time.Now().Local(),
			//InspectionDate: formData.Get("InspectionDate"),
			//ReceiveDate:             formData.get("ReceiveDate"),
			//RepairDate:           formData.Get("RepairDate"),
			//DueDate:                  formData.Get("DueDate"),
			//EstShipDate:              formData.Get("EstShipDate"),

		}

		// Add the RMAs to the DB
		err = Vault.Mongo.C("RMAs").Insert(rma)
		CheckError(err)

		// Go to the list of SalesOrders
		http.Redirect(w, r, "/rma", http.StatusFound)
	}
}

// Display the template
func displayRmaTemplate(w http.ResponseWriter, rmaData *RmaUpdate) {
	// Generate token
	rmaData.Token = genToken()

	t, _ := template.ParseFiles("header.html", "rma_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", rmaData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, rmaData)
}
