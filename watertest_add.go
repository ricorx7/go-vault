package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/albrow/forms"
)

// WaterTestUpdate will contain the Water Test data.
type WaterTestUpdate struct {
	WaterTest      WaterTestResults
	Token          string
	IsSelected     string
	IsSelectedList []OptionItem
}

// Add the Water Test.
func watertestAddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method watertestAddHandler: %s\n", r.Method) // get request method
	if r.Method == "GET" {
		wtData := &WaterTestUpdate{}

		displayWaterTestTemplate(w, wtData)
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

		// Use data to create a WaterTestResults
		wt := &WaterTestResults{
			SerialNumber:              formData.Get("SerialNumber"),
			TestOrientation:           formData.GetInt("TestOrientation"),
			Beam0NoiseFloor:           formData.GetFloat("Beam0NoiseFloor"),
			Beam1NoiseFloor:           formData.GetFloat("Beam1NoiseFloor"),
			Beam2NoiseFloor:           formData.GetFloat("Beam2NoiseFloor"),
			Beam3NoiseFloor:           formData.GetFloat("Beam3NoiseFloor"),
			Beam0SignalLake:           formData.GetFloat("Beam0SignalLake"),
			Beam1SignalLake:           formData.GetFloat("Beam1SignalLake"),
			Beam2SignalLake:           formData.GetFloat("Beam2SignalLake"),
			Beam3SignalLake:           formData.GetFloat("Beam3SignalLake"),
			Beam0SignalOcean:          formData.GetFloat("Beam0SignalOcean"),
			Beam1SignalOcean:          formData.GetFloat("Beam1SignalOcean"),
			Beam2SignalOcean:          formData.GetFloat("Beam2SignalOcean"),
			Beam3SignalOcean:          formData.GetFloat("Beam3SignalOcean"),
			Beam0SnrLake:              formData.GetFloat("Beam0SnrLake"),
			Beam1SnrLake:              formData.GetFloat("Beam1SnrLake"),
			Beam2SnrLake:              formData.GetFloat("Beam2SnrLake"),
			Beam3SnrLake:              formData.GetFloat("Beam3SnrLake"),
			Beam0SnrOcean:             formData.GetFloat("Beam0SnrOcean"),
			Beam1SnrOcean:             formData.GetFloat("Beam1SnrOcean"),
			Beam2SnrOcean:             formData.GetFloat("Beam2SnrOcean"),
			Beam3SnrOcean:             formData.GetFloat("Beam3SnrOcean"),
			GpsDistance:               formData.Get("GpsDistance"),
			GpsDirection:              formData.Get("GpsDirection"),
			BtDistance:                formData.Get("BtDistance"),
			BtDirection:               formData.Get("BtDirection"),
			DistanceError:             formData.Get("DistanceError"),
			DirectionError:            formData.Get("DirectionError"),
			ProfileRangeBeam0:         formData.GetFloat("ProfileRangeBeam0"),
			ProfileRangeBeam1:         formData.GetFloat("ProfileRangeBeam1"),
			ProfileRangeBeam2:         formData.GetFloat("ProfileRangeBeam2"),
			ProfileRangeBeam3:         formData.GetFloat("ProfileRangeBeam3"),
			GlitchCountBeam0:          formData.GetFloat("GlitchCountBeam0"),
			GlitchCountBeam1:          formData.GetFloat("GlitchCountBeam1"),
			GlitchCountBeam2:          formData.GetFloat("GlitchCountBeam2"),
			GlitchCountBeam3:          formData.GetFloat("GlitchCountBeam3"),
			BottomTrackAmplitudeBeam0: formData.GetFloat("BottomTrackAmplitudeBeam0"),
			BottomTrackAmplitudeBeam1: formData.GetFloat("BottomTrackAmplitudeBeam1"),
			BottomTrackAmplitudeBeam2: formData.GetFloat("BottomTrackAmplitudeBeam2"),
			BottomTrackAmplitudeBeam3: formData.GetFloat("BottomTrackAmplitudeBeam3"),
			PlotReport:                formData.Get("PlotReport"),
			Notes:                     formData.Get("Notes"),
			Created:                   time.Now().Local(),
		}

		// Add the WaterTestResults to the DB
		err = Vault.Mongo.C("WaterTestResults").Insert(wt)
		CheckError(err)

		// Go to the list of WaterTests
		http.Redirect(w, r, "/adcp/wt", http.StatusFound)
	}
}

// Display the template
func displayWaterTestTemplate(w http.ResponseWriter, wtData *WaterTestUpdate) {
	// Generate token
	wtData.Token = genToken()

	t, _ := template.ParseFiles("header.html", "watertest_add.html", "footer.html")
	t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "content", wtData)
	t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, wtData)
}
