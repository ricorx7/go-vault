package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-zoo/bone"
)

// AdcpCert holds the ADCP certificate.
type AdcpCert struct {
	Adcp       Adcp               // ADCP data
	CompassCal []CompassCal       // Compass Cal data
	TankTest   []TankTestResults  // Tank Test data
	SnrTest    []SnrTestResults   // SNR Test data
	WaterTest  []WaterTestResults // Water Tesd data
	Errors     struct {
		SerialNumber string
	}
}

// Update the ADCP data
func adcpCertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		serialNum := bone.GetValue(r, "id")                        // Get the value of the "id" parameters in the URL.
		adcp := getAdcp(serialNum)                                 // Get the ADCP data from the DB
		adcpCert := &AdcpCert{Adcp: *adcp}                         // Set the ADCP to struct
		adcpCert.CompassCal = getCompassCalCertData(serialNum)     // Get Compass Cal from the DB
		adcpCert.TankTest = getTankTestResultCertData(serialNum)   // Get Tank Test from the DB
		adcpCert.SnrTest = getSnrTestResultCertData(serialNum)     // Get SNR Test from the DB
		adcpCert.WaterTest = getWaterTestResultCertData(serialNum) // Get Water Test from the DB

		t, _ := template.ParseFiles("header.html", "adcp_cert.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpCert)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpCert)
	}
}

// Get the compass cal data from the DB and enter in the error values.
func getCompassCalCertData(serialNum string) []CompassCal {
	cc := getCompassCalSelected(serialNum)
	for i := 0; i < len(*cc); i++ {
		(*cc)[i].CompasscalBeam1Error = 0.0 - (*cc)[i].Point1PostHdg
		(*cc)[i].CompasscalBeam2Error = 90.0 - (*cc)[i].Point2PostHdg
		(*cc)[i].CompasscalBeam3Error = 180.0 - (*cc)[i].Point3PostHdg
		(*cc)[i].CompasscalBeam4Error = 270.0 - (*cc)[i].Point4PostHdg
	}

	return *cc
}

// Get the tank test data from the DB.
func getTankTestResultCertData(serialNum string) []TankTestResults {
	tt := getTankTestResultsSelectedType(serialNum, "Noise")
	return *tt
}

// Get the Water test data from the DB.
func getWaterTestResultCertData(serialNum string) []WaterTestResults {
	wt := getWaterTestResultsSelected(serialNum)
	return *wt
}

// Get the SNR test data from the DB.
func getSnrTestResultCertData(serialNum string) []SnrTestResults {
	st := getSnrTestResultsSelected(serialNum)
	return *st
}
