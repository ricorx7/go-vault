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
	FirmwareVersion      string        `bson:"FirmwareVersion" json:"FirmwareVersion"`
	FinalCheckLake       CheckListItem `bson:"FinalCheckLake" json:"FinalCheckLake"`
	FinalCheckShipping   CheckListItem `bson:"FinalCheckShipping" json:"FinalCheckShipping"`
	BeamWires            CheckListItem `bson:"BeamWires" json:"BeamWires"`
	ThermistorStranded   CheckListItem `bson:"ThermistorStranded" json:"ThermistorStranded"`
	Temperature          CheckListItem `bson:"Temperature" json:"Temperature"`
	Pressure             CheckListItem `bson:"Pressure" json:"Pressure"`
	PressureSensorSize   string        `bson:"PressureSensorSize" json:"PressureSensorSize"`
	VaccumTest           CheckListItem `bson:"VaccumTest" json:"VaccumTest"`
	VibeTest             CheckListItem `bson:"VibeTest" json:"VibeTest"`
	LakeTest             CheckListItem `bson:"LakeTest" json:"LakeTest"`
	TankTest             CheckListItem `bson:"TankTest" json:"TankTest"`
	BurnInTestBoardStack CheckListItem `bson:"BurnInTestBoardStack" json:"BurnInTestBoardStack"`
	BurnInTestSystem     CheckListItem `bson:"BurnInTestSystem" json:"BurnInTestSystem"`
	AccuracyShort        CheckListItem `bson:"AccuracyShort" json:"AccuracyShort"`
	AccuracyJumper       CheckListItem `bson:"AccuracyJumper" json:"AccuracyJumper"`
	CompassCal           CheckListItem `bson:"CompassCal" json:"CompassCal"`
	BeamOrientation      CheckListItem `bson:"BeamOrientation" json:"BeamOrientation"`
	PressureTestSystem   CheckListItem `bson:"PressureTestSystem" json:"PressureTestSystem"`
	Battery1             CheckListItem `bson:"Battery1" json:"Battery1"`
	Battery1Voltage      string        `bson:"Battery1Voltage" json:"Battery1Voltage"`
	Battery1LoadVoltage  string        `bson:"Battery1LoadVoltage" json:"Battery1LoadVoltage"`
	Battery1SerialNumber string        `bson:"Battery1SerialNumber" json:"Battery1SerialNumber"`
	Battery1Rev          string        `bson:"Battery1Rev" json:"Battery1Rev"`
	Battery1PartNum      string        `bson:"Battery1PartNum" json:"Battery1PartNum"`
	Battery1LotNum       string        `bson:"Battery1LotNum" json:"Battery1LotNum"`
	Battery2             CheckListItem `bson:"Battery2" json:"Battery2"`
	Battery2Voltage      string        `bson:"Battery2Voltage" json:"Battery2Voltage"`
	Battery2LoadVoltage  string        `bson:"Battery2LoadVoltage" json:"Battery2LoadVoltage"`
	Battery2SerialNumber string        `bson:"Battery2SerialNumber" json:"Battery2SerialNumber"`
	Battery2Rev          string        `bson:"Battery2Rev" json:"Battery2Rev"`
	Battery2PartNum      string        `bson:"Battery2PartNum" json:"Battery2PartNum"`
	Battery2LotNum       string        `bson:"Battery2LotNum" json:"Battery2LotNum"`
	Modified             time.Time     `bson:"Modified" json:"Modified"`
	Created              time.Time     `bson:"Created" json:"Created"`
}

// AdcpChecklistUpdate will contain the RMA data.
type AdcpChecklistUpdate struct {
	ADCP                           AdcpChecklist
	OringStatusList                []OptionItem
	UrethaneStatusList             []OptionItem
	StandoffsStatusList            []OptionItem
	ScrewsStatusList               []OptionItem
	NotchStatusList                []OptionItem
	FirmwareStatusList             []OptionItem
	FinalCheckLakeStatusList       []OptionItem
	FinalCheckShippingStatusList   []OptionItem
	BeamWiresStatusList            []OptionItem
	ThermistorStrandedStatusList   []OptionItem
	TemperatureStatusList          []OptionItem
	PressureStatusList             []OptionItem
	VaccumTestStatusList           []OptionItem
	VibeTestStatusList             []OptionItem
	LakeTestStatusList             []OptionItem
	TankTestStatusList             []OptionItem
	BurnInTestBoardStackStatusList []OptionItem
	BurnInTestSystemStatusList     []OptionItem
	AccuracyShortStatusList        []OptionItem
	AccuracyJumperStatusList       []OptionItem
	CompassCalStatusList           []OptionItem
	BeamOrientationStatusList      []OptionItem
	PressureTestSystemStatusList   []OptionItem
	Battery1StatusList             []OptionItem
	Battery2StatusList             []OptionItem
	Token                          string
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
			Created:      time.Now().Local(),
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

		notch := &CheckListItem{
			Status: formData.Get("NotchStatus"),
			Date:   formData.Get("NotchDate"),
			User:   formData.Get("NotchUser"),
		}
		adcp.Notch = *notch

		firmware := &CheckListItem{
			Status: formData.Get("FirmwareStatus"),
			Date:   formData.Get("FirmwareDate"),
			User:   formData.Get("FirmwareUser"),
		}
		adcp.Firmware = *firmware
		adcp.FirmwareVersion = formData.Get("FirmwareVersion")

		finalCheckLake := &CheckListItem{
			Status: formData.Get("FinalCheckLakeStatus"),
			Date:   formData.Get("FinalCheckLakeDate"),
			User:   formData.Get("FinalCheckLakeUser"),
		}
		adcp.FinalCheckLake = *finalCheckLake

		finalCheckShipping := &CheckListItem{
			Status: formData.Get("FinalCheckShippingStatus"),
			Date:   formData.Get("FinalCheckShippingDate"),
			User:   formData.Get("FinalCheckShippingUser"),
		}
		adcp.FinalCheckShipping = *finalCheckShipping

		beamWires := &CheckListItem{
			Status: formData.Get("BeamWiresStatus"),
			Date:   formData.Get("BeamWiresDate"),
			User:   formData.Get("BeamWiresUser"),
		}
		adcp.BeamWires = *beamWires

		thermistorStranded := &CheckListItem{
			Status: formData.Get("ThermistorStrandedStatus"),
			Date:   formData.Get("ThermistorStrandedDate"),
			User:   formData.Get("ThermistorStrandedUser"),
		}
		adcp.ThermistorStranded = *thermistorStranded

		temperature := &CheckListItem{
			Status: formData.Get("TemperatureStatus"),
			Date:   formData.Get("TemperatureDate"),
			User:   formData.Get("TemperatureUser"),
		}
		adcp.Temperature = *temperature

		pressure := &CheckListItem{
			Status: formData.Get("PressureStatus"),
			Date:   formData.Get("PressureDate"),
			User:   formData.Get("PressureUser"),
		}
		adcp.Pressure = *pressure
		adcp.PressureSensorSize = formData.Get("PressureSensorSize")

		vaccumTest := &CheckListItem{
			Status: formData.Get("VaccumTestStatus"),
			Date:   formData.Get("VaccumTestDate"),
			User:   formData.Get("VaccumTestUser"),
		}
		adcp.VaccumTest = *vaccumTest

		vibeTest := &CheckListItem{
			Status: formData.Get("VibeTestStatus"),
			Date:   formData.Get("VibeTestDate"),
			User:   formData.Get("VibeTestUser"),
		}
		adcp.VibeTest = *vibeTest

		lakeTest := &CheckListItem{
			Status: formData.Get("LakeTestStatus"),
			Date:   formData.Get("LakeTestDate"),
			User:   formData.Get("LakeTestUser"),
		}
		adcp.LakeTest = *lakeTest

		tankTest := &CheckListItem{
			Status: formData.Get("TankTestStatus"),
			Date:   formData.Get("TankTestDate"),
			User:   formData.Get("TankTestUser"),
		}
		adcp.TankTest = *tankTest

		burnInTestBoardStack := &CheckListItem{
			Status: formData.Get("BurnInTestBoardStackStatus"),
			Date:   formData.Get("BurnInTestBoardStackDate"),
			User:   formData.Get("BurnInTestBoardStackUser"),
		}
		adcp.BurnInTestBoardStack = *burnInTestBoardStack

		burnInTestSystem := &CheckListItem{
			Status: formData.Get("BurnInTestSystemStatus"),
			Date:   formData.Get("BurnInTestSystemDate"),
			User:   formData.Get("BurnInTestSystemUser"),
		}
		adcp.BurnInTestSystem = *burnInTestSystem

		accuracyShort := &CheckListItem{
			Status: formData.Get("AccuracyShortStatus"),
			Date:   formData.Get("AccuracyShortDate"),
			User:   formData.Get("AccuracyShortUser"),
		}
		adcp.AccuracyShort = *accuracyShort

		accuracyJumper := &CheckListItem{
			Status: formData.Get("AccuracyJumperStatus"),
			Date:   formData.Get("AccuracyJumperDate"),
			User:   formData.Get("AccuracyJumperUser"),
		}
		adcp.AccuracyJumper = *accuracyJumper

		compassCal := &CheckListItem{
			Status: formData.Get("CompassCalStatus"),
			Date:   formData.Get("CompassCalDate"),
			User:   formData.Get("CompassCalUser"),
		}
		adcp.CompassCal = *compassCal

		beamOrientation := &CheckListItem{
			Status: formData.Get("BeamOrientationStatus"),
			Date:   formData.Get("BeamOrientationDate"),
			User:   formData.Get("BeamOrientationUser"),
		}
		adcp.BeamOrientation = *beamOrientation

		pressureTestSystem := &CheckListItem{
			Status: formData.Get("PressureTestSystemStatus"),
			Date:   formData.Get("PressureTestSystemDate"),
			User:   formData.Get("PressureTestSystemUser"),
		}
		adcp.PressureTestSystem = *pressureTestSystem

		battery1 := &CheckListItem{
			Status: formData.Get("Battery1Status"),
			Date:   formData.Get("Battery1Date"),
			User:   formData.Get("Battery1User"),
		}
		adcp.Battery1 = *battery1
		adcp.Battery1Voltage = formData.Get("Battery1Voltage")
		adcp.Battery1LoadVoltage = formData.Get("Battery1LoadVoltage")
		adcp.Battery1LotNum = formData.Get("Battery1LotNum")
		adcp.Battery1PartNum = formData.Get("Battery1PartNum")
		adcp.Battery1Rev = formData.Get("Battery1Rev")
		adcp.Battery1SerialNumber = formData.Get("Battery1SerialNumber")

		battery2 := &CheckListItem{
			Status: formData.Get("Battery2Status"),
			Date:   formData.Get("Battery2Date"),
			User:   formData.Get("Battery2User"),
		}
		adcp.Battery2 = *battery2
		adcp.Battery2Voltage = formData.Get("Battery2Voltage")
		adcp.Battery2LoadVoltage = formData.Get("Battery2LoadVoltage")
		adcp.Battery2LotNum = formData.Get("Battery2LotNum")
		adcp.Battery2PartNum = formData.Get("Battery2PartNum")
		adcp.Battery2Rev = formData.Get("Battery2Rev")
		adcp.Battery2SerialNumber = formData.Get("Battery2SerialNumber")

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
