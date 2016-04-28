package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/albrow/forms"

	"gopkg.in/mgo.v2/bson"
)

// WaterTestData holds the Water Test data.
type WaterTestData struct {
	WaterTests []WaterTestResults // WaterTests
	Filter     string             // Serial Number filter
}

// List all the WaterTests.
func watertestListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method waterTestListHandler:", r.Method) // get request method
	if r.Method == "GET" {

		// Init
		waterTestData := &WaterTestData{}

		// Get data form DB
		err := Vault.Mongo.C("WaterTestResults").Find(bson.M{}).Sort("-Created").All(&waterTestData.WaterTests)
		CheckError(err)
		fmt.Println("Number of WaterTests: ", len(waterTestData.WaterTests))

		// Get the path to the PlotModel
		for index, element := range waterTestData.WaterTests {
			waterTestData.WaterTests[index].PlotReport = getWaterTestPlotModelPath(element.PlotReport, element.SerialNumber)
		}

		// Display data to page
		t, _ := template.ParseFiles("header.html", "watertest_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", waterTestData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, waterTestData)
	} else {

		// Init
		waterTestData := &WaterTestData{}

		// Get the partial serial number
		formData, err := forms.Parse(r)
		CheckError(err)
		var partialSerial = formData.Get("PartialSerialNumber")
		waterTestData.WaterTests = *getWaterTestContain(partialSerial)
		fmt.Println("Number of WaterTests: ", len(waterTestData.WaterTests))

		// Get the path to the PlotModel
		for index, element := range waterTestData.WaterTests {
			waterTestData.WaterTests[index].PlotReport = getWaterTestPlotModelPath(element.PlotReport, element.SerialNumber)
		}

		// Display data to page
		t, _ := template.ParseFiles("header.html", "watertest_list.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", waterTestData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, waterTestData)
	}
}

// Append the file name and the vault drive, to give the link to the
// actual file.
func getWaterTestPlotModelPath(path string, serialNum string) string {
	_, file := filepath.Split(path)
	//vault := "\\\\Vault\\Vault\\"
	vault := "/vault/"

	filePath := vault + serialNum + "\\" + file

	return filePath
}
