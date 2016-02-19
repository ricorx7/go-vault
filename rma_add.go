package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// RmaProduct will store a product for RMA.
type RmaProduct struct {
	PartNumber   string `bson:"PartNumber" json:"PartNumber"`
	Qty          int    `bson:"Qty" json:"Qty"`
	SerialNumber string `bson:"SerialNumber" json:"SerialNumber"`
}

// RMA will keep track of the sales order information.
type RMA struct {
	ID                         bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	OrigSalesOrder             string          `bson:"OrigSalesOrder,omitempty" json:"OrigSalesOrder"`
	RmaDate                    string          `bson:"RmaDate" json:"RmaDate"`
	RmaNumber                  string          `bson:"RmaNumber" json:"RmaNumber"`
	Company                    string          `bson:"Company" json:"Company"`
	ContactName                string          `bson:"ContactName" json:"ContactName"`
	ContactAddress             string          `bson:"ContactAddress" json:"ContactAddress"`
	ContactAddress2            string          `bson:"ContactAddress2" json:"ContactAddress2"`
	ContactAddressCityStateZip string          `bson:"ContactAddressCityStateZip" json:"ContactAddressCityStateZip"`
	ContactAddressCountry      string          `bson:"ContactAddressCountry" json:"ContactAddressCountry"`
	ContactEmail               string          `bson:"ContactEmail" json:"ContactEmail"`
	ContactPhone               string          `bson:"ContactPhone" json:"ContactPhone"`
	ProductDesc                string          `bson:"ProductDesc" json:"ProductDesc"`
	ProductID                  bson.ObjectId   `bson:"ProductID,omitempty" json:"ProductID"`
	Products                   []RmaProduct    `bson:"Products" json:"Products"`
	SerialNumber               string          `bson:"SerialNumber" json:"SerialNumber"`
	ReasonReturn               string          `bson:"ReasonReturn" json:"ReasonReturn"`
	ReturnCompany              string          `bson:"ReturnCompany" json:"ReturnCompany"`
	ReturnContact              string          `bson:"ReturnContact" json:"ReturnContact"`
	ReturnAddress              string          `bson:"ReturnAddress" json:"ReturnAddress"`
	ReturnAddressCont          string          `bson:"ReturnAddressCont" json:"ReturnAddressCont"`
	ReturnAddressCityStateZip  string          `bson:"ReturnAddressCityStateZip" json:"ReturnAddressCityStateZip"`
	ReturnAddressCountry       string          `bson:"ReturnAddressCountry" json:"ReturnAddressCountry"`
	ReturnPhone                string          `bson:"ReturnPhone" json:"ReturnPhone"`
	ReturnEmail                string          `bson:"ReturnEmail" json:"ReturnEmail"`
	ReceiveDate                string          `bson:"ReceiveDate" json:"ReceiveDate"`
	ReceiveInfo                string          `bson:"ReceiveInfo" json:"ReceiveInfo"`
	ReceiveUser                string          `bson:"ReceiveUser" json:"ReceiveUser"`
	InspectionDate             string          `bson:"InspectionDate" json:"InspectionDate"`
	InspectionInfo             string          `bson:"InspectionInfo" json:"InspectionInfo"`
	InspectionUser             string          `bson:"InspectionUser" json:"InspectionUser"`
	RepairDate                 string          `bson:"RepairDate" json:"RepairDate"`
	RepairInfo                 string          `bson:"RepairInfo" json:"RepairInfo"`
	RepairUser                 string          `bson:"RepairUser" json:"RepairUser"`
	RepairEstHours             int             `bson:"RepairEstHours" json:"RepairEstHours"`
	RepairMaterial             []string        `bson:"RepairMaterial" json:"RepairMaterial"`
	RepairMaterialID           []bson.ObjectId `bson:"RepairMaterialID,omitempty" json:"RepairMaterialID"`
	RepairProducts             []RmaProduct    `bson:"RepairProducts" json:"RepairProducts"`
	Billable                   string          `bson:"Billable" json:"Billable"`
	QuoteNum                   string          `bson:"QuoteNum" json:"QuoteNum"`
	OriginalRmaNum             string          `bson:"OriginalRmaNum" json:"OriginalRmaNum"`
	Notes                      string          `bson:"Notes" json:"Notes"`
	Status                     string          `bson:"Status" json:"Status"`
	ShipDate                   string          `bson:"ShipDate" json:"ShipDate"`
	Created                    time.Time       `bson:"Created" json:"Created"`
	Modified                   time.Time       `bson:"Modified" json:"Modified"`
}

// RmaUpdate will contain the RMA data.
type RmaUpdate struct {
	RMA          RMA
	AddProduct   bson.ObjectId
	ProductList  []Product
	StatusList   []OptionItem
	BillableList []OptionItem
	Token        string
}

// Add the RMA.
func rmaAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Display blank form
		rmaData := &RmaUpdate{}
		rmaData.ProductList = *getProductList()
		rmaData.StatusList = getStatusList("Reported")
		rmaData.BillableList = getBillableList("Billable")

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

		// // Validate data
		// val := formData.Validator()
		// val.Require("Company")
		// val.Require("ContactName")
		// val.Require("ContactEmail")
		// val.Require("ContactPhone")
		// val.Require("SerialNumber")
		// val.Require("ReasonReturn")

		// Use data to create a user object
		rma := &RMA{
			RmaNumber:                  formData.Get("RmaNumber"),
			OrigSalesOrder:             formData.Get("OrigSalesOrder"),
			Company:                    formData.Get("Company"),
			ContactName:                formData.Get("ContactName"),
			ContactEmail:               formData.Get("ContactEmail"),
			ContactPhone:               formData.Get("ContactPhone"),
			ContactAddress:             formData.Get("ContactAddress"),
			ContactAddress2:            formData.Get("ContactAddress2"),
			ContactAddressCityStateZip: formData.Get("ContactAddressCityStateZip"),
			ContactAddressCountry:      formData.Get("ContactAddressCountry"),
			ProductDesc:                formData.Get("ProductDesc"),
			ReasonReturn:               formData.Get("ReasonReturn"),
			ReturnCompany:              formData.Get("ReturnCompany"),
			ReturnContact:              formData.Get("ReturnContact"),
			ReturnAddress:              formData.Get("ReturnAddress"),
			ReturnAddressCont:          formData.Get("ReturnAddressCont"),
			ReturnAddressCityStateZip:  formData.Get("ReturnAddressCityStateZip"),
			ReturnAddressCountry:       formData.Get("ReturnAddressCountry"),
			ReturnPhone:                formData.Get("ReturnPhone"),
			ReturnEmail:                formData.Get("ReturnEmail"),
			ReceiveInfo:                formData.Get("ReceiveInfo"),
			ReceiveUser:                formData.Get("ReceiveUser"),
			InspectionInfo:             formData.Get("InspectionInfo"),
			InspectionUser:             formData.Get("InspectionUser"),
			RepairInfo:                 formData.Get("RepairInfo"),
			RepairUser:                 formData.Get("RepairUser"),
			RepairEstHours:             formData.GetInt("RepairEstHours"),
			Billable:                   formData.Get("Billable"),
			QuoteNum:                   formData.Get("QuoteNum"),
			OriginalRmaNum:             formData.Get("OriginalRmaNum"),
			SerialNumber:               formData.Get("SerialNumber"),
			Notes:                      formData.Get("Notes"),
			RmaDate:                    formData.Get("RmaDate"),
			Created:                    time.Now().Local(),
			Modified:                   time.Now().Local(),
			InspectionDate:             formData.Get("InspectionDate"),
			ReceiveDate:                formData.Get("ReceiveDate"),
			RepairDate:                 formData.Get("RepairDate"),
			ShipDate:                   formData.Get("ShipDate"),
			Status:                     formData.Get("Status"),
		}

		fmt.Printf("Submit Button: %s\n", formData.Get("SubmitButton"))

		// Add the new product to the RMA
		if formData.Get("SubmitButton") == "ADD" {
			//if !val.HasErrors() {
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

			rmaData := &RmaUpdate{}
			rmaData.ProductList = *getProductList()
			rmaData.StatusList = getStatusList(rma.Status)
			rmaData.BillableList = getBillableList(rma.Billable)
			rmaData.RMA = *rma

			displayRmaTemplate(w, rmaData)
		} else if formData.Get("SubmitButton") == "ADD REPAIR" {
			//if !val.HasErrors() {
			fmt.Printf("Add repair product to RMA: %s\n", rma.RmaNumber)

			// Add the product to the repair list
			rmaProduct := &RmaProduct{}
			rmaProduct.PartNumber = formData.Get("AddRepairProductPartNumber")
			rmaProduct.Qty = formData.GetInt("AddRepairProductQty")
			rmaProduct.SerialNumber = formData.Get("AddRepairProductSerialNumber")
			product := getProductPartNumber(rmaProduct.PartNumber)
			if product != nil {
				rma.RepairProducts = append(rma.Products, *rmaProduct)
			}

			rmaData := &RmaUpdate{}
			rmaData.ProductList = *getProductList()
			rmaData.StatusList = getStatusList(rma.Status)
			rmaData.BillableList = getBillableList(rma.Billable)
			rmaData.RMA = *rma

			displayRmaTemplate(w, rmaData)
		} else {
			fmt.Printf("RMA Add: %s\n", rma.RmaNumber)

			// Accumulate the Products
			for i := range rma.Products {
				rma.Products[i].PartNumber = r.Form["ProductPartNumber"][i]
				rma.Products[i].SerialNumber = r.Form["ProductSerialNumber"][i]

				qty, err := strconv.Atoi(r.Form["ProductQty"][i])
				if err == nil {
					rma.Products[i].Qty = qty
				}
			}

			// Accumulate the Repair Products
			for i := range rma.RepairProducts {
				rma.RepairProducts[i].PartNumber = r.Form["RepairProductPartNumber"][i]
				rma.RepairProducts[i].SerialNumber = r.Form["RepairProductSerialNumber"][i]

				qty, err := strconv.Atoi(r.Form["RepairProductQty"][i])
				if err == nil {
					rma.RepairProducts[i].Qty = qty
				}
			}

			// Add the RMAs to the DB
			err = Vault.Mongo.C("RMAs").Insert(rma)
			CheckError(err)

			// Go to the list of RMAs
			http.Redirect(w, r, "/rma", http.StatusFound)
		}
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

// Create a Status slice.  Then set the selected flag
// based off the status value given.
func getStatusList(status string) []OptionItem {
	options := []OptionItem{
		OptionItem{"Reported", "Reported", false},
		OptionItem{"Received", "Received", false},
		OptionItem{"Inspected", "Inspected", false},
		OptionItem{"Repaired", "Repaired", false},
		OptionItem{"Returned", "Returned", false},
		OptionItem{"Completed", "Completed", false},
	}

	// Set the selected value based off the status given
	for i := range options {
		if options[i].Value == status {
			options[i].Selected = true
		}
	}

	return options
}

// Create a Billable slice.  Then set the selected flag
// based off the billable value given.
func getBillableList(billable string) []OptionItem {
	options := []OptionItem{
		OptionItem{"Billable", "Billable", false},
		OptionItem{"Warranty", "Warranty", false},
		OptionItem{"N/A", "N/A", false},
	}

	// Set the selected value based off the string given
	for i := range options {
		if options[i].Value == billable {
			options[i].Selected = true
		}
	}

	return options
}
