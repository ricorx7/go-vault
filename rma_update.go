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

		fmt.Printf("FormData %v\n", formData)

		// Use data to create a user object
		rma := getRma(formData.Get("OriginalRmaNumber"))
		fmt.Printf("Set RMA %s Orignal RMA Num %s\n", formData.Get("RmaNumber"), formData.Get("OriginalRmaNumber"))
		rma.RmaNumber = formData.Get("RmaNumber")
		rma.OrigSalesOrder = formData.Get("OrigSalesOrder")
		rma.Company = formData.Get("Company")
		rma.ContactName = formData.Get("ContactName")
		rma.ContactEmail = formData.Get("ContactEmail")
		rma.ContactPhone = formData.Get("ContactPhone")
		rma.ContactAddress = formData.Get("ContactAddress")
		rma.ContactAddress2 = formData.Get("ContactAddress2")
		rma.ContactAddressCityStateZip = formData.Get("ContactAddressCityStateZip")
		rma.ContactAddressCountry = formData.Get("ContactAddressCountry")
		rma.ProductDesc = formData.Get("ProductDesc")
		rma.ReasonReturn = formData.Get("ReasonReturn")
		rma.ReturnCompany = formData.Get("ReturnCompany")
		rma.ReturnContact = formData.Get("ReturnContact")
		rma.ReturnAddress = formData.Get("ReturnAddress")
		rma.ReturnAddressCont = formData.Get("ReturnAddressCont")
		rma.ReturnAddressCityStateZip = formData.Get("ReturnAddressCityStateZip")
		rma.ReturnAddressCountry = formData.Get("ReturnAddressCountry")
		rma.ReturnPhone = formData.Get("ReturnPhone")
		rma.ReturnEmail = formData.Get("ReturnEmail")
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

		// Accumulate the RepairProducts
		for i := range rma.RepairProducts {
			rma.RepairProducts[i].PartNumber = r.Form["ProductPartNumber"][i]
			rma.RepairProducts[i].SerialNumber = r.Form["ProductSerialNumber"][i]

			qty, err := strconv.Atoi(r.Form["ProductQty"][i])
			if err == nil {
				rma.RepairProducts[i].Qty = qty
			}
		}

		// Add the new product to the RMA
		if formData.Get("SubmitButton") == "ADD" {
			// fmt.Printf("Add product to RMA: %s\n", rma.RmaNumber)
			//
			// // Add the product to the list
			// rmaProduct := &RmaProduct{}
			// rmaProduct.PartNumber = formData.Get("AddProductPartNumber")
			// rmaProduct.Qty = formData.GetInt("AddProductQty")
			// rmaProduct.SerialNumber = formData.Get("AddProductSerialNumber")
			// product := getProductPartNumber(rmaProduct.PartNumber)
			// if product != nil {
			// 	rma.Products = append(rma.Products, *rmaProduct)
			// }
			// // Update the RMA in DB
			// updateRma(rma)
			addRmaProduct(formData.Get("AddProductPartNumber"),
				formData.GetInt("AddProductQty"),
				formData.Get("AddProductSerialNumber"),
				rma)

			// Go back to the update page
			http.Redirect(w, r, "/rma/update/"+rma.RmaNumber, http.StatusFound)
		} else if formData.Get("SubmitButton") == "ADD REPAIR" {
			// fmt.Printf("Add Repair product to RMA: %s\n", rma.RmaNumber)
			//
			// // Add the product to the list
			// rmaProduct := &RmaProduct{}
			// rmaProduct.PartNumber = formData.Get("AddRepairProductPartNumber")
			// rmaProduct.Qty = formData.GetInt("AddRepairProductQty")
			// rmaProduct.SerialNumber = formData.Get("AddRepairProductSerialNumber")
			// product := getProductPartNumber(rmaProduct.PartNumber)
			// if product != nil {
			// 	rma.RepairProducts = append(rma.RepairProducts, *rmaProduct)
			// }
			// // Update the RMA in DB
			// updateRma(rma)

			addRmaRepairProduct(formData.Get("AddRepairProductPartNumber"),
				formData.GetInt("AddRepairProductQty"),
				formData.Get("AddRepairProductSerialNumber"),
				rma)

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

// Add a product to the RMA Repair list.
func addRmaRepairProduct(partNum string, qty int, serialNum string, rma *RMA) {
	fmt.Printf("Add Repair product to RMA: %s\n", rma.RmaNumber)

	// Add the product to the list
	rmaProduct := &RmaProduct{}
	rmaProduct.PartNumber = partNum
	rmaProduct.Qty = qty
	rmaProduct.SerialNumber = serialNum
	product := getProductPartNumber(rmaProduct.PartNumber)
	if product != nil {
		rma.RepairProducts = append(rma.RepairProducts, *rmaProduct)
	}

	fmt.Printf("RMA %v\n", rma.RepairProducts)

	// Update the RMA in DB
	updateRma(rma)
}

// Add a product to the received parts list.
func addRmaProduct(partNum string, qty int, serialNum string, rma *RMA) {
	fmt.Printf("Add product to RMA: %s\n", rma.RmaNumber)

	// Add the product to the list
	rmaProduct := &RmaProduct{}
	rmaProduct.PartNumber = partNum
	rmaProduct.Qty = qty
	rmaProduct.SerialNumber = serialNum
	product := getProductPartNumber(rmaProduct.PartNumber)
	if product != nil {
		rma.Products = append(rma.Products, *rmaProduct)
	}
	// Update the RMA in DB
	updateRma(rma)
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
		fmt.Printf("Can't find RMA %v\n", err)
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
		"RmaNumber":                  rma.RmaNumber,
		"RmaDate":                    rma.RmaDate,
		"OrigSalesOrder":             rma.OrigSalesOrder,
		"Company":                    rma.Company,
		"ContactName":                rma.ContactName,
		"ContactEmail":               rma.ContactEmail,
		"ContactPhone":               rma.ContactPhone,
		"ContactAddress":             rma.ContactAddress,
		"ContactAddress2":            rma.ContactAddress2,
		"ContactAddressCityStateZip": rma.ContactAddressCityStateZip,
		"ContactAddressCountry":      rma.ContactAddressCountry,
		"ProductDesc":                rma.ProductDesc,
		"ReasonReturn":               rma.ReasonReturn,
		"ReturnCompany":              rma.ReturnCompany,
		"ReturnContact":              rma.ReturnContact,
		"ReturnAddress":              rma.ReturnAddress,
		"ReturnAddressCont":          rma.ReturnAddressCont,
		"ReturnAddressCityStateZip":  rma.ReturnAddressCityStateZip,
		"ReturnAddressCountry":       rma.ReturnAddressCountry,
		"ReturnPhone":                rma.ReturnPhone,
		"ReturnEmail":                rma.ReturnEmail,
		"ReceiveInfo":                rma.ReceiveInfo,
		"ReceiveUser":                rma.ReceiveUser,
		"ReceiveDate":                rma.ReceiveDate,
		"InspectionInfo":             rma.InspectionInfo,
		"InspectionUser":             rma.InspectionUser,
		"InspectionDate":             rma.InspectionDate,
		"RepairInfo":                 rma.RepairInfo,
		"RepairUser":                 rma.RepairUser,
		"RepairDate":                 rma.RepairDate,
		"RepairEstHours":             rma.RepairEstHours,
		"RepairProducts":             rma.RepairProducts,
		"Billable":                   rma.Billable,
		"QuoteNum":                   rma.QuoteNum,
		"OriginalRmaNum":             rma.OriginalRmaNum,
		"SerialNumber":               rma.SerialNumber,
		"Products":                   rma.Products,
		"Notes":                      rma.Notes,
		"ShipDate":                   rma.ShipDate,
		"Modified":                   rma.Modified,
		"Status":                     rma.Status}})
	if err != nil {
		fmt.Printf("Can't update RMA %v\n", err)
	}
}
