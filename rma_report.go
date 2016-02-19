package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-zoo/bone"
)

// RmaReport holds the RMA report.
type RmaReport struct {
	RMA    RMA // RMA data
	Errors struct {
		SerialNumber string
	}
}

// Display the RMA Report
func rmaReportHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {

		rmaID := bone.GetValue(r, "id") // Get the value of the "id" parameters in the URL.
		rma := getRma(rmaID)            // Get the RMA
		rmaReport := &RmaReport{}       // Set the RMA Report
		rmaReport.RMA = *rma

		t, _ := template.ParseFiles("header.html", "rma_report.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", rmaReport)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, rmaReport)
	}
}
