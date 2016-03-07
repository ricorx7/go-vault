package main

// Update the ADCP data
import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/albrow/forms"
	"github.com/go-zoo/bone"
)

func adcpChecklistUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		val := bone.GetValue(r, "id")

		// Check if the checklist exist, if it does not, add it
		CheckAdcpChecklist(val)

		// Get the checklist from the DB
		checklist := getAdcpChecklist(val)

		// Generate a token
		token := genToken()

		adcpChecklist := &AdcpChecklistUpdate{ADCP: *checklist, Token: token}
		adcpChecklist.OringStatusList = getChecklistStatusList(checklist.Oring.Status)
		adcpChecklist.UrethaneStatusList = getChecklistStatusList(checklist.Urethane.Status)
		adcpChecklist.StandoffsStatusList = getChecklistStatusList(checklist.Standoffs.Status)
		adcpChecklist.ScrewsStatusList = getChecklistStatusList(checklist.Standoffs.Status)
		adcpChecklist.NotchStatusList = getChecklistStatusList(checklist.Notch.Status)
		adcpChecklist.FirmwareStatusList = getChecklistStatusList(checklist.Firmware.Status)
		adcpChecklist.FinalCheckLakeStatusList = getChecklistStatusList(checklist.FinalCheckLake.Status)
		adcpChecklist.FinalCheckShippingStatusList = getChecklistStatusList(checklist.FinalCheckShipping.Status)
		adcpChecklist.BeamWiresStatusList = getChecklistStatusList(checklist.BeamWires.Status)
		adcpChecklist.ThermistorStrandedStatusList = getChecklistStatusList(checklist.ThermistorStranded.Status)
		adcpChecklist.TemperatureStatusList = getChecklistStatusList(checklist.Temperature.Status)
		adcpChecklist.PressureStatusList = getChecklistStatusList(checklist.Pressure.Status)
		adcpChecklist.VaccumTestStatusList = getChecklistStatusList(checklist.VaccumTest.Status)
		adcpChecklist.VibeTestStatusList = getChecklistStatusList(checklist.VibeTest.Status)
		adcpChecklist.LakeTestStatusList = getChecklistStatusList(checklist.LakeTest.Status)
		adcpChecklist.ReviewLakeTestDataStatusList = getChecklistStatusList(checklist.ReviewLakeTestData.Status)
		adcpChecklist.TankTestStatusList = getChecklistStatusList(checklist.TankTest.Status)
		adcpChecklist.ReviewTankTestDataStatusList = getChecklistStatusList(checklist.ReviewTankTestData.Status)
		adcpChecklist.BurnInTestBoardStackStatusList = getChecklistStatusList(checklist.BurnInTestBoardStack.Status)
		adcpChecklist.BurnInTestSystemStatusList = getChecklistStatusList(checklist.BurnInTestSystem.Status)
		adcpChecklist.AccuracyJumperStatusList = getChecklistStatusList(checklist.AccuracyJumper.Status)
		adcpChecklist.AccuracyShortStatusList = getChecklistStatusList(checklist.AccuracyShort.Status)
		adcpChecklist.CompassCalStatusList = getChecklistStatusList(checklist.CompassCal.Status)
		adcpChecklist.BeamOrientationStatusList = getChecklistStatusList(checklist.BeamOrientation.Status)
		adcpChecklist.PressureTestSystemStatusList = getChecklistStatusList(checklist.PressureTestSystem.Status)
		adcpChecklist.Battery1StatusList = getChecklistStatusList(checklist.Battery1.Status)
		adcpChecklist.Battery2StatusList = getChecklistStatusList(checklist.Battery2.Status)

		t, _ := template.ParseFiles("header.html", "adcp_checklist_update.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpChecklist)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpChecklist)

	} else {
		// Parse the form
		formData, err := forms.Parse(r)
		CheckError(err)

		fmt.Println("Update the ADCP Checklist")

		// Check token
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
			fmt.Println("Good token")
		} else {
			// give error if no token
			fmt.Println("Bad token")
		}

		// Update the checklist in the DB
		if updateAdcpChecklistForm(formData) > 0 {
			// Go to the list of ADCP
			http.Redirect(w, r, "/adcp", http.StatusFound)
		}
	}
}

// Update the checklist data from the form.
func updateAdcpChecklistForm(formData *forms.Data) int {
	// Use data to create a adcp object
	cl := getAdcpChecklist(formData.Get("SerialNumber"))
	cl.SerialNumber = formData.Get("SerialNumber")
	cl.Oring.Status = formData.Get("OringStatus")
	cl.Oring.Date = formData.Get("OringDate")
	cl.Oring.User = formData.Get("OringUser")
	cl.Urethane.Status = formData.Get("UrethaneStatus")
	cl.Urethane.Date = formData.Get("UrethaneDate")
	cl.Urethane.User = formData.Get("UrethaneUser")
	cl.Standoffs.Status = formData.Get("StandoffsStatus")
	cl.Standoffs.Date = formData.Get("StandoffsDate")
	cl.Standoffs.User = formData.Get("StandoffsUser")
	cl.Screws.Status = formData.Get("ScrewsStatus")
	cl.Screws.Date = formData.Get("ScrewsDate")
	cl.Screws.User = formData.Get("ScrewsUser")
	cl.Notch.Status = formData.Get("NotchStatus")
	cl.Notch.Date = formData.Get("NotchDate")
	cl.Notch.User = formData.Get("NotchUser")
	cl.Firmware.Status = formData.Get("FirmwareStatus")
	cl.Firmware.Date = formData.Get("FirmwareDate")
	cl.Firmware.User = formData.Get("FirmwareUser")
	cl.FirmwareVersion = formData.Get("FirmwareVersion")
	cl.FinalCheckLake.Status = formData.Get("FinalCheckLakeStatus")
	cl.FinalCheckLake.Date = formData.Get("FinalCheckLakeDate")
	cl.FinalCheckLake.User = formData.Get("FinalCheckLakeUser")
	cl.FinalCheckShipping.Status = formData.Get("FinalCheckShippingStatus")
	cl.FinalCheckShipping.Date = formData.Get("FinalCheckShippingDate")
	cl.FinalCheckShipping.User = formData.Get("FinalCheckShippingUser")
	cl.BeamWires.Status = formData.Get("BeamWiresStatus")
	cl.BeamWires.Date = formData.Get("BeamWiresDate")
	cl.BeamWires.User = formData.Get("BeamWiresUser")
	cl.ThermistorStranded.Status = formData.Get("ThermistorStrandedStatus")
	cl.ThermistorStranded.Date = formData.Get("ThermistorStrandedDate")
	cl.ThermistorStranded.User = formData.Get("ThermistorStrandedUser")
	cl.Temperature.Status = formData.Get("TemperatureStatus")
	cl.Temperature.Date = formData.Get("TemperatureDate")
	cl.Temperature.User = formData.Get("TemperatureUser")
	cl.Pressure.Status = formData.Get("PressureStatus")
	cl.Pressure.Date = formData.Get("PressureDate")
	cl.Pressure.User = formData.Get("PressureUser")
	cl.PressureSensorSize = formData.Get("PressureSensorSize")
	cl.VaccumTest.Status = formData.Get("VaccumTestStatus")
	cl.VaccumTest.Date = formData.Get("VaccumTestDate")
	cl.VaccumTest.User = formData.Get("VaccumTestUser")
	cl.VibeTest.Status = formData.Get("VibeTestStatus")
	cl.VibeTest.Date = formData.Get("VibeTestDate")
	cl.VibeTest.User = formData.Get("VibeTestUser")
	cl.LakeTest.Status = formData.Get("LakeTestStatus")
	cl.LakeTest.Date = formData.Get("LakeTestDate")
	cl.LakeTest.User = formData.Get("LakeTestUser")
	cl.ReviewLakeTestData.Status = formData.Get("ReviewLakeTestDataStatus")
	cl.ReviewLakeTestData.Date = formData.Get("ReviewLakeTestDataDate")
	cl.ReviewLakeTestData.User = formData.Get("ReviewLakeTestDataUser")
	cl.ReviewLakeTestNotes = formData.Get("ReviewLakeTestNotes")
	cl.TankTest.Status = formData.Get("TankTestStatus")
	cl.TankTest.Date = formData.Get("TankTestDate")
	cl.TankTest.User = formData.Get("TankTestUser")
	cl.ReviewTankTestData.Status = formData.Get("ReviewTankTestDataStatus")
	cl.ReviewTankTestData.Date = formData.Get("ReviewTankTestDataDate")
	cl.ReviewTankTestData.User = formData.Get("ReviewTankTestDataUser")
	cl.ReviewTankTestNotes = formData.Get("ReviewTankTestNotes")
	cl.BurnInTestBoardStack.Status = formData.Get("BurnInTestBoardStackStatus")
	cl.BurnInTestBoardStack.Date = formData.Get("BurnInTestBoardStackDate")
	cl.BurnInTestBoardStack.User = formData.Get("BurnInTestBoardStackUser")
	cl.BurnInTestSystem.Status = formData.Get("BurnInTestSystemStatus")
	cl.BurnInTestSystem.Date = formData.Get("BurnInTestSystemDate")
	cl.BurnInTestSystem.User = formData.Get("BurnInTestSystemUser")
	cl.AccuracyJumper.Status = formData.Get("AccuracyJumperStatus")
	cl.AccuracyJumper.Date = formData.Get("AccuracyJumperDate")
	cl.AccuracyJumper.User = formData.Get("AccuracyJumperUser")
	cl.AccuracyShort.Status = formData.Get("AccuracyShortStatus")
	cl.AccuracyShort.Date = formData.Get("AccuracyShortDate")
	cl.AccuracyShort.User = formData.Get("AccuracyShortUser")
	cl.CompassCal.Status = formData.Get("CompassCalStatus")
	cl.CompassCal.Date = formData.Get("CompassCalDate")
	cl.CompassCal.User = formData.Get("CompassCalUser")
	cl.BeamOrientation.Status = formData.Get("BeamOrientationStatus")
	cl.BeamOrientation.Date = formData.Get("BeamOrientationDate")
	cl.BeamOrientation.User = formData.Get("BeamOrientationUser")
	cl.PressureTestSystem.Status = formData.Get("PressureTestSystemStatus")
	cl.PressureTestSystem.Date = formData.Get("PressureTestSystemDate")
	cl.PressureTestSystem.User = formData.Get("PressureTestSystemUser")
	cl.Battery1.Status = formData.Get("Battery1Status")
	cl.Battery1.Date = formData.Get("Battery1Date")
	cl.Battery1.User = formData.Get("Battery1User")
	cl.Battery1Voltage = formData.Get("Battery1Voltage")
	cl.Battery1LoadVoltage = formData.Get("Battery1LoadVoltage")
	cl.Battery1Rev = formData.Get("Battery1Rev")
	cl.Battery1SerialNumber = formData.Get("Battery1SerialNumber")
	cl.Battery1PartNum = formData.Get("Battery1PartNum")
	cl.Battery1LotNum = formData.Get("Battery1LotNum")
	cl.Battery2.Status = formData.Get("Battery2Status")
	cl.Battery2.Date = formData.Get("Battery2Date")
	cl.Battery2.User = formData.Get("Battery2User")
	cl.Battery2Voltage = formData.Get("Battery2Voltage")
	cl.Battery2LoadVoltage = formData.Get("Battery2LoadVoltage")
	cl.Battery2Rev = formData.Get("Battery2Rev")
	cl.Battery2SerialNumber = formData.Get("Battery2SerialNumber")
	cl.Battery2PartNum = formData.Get("Battery2PartNum")
	cl.Battery2LotNum = formData.Get("Battery2LotNum")
	cl.Modified = time.Now()

	fmt.Println("Serial: ", cl.SerialNumber)

	// Save the data to the DB
	updateAdcpChecklist(cl)

	return 1
}

// Find the Checklist from the database
func getAdcpChecklist(serialNum string) *AdcpChecklist {
	fmt.Println("Get ADCP Checklist - SerialNumber: ", serialNum)

	var data AdcpChecklist
	err := Vault.Mongo.C("AdcpChecklists").Find(bson.M{"SerialNumber": serialNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find ADCP Checklist %v\n", err)
	}

	fmt.Println("SerialNum: ", data.SerialNumber)
	fmt.Println("ID:", data.ID)

	return &data
}

// Update the Checklist data.
func updateAdcpChecklist(data *AdcpChecklist) {
	fmt.Println("updateAdcpChecklist SerialNumber: ", data.SerialNumber)
	fmt.Println("updateAdcpChecklist ID: ", data.ID)

	err := Vault.Mongo.C("AdcpChecklists").Update(bson.M{"_id": data.ID}, bson.M{"$set": bson.M{
		"SerialNumber":         data.SerialNumber,
		"Oring":                data.Oring,
		"Urethane":             data.Urethane,
		"Screws":               data.Screws,
		"Standoffs":            data.Standoffs,
		"Notch":                data.Notch,
		"Firmware":             data.Firmware,
		"FirmwareVersion":      data.FirmwareVersion,
		"FinalCheckLake":       data.FinalCheckLake,
		"FinalCheckShipping":   data.FinalCheckShipping,
		"BeamWires":            data.BeamWires,
		"ThermistorStranded":   data.ThermistorStranded,
		"Temperature":          data.Temperature,
		"Pressure":             data.Pressure,
		"PressureSensorSize":   data.PressureSensorSize,
		"VaccumTest":           data.VaccumTest,
		"VibeTest":             data.VibeTest,
		"LakeTest":             data.LakeTest,
		"ReviewLakeTestData":   data.ReviewLakeTestData,
		"ReviewLakeTestNotes":  data.ReviewLakeTestNotes,
		"TankTest":             data.TankTest,
		"ReviewTankTestData":   data.ReviewTankTestData,
		"ReviewTankTestNotes":  data.ReviewTankTestNotes,
		"BurnInTestBoardStack": data.BurnInTestBoardStack,
		"BurnInTestSystem":     data.BurnInTestSystem,
		"AccuracyShort":        data.AccuracyShort,
		"AccuracyJumper":       data.AccuracyJumper,
		"CompassCal":           data.CompassCal,
		"BeamOrientation":      data.BeamOrientation,
		"PressureTestSystem":   data.PressureTestSystem,
		"Battery1":             data.Battery1,
		"Battery1Voltage":      data.Battery1Voltage,
		"Battery1LoadVoltage":  data.Battery1LoadVoltage,
		"Battery1SerialNumber": data.Battery1SerialNumber,
		"Battery1Rev":          data.Battery1Rev,
		"Battery1PartNum":      data.Battery1PartNum,
		"Battery1LotNum":       data.Battery1LotNum,
		"Battery2":             data.Battery2,
		"Battery2Voltage":      data.Battery2Voltage,
		"Battery2LoadVoltage":  data.Battery2LoadVoltage,
		"Battery2SerialNumber": data.Battery2SerialNumber,
		"Battery2Rev":          data.Battery2Rev,
		"Battery2PartNum":      data.Battery2PartNum,
		"Battery2LotNum":       data.Battery2LotNum,
		"Modified":             data.Modified}})
	if err != nil {
		fmt.Printf("Can't update ADCP Checklist %v\n", err)
	}
}

// CheckAdcpChecklist will check if the checklist has been created.
// If it is, send true, if not, send false.
func CheckAdcpChecklist(serialNum string) {
	fmt.Printf("CheckAdcpChecklist: %s\n", serialNum)

	var data []Adcp
	err := Vault.Mongo.C("AdcpChecklists").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	CheckError(err)

	if len(data) > 0 {
		return
	}

	fmt.Printf("Checklist does not exist: %s\n", serialNum)

	// Add the checklist if it does not exist
	checklist := &AdcpChecklist{
		SerialNumber: serialNum,
		Modified:     time.Now(),
		Created:      time.Now(),
	}
	err = Vault.Mongo.C("AdcpChecklists").Insert(checklist)
	CheckError(err)
}

// GetChecklistStatusList will create a Status slice.  Then set the selected flag
// based off the status value given.
func getChecklistStatusList(status string) []OptionItem {
	options := []OptionItem{
		OptionItem{"Not Checked", "Not Checked", false},
		OptionItem{"Completed", "Completed", false},
		OptionItem{"N/A", "N/A", false},
	}

	// Set the selected value based off the status given
	for i := range options {
		if options[i].Value == status {
			options[i].Selected = true
		}
	}

	return options
}
