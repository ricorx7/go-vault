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

// AdcpDataUpdate holds the ADCP data.
type AdcpDataUpdate struct {
	Adcp   Adcp
	Token  string
	Errors struct {
		SerialNumber string
	}
}

// Update the ADCP data
func adcpUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		val := bone.GetValue(r, "id")
		adcp := getAdcp(val)

		// Generate a token
		token := genToken()

		//var adcpData = {"SerialNumber": "1000056", "Desc": "1200 kHz", "token":token}
		//adcp := &Adcp{SerialNumber: "100001212", Frequency: "1200 kHz"}
		adcpData := &AdcpDataUpdate{Adcp: *adcp, Token: token}

		t, _ := template.ParseFiles("header.html", "adcp_update.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpData)
	} else {
		// Parse the form
		formData, err := forms.Parse(r)
		CheckError(err)

		fmt.Println("Update the ADCP")

		// Check token
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
			fmt.Println("Good token")
		} else {
			// give error if no token
			fmt.Println("Bad token")
		}

		// Use data to create a user object
		adcp := getAdcp(formData.Get("SerialNumber"))
		adcp.Customer = formData.Get("Customer")
		adcp.OrderNumber = formData.Get("OrderNumber")

		fmt.Println("Serial: ", adcp.SerialNumber)
		fmt.Println("Customer: ", adcp.Customer)
		fmt.Println("OrderNumber: ", adcp.OrderNumber)

		// Save the data to the DB
		updateAdcp2(adcp)

		// Go to the list of ADCP
		http.Redirect(w, r, "/adcp", http.StatusFound)
	}
}

// Find the ADCP from the database
func getAdcp(serialNum string) *Adcp {
	fmt.Println("Get ADCP - SerialNum: ", serialNum)

	var data Adcp
	err := Vault.Mongo.C("adcps").Find(bson.M{"SerialNumber": serialNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find document %v\n", err)
	}
	fmt.Println("SerialNum: ", data.SerialNumber)
	fmt.Println("Customer: ", data.Customer)
	fmt.Println("ID:", data.ID)

	return &data
}

// Update the ADCP data
func updateAdcp(adcp *Adcp) {
	fmt.Println("ID:", adcp.ID)

	// Update
	qry := bson.M{"SerialNumber": adcp.SerialNumber}
	change := bson.M{"$set": bson.M{"Customer": adcp.Customer, "OrderNumber": adcp.OrderNumber, "modified": time.Now()}}
	err := Vault.Mongo.C("adcps").Update(qry, change)
	CheckError(err)
}

// func updateAdcp1(adcp *Adcp) {
// 	change := mgo.Change{
// 		Update:    bson.M{"$inc": bson.M{"Customer": adcp.Customer, "OrderNumber": adcp.OrderNumber, "modified": time.Now()}},
// 		ReturnNew: true,
// 	}
// 	info, err := Vault.Mongo.C("adcps").Find(bson.M{"_id": adcp._id}).Apply(change, &doc)
// 	fmt.Println(doc.N)
// }

func updateAdcp2(adcp *Adcp) {
	fmt.Println("UpdateAdcp2 - ID", adcp.ID)

	//err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp._id}, bson.M{"$inc": bson.M{"Customer": adcp.Customer}})
	err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp.ID}, bson.M{"$set": bson.M{"Customer": adcp.Customer, "OrderNumber": adcp.OrderNumber}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
	}
}
