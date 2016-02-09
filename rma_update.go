package main

import (
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2/bson"

	"github.com/albrow/forms"
	"github.com/go-zoo/bone"
)

// Update the RMA.
func rmaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		rmaNum := bone.GetValue(r, "id")
		rma := getRma(rmaNum)
		rmaData := &RmaUpdate{}
		rmaData.RMA = *rma

		displayRmaUpdateTemplate(w, rmaData)
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
		val.Require("RmaNumber")

		// Use data to create a user object
		rma := getRma(formData.Get("RmaNumber"))
		rma.RmaNumber = formData.Get("RmaNumber")
		rma.Company = formData.Get("Company")
		rma.ContactName = formData.Get("ContactName")
		rma.ContactEmail = formData.Get("ContactEmail")
		rma.ContactPhone = formData.Get("ContactPhone")
		rma.ProductDesc = formData.Get("ProductDesc")
		rma.ReasonReturn = formData.Get("ReasonReturn")
		rma.ReceiveInfo = formData.Get("ReceiveInfo")
		rma.ReceiveUser = formData.Get("ReceiveUser")
		rma.InspectionInfo = formData.Get("InspectionInfo")
		rma.InspectionUser = formData.Get("InspectionUser")
		rma.RepairInfo = formData.Get("RepairInfo")
		rma.RepairUser = formData.Get("RepairUser")
		rma.RepairEstHours = formData.GetInt("RepairEstHours")
		rma.Billable = formData.Get("Billable")
		rma.QuoteNum = formData.Get("QuoteNum")
		rma.OriginalRmaNum = formData.Get("OriginalRmaNum")
		rma.SerialNumber = formData.Get("SerialNumber")
		rma.Notes = formData.Get("Notes")

		fmt.Printf("RMA Update: %s\n", rma.RmaNumber)

		// Update the RMA in DB
		updateRma(rma)

		// Go to the list of ADCP
		http.Redirect(w, r, "/rma", http.StatusFound)
	}
}

// Display the template
func displayRmaUpdateTemplate(w http.ResponseWriter, rmaData *RmaUpdate) {
	// Generate token
	rmaData.Token = genToken()

	// Redirect back to the page with error message
	t, _ := template.ParseFiles("header.html", "rma_update.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", rmaData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, rmaData)
}

// Find the RMA from the database
func getRma(rmaNum string) *RMA {
	fmt.Println("Get RMA - RMANumber: ", rmaNum)

	var data RMA
	err := Vault.Mongo.C("RMAs").Find(bson.M{"RmaNumber": rmaNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find SalesOrder %v\n", err)
	}

	fmt.Println("RMA: ", data.RmaNumber)
	fmt.Println("SerialNum: ", data.SerialNumber)
	fmt.Println("ReasonReturn: ", data.ReasonReturn)
	fmt.Println("ID:", data.ID)

	return &data
}

func updateRma(rma *RMA) {
	fmt.Println("updateRMA - ID", rma.ID)

	//err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp._id}, bson.M{"$inc": bson.M{"Customer": adcp.Customer}})
	err := Vault.Mongo.C("RMAs").Update(bson.M{"_id": rma.ID}, bson.M{"$set": bson.M{
		"RmaNumber":      rma.RmaNumber,
		"Company":        rma.Company,
		"ContactName":    rma.ContactName,
		"ContactEmail":   rma.ContactEmail,
		"ContactPhone":   rma.ContactPhone,
		"ProductDesc":    rma.ProductDesc,
		"ReasonReturn":   rma.ReasonReturn,
		"ReceiveInfo":    rma.ReceiveInfo,
		"ReceiveUser":    rma.ReceiveUser,
		"InspectionInfo": rma.InspectionInfo,
		"InspectionUser": rma.InspectionUser,
		"RepairInfo":     rma.RepairInfo,
		"RepairUser":     rma.RepairUser,
		"RepairEstHours": rma.RepairEstHours,
		"Billable":       rma.Billable,
		"QuoteNum":       rma.QuoteNum,
		"OriginalRmaNum": rma.OriginalRmaNum,
		"SerialNumber":   rma.SerialNumber,
		"Notes":          rma.Notes}})
	if err != nil {
		fmt.Printf("Can't update RMA %v\n", err)
	}
}
