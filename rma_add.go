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
	ID               bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	OrigSalesOrder   string          `bson:"OrigSalesOrder,omitempty" json:"OrigSalesOrder"`
	RmaDate          string          `bson:"RmaDate" json:"RmaDate"`
	RmaNumber        string          `bson:"RmaNumber" json:"RmaNumber"`
	Company          string          `bson:"Company" json:"Company"`
	ContactName      string          `bson:"ContactName" json:"ContactName"`
	ContactAddress   string          `bson:"ContactAddress" json:"ContactAddress"`
	ContactEmail     string          `bson:"ContactEmail" json:"ContactEmail"`
	ContactPhone     string          `bson:"ContactPhone" json:"ContactPhone"`
	ProductDesc      string          `bson:"ProductDesc" json:"ProductDesc"`
	ProductID        bson.ObjectId   `bson:"ProductID,omitempty" json:"ProductID"`
	Products         []RmaProduct    `bson:"Products" json:"Products"`
	SerialNumber     string          `bson:"SerialNumber" json:"SerialNumber"`
	ReasonReturn     string          `bson:"ReasonReturn" json:"ReasonReturn"`
	ReceiveDate      string          `bson:"ReceiveDate" json:"ReceiveDate"`
	ReceiveInfo      string          `bson:"ReceiveInfo" json:"ReceiveInfo"`
	ReceiveUser      string          `bson:"ReceiveUser" json:"ReceiveUser"`
	InspectionDate   string          `bson:"InspectionDate" json:"InspectionDate"`
	InspectionInfo   string          `bson:"InspectionInfo" json:"InspectionInfo"`
	InspectionUser   string          `bson:"InspectionUser" json:"InspectionUser"`
	RepairDate       string          `bson:"RepairDate" json:"RepairDate"`
	RepairInfo       string          `bson:"RepairInfo" json:"RepairInfo"`
	RepairUser       string          `bson:"RepairUser" json:"RepairUser"`
	RepairEstHours   int             `bson:"RepairEstHours" json:"RepairEstHours"`
	RepairMaterial   []string        `bson:"RepairMaterial" json:"RepairMaterial"`
	RepairMaterialID []bson.ObjectId `bson:"RepairMaterialID,omitempty" json:"RepairMaterialID"`
	Billable         string          `bson:"Billable" json:"Billable"`
	QuoteNum         string          `bson:"QuoteNum" json:"QuoteNum"`
	OriginalRmaNum   string          `bson:"OriginalRmaNum" json:"OriginalRmaNum"`
	Notes            string          `bson:"Notes" json:"Notes"`
	Status           string          `bson:"Status" json:"Status"`
	ShipDate         string          `bson:"ShipDate" json:"ShipDate"`
	Created          time.Time       `bson:"Created" json:"Created"`
	Modified         time.Time       `bson:"Modified" json:"Modified"`
}

// RmaUpdate will contain the RMA data.
type RmaUpdate struct {
<<<<<<< HEAD
	RMA          RMA
	AddProduct   bson.ObjectId
	ProductList  []Product
	StatusList   []OptionItem
	BillableList []OptionItem
	Token        string
=======
	RMA         RMA
	AddProduct  bson.ObjectId
	ProductList []Product
	StatusList  []OptionItem
	Token       string
>>>>>>> 748ed8acfcc3d7ae2c2901f6e9bd93d330491997
}

// Add the RMA.
func rmaAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Display blank form
		rmaData := &RmaUpdate{}
		rmaData.ProductList = *getProductList()
		rmaData.StatusList = getStatusList("Reported")
<<<<<<< HEAD
		rmaData.BillableList = getBillableList("Billable")
=======
>>>>>>> 748ed8acfcc3d7ae2c2901f6e9bd93d330491997

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
			RmaNumber:      formData.Get("RmaNumber"),
			OrigSalesOrder: formData.Get("OrigSalesOrder"),
			Company:        formData.Get("Company"),
			ContactName:    formData.Get("ContactName"),
			ContactEmail:   formData.Get("ContactEmail"),
			ContactPhone:   formData.Get("ContactPhone"),
			ContactAddress: formData.Get("ContactAddress"),
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
			RmaDate:        formData.Get("RmaDate"),
			Created:        time.Now().Local(),
			Modified:       time.Now().Local(),
			InspectionDate: formData.Get("InspectionDate"),
			ReceiveDate:    formData.Get("ReceiveDate"),
			RepairDate:     formData.Get("RepairDate"),
			ShipDate:       formData.Get("ShipDate"),
			Status:         formData.Get("Status"),
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
			// Update the RMA in DB
			//updateRma(rma)
			// } else {
			// 	fmt.Println("Error with values entered")
			// }
			rmaData := &RmaUpdate{}
			rmaData.ProductList = *getProductList()
			rmaData.StatusList = getStatusList(rma.Status)
<<<<<<< HEAD
			rmaData.BillableList = getBillableList(rma.Billable)
=======
>>>>>>> 748ed8acfcc3d7ae2c2901f6e9bd93d330491997
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
<<<<<<< HEAD

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
=======
>>>>>>> 748ed8acfcc3d7ae2c2901f6e9bd93d330491997
