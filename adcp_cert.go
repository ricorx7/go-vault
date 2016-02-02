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

		// Get the value of the "id" parameters.
		val := bone.GetValue(r, "id")
		adcp := getAdcp(val)

		adcpCert := &AdcpCert{Adcp: *adcp}
		cc := getCompassCalSelected(val)
		adcpCert.CompassCal = *cc
		if len(adcpCert.CompassCal) > 0 {
			fmt.Println("Compass Cal: ", adcpCert.CompassCal[0].Point1PostHdg)
		}
		//adcpCert.CompassCal = *cc

		t, _ := template.ParseFiles("adcp_cert.html", "footer.html")
		//t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", adcpCert)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, adcpCert)
	}
}
