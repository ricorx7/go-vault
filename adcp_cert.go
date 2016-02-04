package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-zoo/bone"
)

// AdcpCert holds the ADCP certificate.
type AdcpCert struct {
	Adcp       Adcp
	CompassCal []CompassCal
	Errors     struct {
		SerialNumber string
	}
}

// Update the ADCP data
func adcpCertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		serialNum := bone.GetValue(r, "id")                    // Get the value of the "id" parameters in the URL.
		adcp := getAdcp(serialNum)                             // Get the ADCP data from the DB
		adcpCert := &AdcpCert{Adcp: *adcp}                     // Set the ADCP to struct
		adcpCert.CompassCal = getCompassCalCertData(serialNum) // Get Compass Cal from the DB

		t, _ := template.ParseFiles("header.html", "adcp_cert1.html", "footer.html")
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
