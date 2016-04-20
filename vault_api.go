package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"

	"gopkg.in/mgo.v2/bson"
)

// Get the ADCP data from the Vault.
func vaultAPIAdcpGet(w http.ResponseWriter, r *http.Request) {
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
