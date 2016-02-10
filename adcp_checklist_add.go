package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// CheckListItem will store the checklist results.
type CheckListItem struct {
	Status string `bson:"Status" json:"Status"` // Not Test, N/A, Tested
	Date   string `bson:"Date" json:"Date"`
	User   string `bson:"User" json:"User"`
}

// AdcpChecklist will keep track of the product information.
type AdcpChecklist struct {
	ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	SerialNumber         string        `bson:"SerialNumber" json:"SerialNumber"`
	Oring                CheckListItem `bson:"Oring" json:"Oring"`
	Urethane             CheckListItem `bson:"Urethane" json:"Urethane"`
	Screws               CheckListItem `bson:"Screws" json:"Screws"`
	Standoffs            CheckListItem `bson:"Standoffs" json:"Standoffs"`
	Notch                CheckListItem `bson:"Notch" json:"Notch"`
	Firmware             CheckListItem `bson:"Firmware" json:"Firmware"`
	FinalCheckLake       CheckListItem `bson:"FinalCheckLake" json:"FinalCheckLake"`
	FinalCheckShipping   CheckListItem `bson:"FinalCheckShipping" json:"FinalCheckShipping"`
	BeamWires            CheckListItem `bson:"BeamWires" json:"BeamWires"`
	ThermistorStranded   CheckListItem `bson:"ThermistorStranded" json:"ThermistorStranded"`
	Temperature          CheckListItem `bson:"Temperature" json:"Temperature"`
	Pressure             CheckListItem `bson:"Pressure" json:"Pressure"`
	VaccumTest           CheckListItem `bson:"VaccumTest" json:"VaccumTest"`
	VibeTest             CheckListItem `bson:"VibeTest" json:"VibeTest"`
	BurnInTestBoardStack CheckListItem `bson:"BurnInTestBoardStack" json:"BurnInTestBoardStack"`
	BurnInTestSystem     CheckListItem `bson:"BurnInTestSystem" json:"BurnInTestSystem"`
	AccuracyShort        CheckListItem `bson:"AccurayShort" json:"AccuracyShort"`
	AccuracyJumper       CheckListItem `bson:"AccuracyJumper" json:"AccuracyJumper"`
	CompassCal           CheckListItem `bson:"CompassCal" json:"CompassCal"`
	BeamOrientation      CheckListItem `bson:"BeamOrientation" json:"BeamOrientation"`
	Modified             time.Time     `bson:"Modified" json:"Modified"`
}

// AdcpChecklistUpdate will contain the RMA data.
type AdcpChecklistUpdate struct {
	ADCP  AdcpChecklist
	Token string
}

// Add the Product.
func adcpChecklistAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Display blank form
		data := &AdcpChecklistUpdate{}
		displayAdcpChecklistAddTemplate(w, data)
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
		val.Require("SerialNumber")

		// Use data to create a user object
		adcp := &AdcpChecklist{
			SerialNumber: formData.Get("SerialNumber"),
			Modified:     time.Now().Local(),
		}
		oring := &CheckListItem{
			Status: formData.Get("OringStatus"),
			Date:   formData.Get("OringDate"),
			User:   formData.Get("OringUser"),
		}
		adcp.Oring = *oring

		urethane := &CheckListItem{
			Status: formData.Get("UrethaneStatus"),
			Date:   formData.Get("UrethaneDate"),
			User:   formData.Get("UrethaneUser"),
		}
		adcp.Urethane = *urethane

		standoff := &CheckListItem{
			Status: formData.Get("StandoffsStatus"),
			Date:   formData.Get("StandoffsDate"),
			User:   formData.Get("StandoffsUser"),
		}
		adcp.Standoffs = *standoff

		screws := &CheckListItem{
			Status: formData.Get("ScrewsStatus"),
			Date:   formData.Get("ScrewsDate"),
			User:   formData.Get("ScrewsUser"),
		}
		adcp.Screws = *screws

		checklist := AdcpChecklistUpdate{}
		checklist.ADCP = *adcp

		// Add the Products to the DB
		err = Vault.Mongo.C("AdcpChecklist").Insert(checklist)
		CheckError(err)

		// Go to the list of products
		http.Redirect(w, r, "/adcp", http.StatusFound)
	}
}

// Display the template
func displayAdcpChecklistAddTemplate(w http.ResponseWriter, data *AdcpChecklistUpdate) {
	// Generate token
	data.Token = genToken()

	t, _ := template.ParseFiles("header.html", "adcp_checklist_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", data)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, data)
}
