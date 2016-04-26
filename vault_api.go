package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"

	"gopkg.in/mgo.v2/bson"
)

// Get the ADCP data from the Vault.
func vaultAPIAdcpGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	adcpData := &AdcpData{}

	// Get data form DB
	err := Vault.Mongo.C("adcps").Find(bson.M{}).Sort("-created").All(&adcpData.Adcps)
	CheckError(err)
	fmt.Println("Number of ADCPs: ", len(adcpData.Adcps))

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(adcpData); err != nil {
		panic(err)
	}
}

// Get the Tank Test data from the Vault.
func vaultAPITankHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	var waterTestData []WaterTestResults

	// Get data form DB
	err := Vault.Mongo.C("TankTestResults").Find(bson.M{}).Sort("-Created").All(&waterTestData)
	CheckError(err)
	fmt.Println("Number of WaterTests: ", len(waterTestData))

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(waterTestData); err != nil {
		panic(err)
	}
}

// Get the Tank Test data with the given serial number from the vault.
func vaultAPITankSerialHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serial := bone.GetValue(r, "id")

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(getTankTestResults(serial)); err != nil {
		panic(err)
	}
}

// Get the Tank Test data with the given serial number from the vault.
func vaultAPITankSelectedSerialHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serial := bone.GetValue(r, "id")

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(getTankTestResults(serial)); err != nil {
		panic(err)
	}
}

// Get the Moving Tank Test data with the given serial number from the vault.
func vaultAPITankSelectedSerialMovingHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serial := bone.GetValue(r, "id")

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(getTankTestResultsSelectedType(serial, "Moving")); err != nil {
		panic(err)
	}
}

// Get the Noise Tank Test data with the given serial number from the vault.
func vaultAPITankSelectedSerialNoiseHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serial := bone.GetValue(r, "id")

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(getTankTestResultsSelectedType(serial, "Noise")); err != nil {
		panic(err)
	}
}

// Get the Ringing Tank Test data with the given serial number from the vault.
func vaultAPITankSelectedSerialRingingHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serial := bone.GetValue(r, "id")

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(getTankTestResultsSelectedType(serial, "Ringing")); err != nil {
		panic(err)
	}
}

// Get the Water Test data from the vault.
func vaultAPIWaterTestGetHandler(w http.ResponseWriter, r *http.Request) {
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

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(waterTestData); err != nil {
		panic(err)
	}
}

// Set the Water Test Selected value.  This will invert the value that is in the database.
func vaultAPIWaterTestSelectGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	watertest := getWaterTestResultsID(id)
	watertest.IsSelected = !watertest.IsSelected // Invert the value

	// Pass the data back to the database
	updateWaterTest(watertest)

	fmt.Printf("%v\n", watertest)

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(watertest); err != nil {
		panic(err)
	}
}
