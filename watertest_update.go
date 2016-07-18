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

// Update the WaterTest.
func watertestUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method watertestUpdateHandler: %s\n", r.Method) // get request method
	if r.Method == "GET" {

		// Get the value of the "id" parameters.
		id := bone.GetValue(r, "id")

		watertest := getWaterTestResultsID(id)
		watertestData := &WaterTestUpdate{}
		watertestData.WaterTest = *watertest
		watertestData.IsSelectedList = getSelectedList(watertest.IsSelected)

		fmt.Printf("%v\n", watertestData.WaterTest)
		fmt.Printf("SerialNum: %v\n", watertestData.WaterTest.SerialNumber)
		fmt.Printf("GpsDistance: %v\n", watertestData.WaterTest.GpsDistance)

		displayWaterTestUpdateTemplate(w, watertestData)
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
		val.Require("ID")

		fmt.Printf("ID: %s\n", formData.Get("ID"))

		//wt := getWaterTestResultsID(bson.ObjectId(formData.Get("ID")))
		wt := getWaterTestResultsID(formData.Get("ID"))
		wt.SerialNumber = formData.Get("SerialNumber")
		wt.TestOrientation = formData.GetInt("TestOrientation")
		wt.Beam0NoiseFloor = formData.GetFloat("Beam0NoiseFloor")
		wt.Beam1NoiseFloor = formData.GetFloat("Beam1NoiseFloor")
		wt.Beam2NoiseFloor = formData.GetFloat("Beam2NoiseFloor")
		wt.Beam3NoiseFloor = formData.GetFloat("Beam3NoiseFloor")
		wt.Beam0SignalLake = formData.GetFloat("Beam0SignalLake")
		wt.Beam1SignalLake = formData.GetFloat("Beam1SignalLake")
		wt.Beam2SignalLake = formData.GetFloat("Beam2SignalLake")
		wt.Beam3SignalLake = formData.GetFloat("Beam3SignalLake")
		wt.Beam0SignalOcean = formData.GetFloat("Beam0SignalOcean")
		wt.Beam1SignalOcean = formData.GetFloat("Beam1SignalOcean")
		wt.Beam2SignalOcean = formData.GetFloat("Beam2SignalOcean")
		wt.Beam3SignalOcean = formData.GetFloat("Beam3SignalOcean")
		wt.Beam0SnrLake = formData.GetFloat("Beam0SnrLake")
		wt.Beam1SnrLake = formData.GetFloat("Beam1SnrLake")
		wt.Beam2SnrLake = formData.GetFloat("Beam2SnrLake")
		wt.Beam3SnrLake = formData.GetFloat("Beam3SnrLake")
		wt.Beam0SnrOcean = formData.GetFloat("Beam0SnrOcean")
		wt.Beam1SnrOcean = formData.GetFloat("Beam1SnrOcean")
		wt.Beam2SnrOcean = formData.GetFloat("Beam2SnrOcean")
		wt.Beam3SnrOcean = formData.GetFloat("Beam3SnrOcean")
		wt.GpsDistance = formData.Get("GpsDistance")
		wt.GpsDirection = formData.Get("GpsDirection")
		wt.BtDistance = formData.Get("BtDistance")
		wt.BtDirection = formData.Get("BtDirection")
		wt.DistanceError = formData.Get("DistanceError")
		wt.DirectionError = formData.Get("DirectionError")
		wt.ProfileRangeBeam0 = formData.GetFloat("ProfileRangeBeam0")
		wt.ProfileRangeBeam1 = formData.GetFloat("ProfileRangeBeam1")
		wt.ProfileRangeBeam2 = formData.GetFloat("ProfileRangeBeam2")
		wt.ProfileRangeBeam3 = formData.GetFloat("ProfileRangeBeam3")
		wt.GlitchCountBeam0 = formData.GetFloat("GlitchCountBeam0")
		wt.GlitchCountBeam1 = formData.GetFloat("GlitchCountBeam1")
		wt.GlitchCountBeam2 = formData.GetFloat("GlitchCountBeam2")
		wt.GlitchCountBeam3 = formData.GetFloat("GlitchCountBeam3")
		wt.BottomTrackAmplitudeBeam0 = formData.GetFloat("BottomTrackAmplitudeBeam0")
		wt.BottomTrackAmplitudeBeam1 = formData.GetFloat("BottomTrackAmplitudeBeam1")
		wt.BottomTrackAmplitudeBeam2 = formData.GetFloat("BottomTrackAmplitudeBeam2")
		wt.BottomTrackAmplitudeBeam3 = formData.GetFloat("BottomTrackAmplitudeBeam3")
		wt.PlotReport = formData.Get("PlotReport")
		wt.Notes = formData.Get("Notes")
		wt.Modified = time.Now().Local()

		// Get selected checkbox
		isChecked := formData.Get("IsSelected")
		fmt.Printf("WaterTest selected: %v\n", isChecked)
		if isChecked == "Selected" {
			wt.IsSelected = true
		} else {
			wt.IsSelected = false
		}

		fmt.Printf("WaterTest Update: %s\n", wt.SerialNumber)

		// Update the WaterTest in DB
		updateWaterTest(wt)

		// Go to the list of Products
		http.Redirect(w, r, "/adcp/wt", http.StatusFound)
	}
}

// Display the template
func displayWaterTestUpdateTemplate(w http.ResponseWriter, data *WaterTestUpdate) {
	// Generate token
	data.Token = genToken()

	// Redirect back to the page with error message
	t, _ := template.ParseFiles("header.html", "watertest_update.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", data)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, data)
}

// Update the WaterTest in the database.
func updateWaterTest(wt *WaterTestResults) {
	fmt.Println("updateWaterTest - ID", wt.ID)

	err := Vault.Mongo.C("WaterTestResults").Update(bson.M{"_id": wt.ID}, bson.M{"$set": bson.M{
		"SerialNumber":              wt.SerialNumber,
		"IsSelected":                wt.IsSelected,
		"TestOrientation":           wt.TestOrientation,
		"Beam0NoiseFloor":           wt.Beam0NoiseFloor,
		"Beam1NoiseFloor":           wt.Beam1NoiseFloor,
		"Beam2NoiseFloor":           wt.Beam2NoiseFloor,
		"Beam3NoiseFloor":           wt.Beam3NoiseFloor,
		"Beam0SignalLake":           wt.Beam0SignalLake,
		"Beam1SignalLake":           wt.Beam1SignalLake,
		"Beam2SignalLake":           wt.Beam2SignalLake,
		"Beam3SignalLake":           wt.Beam3SignalLake,
		"Beam0SignalOcean":          wt.Beam0SignalOcean,
		"Beam1SignalOcean":          wt.Beam1SignalOcean,
		"Beam2SignalOcean":          wt.Beam2SignalOcean,
		"Beam3SignalOcean":          wt.Beam3SignalOcean,
		"Beam0SnrLake":              wt.Beam0SnrLake,
		"Beam1SnrLake":              wt.Beam1SnrLake,
		"Beam2SnrLake":              wt.Beam2SnrLake,
		"Beam3SnrLake":              wt.Beam3SnrLake,
		"Beam0SnrOcean":             wt.Beam0SnrOcean,
		"Beam1SnrOcean":             wt.Beam1SnrOcean,
		"Beam2SnrOcean":             wt.Beam2SnrOcean,
		"Beam3SnrOcean":             wt.Beam3SnrOcean,
		"GpsDistance":               wt.GpsDistance,
		"GpsDirection":              wt.GpsDirection,
		"BtDistance":                wt.BtDistance,
		"BtDirection":               wt.BtDirection,
		"DistanceError":             wt.DistanceError,
		"DirectionError":            wt.DirectionError,
		"ProfileRangeBeam0":         wt.ProfileRangeBeam0,
		"ProfileRangeBeam1":         wt.ProfileRangeBeam1,
		"ProfileRangeBeam2":         wt.ProfileRangeBeam2,
		"ProfileRangeBeam3":         wt.ProfileRangeBeam3,
		"GlitchCountBeam0":          wt.GlitchCountBeam0,
		"GlitchCountBeam1":          wt.GlitchCountBeam1,
		"GlitchCountBeam2":          wt.GlitchCountBeam2,
		"GlitchCountBeam3":          wt.GlitchCountBeam3,
		"BottomTrackAmplitudeBeam0": wt.BottomTrackAmplitudeBeam0,
		"BottomTrackAmplitudeBeam1": wt.BottomTrackAmplitudeBeam1,
		"BottomTrackAmplitudeBeam2": wt.BottomTrackAmplitudeBeam2,
		"BottomTrackAmplitudeBeam3": wt.BottomTrackAmplitudeBeam3,
		"PlotReport":                wt.PlotReport,
		"Notes":                     wt.Notes,
		"Modified":                  time.Now()}})
	if err != nil {
		fmt.Printf("Can't update WaterTest %v\n", err)
	}
}

// getSelectedList will create a IsSelected slice.  Then set the selected flag
// based off the isSelected flag value given.
func getSelectedList(isSelected bool) []OptionItem {
	options := []OptionItem{
		OptionItem{"Not Selected", "Not Selected", false},
		OptionItem{"Selected", "Selected", false},
	}

	// Set the selected value based off the status given
	if isSelected {
		options[0].Selected = false
		options[1].Selected = true
	} else {
		options[0].Selected = true
		options[1].Selected = false
	}

	return options
}
