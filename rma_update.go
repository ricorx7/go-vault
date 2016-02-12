package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

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
		rmaData.ProductList = *getProductList()
		rmaData.StatusList = getStatusList(rma.Status)
		rmaData.BillableList = getBillableList(rma.Billable)

		fmt.Printf("StatusList %v\n", rmaData.StatusList)

		displayRmaUpdateTemplate(w, rmaData)
	} else {
		fmt.Printf("Request: %v", r)

		// Parse the form
		formData, err := forms.Parse(r)
		CheckError(err)

		fmt.Printf("Button clicked: %s\n", formData.Get("SubmitButton"))
		fmt.Printf("Selected Product: %s\n", bson.ObjectId(formData.Get("AddProduct")))
		//fmt.Printf("Selected Product: %s\n", formData.Get("AddProduct"))

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
		val.TypeInt("AddProductQty")

		fmt.Printf("FormData %v\n", formData)

		// Use data to create a user object
		rma := getRma(formData.Get("RmaNumber"))
		rma.OrigSalesOrder = formData.Get("OrigSalesOrder")
		rma.Company = formData.Get("Company")
		rma.ContactName = formData.Get("ContactName")
		rma.ContactEmail = formData.Get("ContactEmail")
		rma.ContactPhone = formData.Get("ContactPhone")
		rma.ContactAddress = formData.Get("ContactAddress")
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
		rma.Modified = time.Now().Local()
		rma.InspectionDate = formData.Get("InspectionDate")
		rma.ReceiveDate = formData.Get("ReceiveDate")
		rma.RepairDate = formData.Get("RepairDate")
		rma.RmaDate = formData.Get("RmaDate")
		rma.ShipDate = formData.Get("ShipDate")
		rma.Status = formData.Get("Status")

		// Accumulate the Products
		for i := range rma.Products {
			rma.Products[i].PartNumber = r.Form["ProductPartNumber"][i]
			rma.Products[i].SerialNumber = r.Form["ProductSerialNumber"][i]

			qty, err := strconv.Atoi(r.Form["ProductQty"][i])
			if err == nil {
				rma.Products[i].Qty = qty
			}
		}

		// Add the new product to the RMA
		if formData.Get("SubmitButton") == "ADD" {
			if !val.HasErrors() {
				fmt.Printf("Add product to RMA: %s\n", rma.RmaNumber)

				// Add the product to the list
				rmaProduct := &RmaProduct{}
				rmaProduct.PartNumber = formData.Get("AddProductPartNumber")
				rmaProduct.Qty = formData.GetInt("AddProductQty")
				rmaProduct.SerialNumber = formData.Get("AddProductSerialNumber")
				product := getProductPartNumber(rmaProduct.PartNumber)
				if product != nil {
					rma.Products = append(rma.Products, *rmaProduct)
				}
				// Update the RMA in DB
				updateRma(rma)
			} else {
				fmt.Println("Error with values entered")
			}

			// Go back to the update page
			http.Redirect(w, r, "/rma/update/"+rma.RmaNumber, http.StatusFound)
		} else {
			fmt.Printf("RMA Update: %s\n", rma.RmaNumber)

			// Update the RMA in DB
			updateRma(rma)

			// Go to the list of RMA
			http.Redirect(w, r, "/rma", http.StatusFound)
		}
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

// Update the RMA data.
func updateRma(rma *RMA) {
	fmt.Println("updateRMA - ID", rma.ID)
	fmt.Printf("Products: %v\n", rma.Products)

	//err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp._id}, bson.M{"$inc": bson.M{"Customer": adcp.Customer}})
	err := Vault.Mongo.C("RMAs").Update(bson.M{"_id": rma.ID}, bson.M{"$set": bson.M{
		"RmaNumber":      rma.RmaNumber,
		"RmaDate":        rma.RmaDate,
		"OrigSalesOrder": rma.OrigSalesOrder,
		"Company":        rma.Company,
		"ContactName":    rma.ContactName,
		"ContactEmail":   rma.ContactEmail,
		"ContactPhone":   rma.ContactPhone,
		"ContactAddress": rma.ContactAddress,
		"ProductDesc":    rma.ProductDesc,
		"ReasonReturn":   rma.ReasonReturn,
		"ReceiveInfo":    rma.ReceiveInfo,
		"ReceiveUser":    rma.ReceiveUser,
		"ReceiveDate":    rma.ReceiveDate,
		"InspectionInfo": rma.InspectionInfo,
		"InspectionUser": rma.InspectionUser,
		"InspectionDate": rma.InspectionDate,
		"RepairInfo":     rma.RepairInfo,
		"RepairUser":     rma.RepairUser,
		"RepairDate":     rma.RepairDate,
		"RepairEstHours": rma.RepairEstHours,
		"Billable":       rma.Billable,
		"QuoteNum":       rma.QuoteNum,
		"OriginalRmaNum": rma.OriginalRmaNum,
		"SerialNumber":   rma.SerialNumber,
		"Products":       rma.Products,
		"Notes":          rma.Notes,
		"ShipDate":       rma.ShipDate,
		"Modified":       rma.Modified,
		"Status":         rma.Status}})
	if err != nil {
		fmt.Printf("Can't update RMA %v\n", err)
	}
}
