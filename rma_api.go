package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"

	"gopkg.in/mgo.v2/bson"
)

// RmaProduct will store a product for RMA.
type RmaProduct struct {
	PartNumber   string `bson:"PartNumber" json:"PartNumber"`
	Qty          int    `bson:"Qty" json:"Qty"`
	SerialNumber string `bson:"SerialNumber" json:"SerialNumber"`
}

// RMA will keep track of the sales order information.
// RMA Type is the first 3 digits
// RMA Number is the unique number for all RMA.
// RMA Types:
// 29 = Warrenty
// 28 = Billable
// 259 = Demo Repair
type RMA struct {
	ID                         bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	OrigSalesOrder             string          `bson:"OrigSalesOrder,omitempty" json:"OrigSalesOrder"`
	RmaDate                    string          `bson:"RmaDate" json:"RmaDate"`
	RmaType                    string          `bson:"RmaType" json:"RmaType"`
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

// RmaData holds the Water Test data.
type RmaData struct {
	RMA    []RMA  // RMAs
	Filter string // Filter
}

// Get the RMA data from the vault.
func vaultAPIRmaGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	rmaData := &RmaData{}

	// Get the limit
	limit := 0
	if len(r.URL.Query().Get("limit")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err == nil {
			limit = val
		}
	}

	// Get the offset
	offset := 0
	if len(r.URL.Query().Get("offset")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err == nil {
			offset = val
		}
	}

	// Get the filter
	// Get the data from the database
	filter := ""
	if len(r.URL.Query().Get("filter")) > 0 {
		filter = r.URL.Query().Get("filter")
		err := Vault.Mongo.C("RMAs").Find(bson.M{"SerialNumber": &bson.RegEx{Pattern: filter}}).Sort("-Created").All(&rmaData.RMA)
		CheckError(err)
	} else {
		err := Vault.Mongo.C("RMAs").Find(bson.M{}).Skip(offset).Limit(limit).Sort("-Created").All(&rmaData.RMA)
		CheckError(err)
	}

	// Get data form DB
	fmt.Println("Number of RMAs: ", len(rmaData.RMA))
	fmt.Printf("limit: %s\n", r.URL.Query().Get("limit"))
	fmt.Printf("offset: %s\n", r.URL.Query().Get("offset"))
	fmt.Printf("filter: %s\n", r.URL.Query().Get("filter"))

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(rmaData); err != nil {
		CheckError(err)
		panic(err)
	}
}

// Edit the RMA data from the vault.
func vaultAPIRmaEditHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			// ID is the ADCP serial number

			rma := getRmaResultsID(id) // Get the data from the database

			fmt.Printf("Edit RMA: %v\n", rma)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			// Pass back as a JSON
			if err := json.NewEncoder(w).Encode(rma); err != nil {
				panic(err)
			}
		}
	case "POST":
		{
			// ID is the database ID

			// Verify the data exist
			if r.Body == nil {
				http.Error(w, "Send a request body", 400)
				fmt.Printf("Edit waterest 1: \n")
				return
			}

			defer r.Body.Close()

			// Read in the data
			fmt.Println("response Headers:", r.Header)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("response Body:", string(body))

			// Convert to JSON
			var rma RMA
			err := json.Unmarshal(body, &rma)
			if err != nil {
				fmt.Println("Error with unmarsharl: ", err)
			}

			fmt.Printf("POST RMA: %v\n", rma)

			// Store the new data to the database
			updateRma(&rma)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	default:
		{

		}

	}
}

// Find the RMA from the database based off the ID
func getRmaResultsID(id string) *RMA {
	var data RMA

	err := Vault.Mongo.C("RMAs").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)
	if err != nil {
		fmt.Printf("Can't find from ID RMA data %v\n", err)
	}
	return &data
}
