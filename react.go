package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func reactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method reactHander:", r.Method) // get request method
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
		t, _ := template.ParseFiles("header.html", "react.html", "footer.html")
		t.ExecuteTemplate(w, "header", nil)
		t.ExecuteTemplate(w, "content", waterTestData)
		t.ExecuteTemplate(w, "footer", nil)
		t.Execute(w, waterTestData)
	}
}
