package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

// Get the ADCP from the given serial number value.
func vaultAPIAdcpSerialGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			adcp := getAdcp(id)

			fmt.Printf("Get ADCP from serial: %s  %v", id, adcp)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(adcp); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Get the ADCP Cert info the given serial number value.
func vaultAPIAdcpCertGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	serialNum := bone.GetValue(r, "id") // Get the value of the "id" parameters in the URL.

	switch r.Method {
	case "GET":
		{
			adcp := getAdcp(serialNum)                                 // Get the ADCP data from the DB
			adcpCert := &AdcpCert{Adcp: *adcp}                         // Set the ADCP to struct
			adcpCert.CompassCal = getCompassCalCertData(serialNum)     // Get Compass Cal from the DB
			adcpCert.TankTest = getTankTestResultCertData(serialNum)   // Get Tank Test from the DB
			adcpCert.SnrTest = getSnrTestResultCertData(serialNum)     // Get SNR Test from the DB
			adcpCert.WaterTest = getWaterTestResultCertData(serialNum) // Get Water Test from the DB

			fmt.Printf("Get ADCP from serial: %s  %v", serialNum, adcpCert)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(adcpCert); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Get the compass cal from the given serial number value.
func vaultAPICompassCalSerialGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			compasscal := getCompassCal(id)

			fmt.Printf("Get CompassCal from serial: %s  %v", id, compasscal)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(compasscal); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Get the Tank Test data from the Vault.
func vaultAPITankHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	var waterTestData []WaterTestResults

	// Get data form DB
	err := Vault.Mongo.C("TankTestResults").Find(bson.M{}).Sort("-Created").All(&waterTestData)
	CheckError(err)
	fmt.Println("Number of TankTests: ", len(waterTestData))

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

	// Get the limit
	limit := 0
	if len(r.URL.Query().Get("limit")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err == nil {
			limit = val
		}
	}

	// Get the offset
	offset := 0
	if len(r.URL.Query().Get("offset")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err == nil {
			offset = val
		}
	}

	// Get the filter
	// Get the data from the database
	filter := ""
	if len(r.URL.Query().Get("filter")) > 0 {
		filter = r.URL.Query().Get("filter")
		err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": &bson.RegEx{Pattern: filter}}).Sort("-Created").All(&waterTestData.WaterTests)
		CheckError(err)
	} else {
		err := Vault.Mongo.C("WaterTestResults").Find(bson.M{}).Skip(offset).Limit(limit).Sort("-Created").All(&waterTestData.WaterTests)
		CheckError(err)
	}

	// Get data form DB
	//err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": filter}).Skip(offset).Limit(limit).Sort("-Created").All(&waterTestData.WaterTests)
	//CheckError(err)
	fmt.Println("Number of WaterTests: ", len(waterTestData.WaterTests))
	fmt.Printf("limit: %s\n", r.URL.Query().Get("limit"))
	fmt.Printf("offset: %s\n", r.URL.Query().Get("offset"))
	fmt.Printf("filter: %s\n", r.URL.Query().Get("filter"))

	// Get the path to the PlotModel
	for index, element := range waterTestData.WaterTests {
		waterTestData.WaterTests[index].PlotReport = getWaterTestPlotModelPath(element.PlotReport, element.SerialNumber)
	}

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(waterTestData); err != nil {
		CheckError(err)
		panic(err)
	}
}

// Set the Water Test Selected value.  This will invert the value that is in the database.
func vaultAPIWaterTestSelectGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			watertest := getWaterTestResultsID(id)
			watertest.IsSelected = !watertest.IsSelected // Invert the value

			// Pass the data back to the database
			updateWaterTest(watertest)

			fmt.Printf("given waterest: %v\n", watertest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(watertest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Edit the Water Test data from the vault.
func vaultAPIWaterTestEditHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			// ID is the ADCP serial number

			watertest := getWaterTestResultsID(id) // Get the data from the database

			fmt.Printf("Edit waterest: %v\n", watertest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			// Pass back as a JSON
			if err := json.NewEncoder(w).Encode(watertest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{
			// ID is the database ID

			// Verify the data exist
			if r.Body == nil {
				http.Error(w, "Send a request body", 400)
				fmt.Printf("Edit waterest 1: \n")
				return
			}

			defer r.Body.Close()

			// Read in the data
			fmt.Println("response Headers:", r.Header)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("response Body:", string(body))

			// Convert to JSON
			var wt WaterTestResults
			err := json.Unmarshal(body, &wt)
			if err != nil {
				fmt.Println("Error with unmarsharl: ", err)
			}

			fmt.Printf("POST Watertest: %v\n", wt)

			// Store the new data to the database
			updateWaterTest(&wt)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	default:
		{

		}

	}
}

// Get the Tank Test data from the vault.
func vaultAPITankTestGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	tankTestData := &TankTestData{}

	// Get the limit
	limit := 0
	if len(r.URL.Query().Get("limit")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err == nil {
			limit = val
		}
	}

	// Get the offset
	offset := 0
	if len(r.URL.Query().Get("offset")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err == nil {
			offset = val
		}
	}

	// Get the filter
	// Get the data from the database
	filter := ""
	if len(r.URL.Query().Get("filter")) > 0 {
		filter = r.URL.Query().Get("filter")
		err := Vault.Mongo.C("TankTestResults").Find(bson.M{"SerialNumber": &bson.RegEx{Pattern: filter}}).Sort("-Created").All(&tankTestData.TankTests)
		CheckError(err)
	} else {
		err := Vault.Mongo.C("TankTestResults").Find(bson.M{}).Skip(offset).Limit(limit).Sort("-Created").All(&tankTestData.TankTests)
		CheckError(err)
	}

	// Get data form DB
	//err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": filter}).Skip(offset).Limit(limit).Sort("-Created").All(&waterTestData.WaterTests)
	//CheckError(err)
	fmt.Println("Number of TankTests: ", len(tankTestData.TankTests))
	fmt.Printf("limit: %s\n", r.URL.Query().Get("limit"))
	fmt.Printf("offset: %s\n", r.URL.Query().Get("offset"))
	fmt.Printf("filter: %s\n", r.URL.Query().Get("filter"))

	// Get the path to the PlotModel
	for index, element := range tankTestData.TankTests {
		tankTestData.TankTests[index].PlotReport = getWaterTestPlotModelPath(element.PlotReport, element.SerialNumber)
	}

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tankTestData); err != nil {
		CheckError(err)
		panic(err)
	}
}

// Set the Tank Test Selected value.  This will invert the value that is in the database.
func vaultAPITankTestSelectGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			tanktest := getTankTestResultsID(id)
			tanktest.IsSelected = !tanktest.IsSelected // Invert the value

			// Pass the data back to the database
			updateTankTest(tanktest)

			fmt.Printf("given tankTest: %v\n", tanktest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(tanktest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Edit the Tank Test data from the vault.
func vaultAPITankTestEditHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			// ID is the ADCP serial number

			tanktest := getTankTestResultsID(id) // Get the data from the database

			fmt.Printf("Edit tanktest: %v\n", tanktest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			// Pass back as a JSON
			if err := json.NewEncoder(w).Encode(tanktest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{
			// ID is the database ID

			// Verify the data exist
			if r.Body == nil {
				http.Error(w, "Send a request body", 400)
				fmt.Printf("Edit tanktest 1: \n")
				return
			}

			defer r.Body.Close()

			// Read in the data
			fmt.Println("response Headers:", r.Header)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("response Body:", string(body))

			// Convert to JSON
			var tt TankTestResults
			err := json.Unmarshal(body, &tt)
			if err != nil {
				fmt.Println("Error with unmarsharl: ", err)
			}

			fmt.Printf("POST tanktest: %v\n", tt)

			// Store the new data to the database
			updateTankTest(&tt)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	default:
		{

		}

	}
}

// Get the SNR Test data from the vault.
func vaultAPISnrTestGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	snrTestData := &SnrTestData{}

	// Get the limit
	limit := 0
	if len(r.URL.Query().Get("limit")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err == nil {
			limit = val
		}
	}

	// Get the offset
	offset := 0
	if len(r.URL.Query().Get("offset")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err == nil {
			offset = val
		}
	}

	// Get the filter
	// Get the data from the database
	filter := ""
	if len(r.URL.Query().Get("filter")) > 0 {
		filter = r.URL.Query().Get("filter")
		err := Vault.Mongo.C("SnrTestResults").Find(bson.M{"SerialNumber": &bson.RegEx{Pattern: filter}}).Sort("-Created").All(&snrTestData.SnrTests)
		CheckError(err)
	} else {
		err := Vault.Mongo.C("SnrTestResults").Find(bson.M{}).Skip(offset).Limit(limit).Sort("-Created").All(&snrTestData.SnrTests)
		CheckError(err)
	}

	// Get data form DB
	//err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": filter}).Skip(offset).Limit(limit).Sort("-Created").All(&waterTestData.WaterTests)
	//CheckError(err)
	fmt.Println("Number of SnrTests: ", len(snrTestData.SnrTests))
	fmt.Printf("limit: %s\n", r.URL.Query().Get("limit"))
	fmt.Printf("offset: %s\n", r.URL.Query().Get("offset"))
	fmt.Printf("filter: %s\n", r.URL.Query().Get("filter"))

	// Get the path to the PlotModel
	for index, element := range snrTestData.SnrTests {
		snrTestData.SnrTests[index].PlotReport = getWaterTestPlotModelPath(element.PlotReport, element.SerialNumber)
	}

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(snrTestData); err != nil {
		CheckError(err)
		panic(err)
	}
}

// Set the SNR Test Selected value.  This will invert the value that is in the database.
func vaultAPISnrTestSelectGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			snrtest := getSnrTestResultsID(id)
			snrtest.IsSelected = !snrtest.IsSelected // Invert the value

			// Pass the data back to the database
			updateSnrTest(snrtest)

			fmt.Printf("given SnrTest: %v\n", snrtest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(snrtest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Edit the SNR Test data from the vault.
func vaultAPISnrTestEditHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			// ID is the ADCP serial number

			snrtest := getSnrTestResultsID(id) // Get the data from the database

			fmt.Printf("Edit snrtest: %v\n", snrtest)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			// Pass back as a JSON
			if err := json.NewEncoder(w).Encode(snrtest); err != nil {
				panic(err)
			}
		}
	case "POST":
		{
			// ID is the database ID

			// Verify the data exist
			if r.Body == nil {
				http.Error(w, "Send a request body", 400)
				fmt.Printf("Edit snrtest 1: \n")
				return
			}

			defer r.Body.Close()

			// Read in the data
			fmt.Println("response Headers:", r.Header)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("response Body:", string(body))

			// Convert to JSON
			var snr SnrTestResults
			err := json.Unmarshal(body, &snr)
			if err != nil {
				fmt.Println("Error with unmarsharl: ", err)
			}

			fmt.Printf("POST snrtest: %v\n", snr)

			// Store the new data to the database
			updateSnrTest(&snr)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	default:
		{

		}

	}
}

// Get the Compass Cal data from the vault.
func vaultAPICompassCalGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init
	compassCalData := &CompassCalData{}

	// Get the limit
	limit := 0
	if len(r.URL.Query().Get("limit")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err == nil {
			limit = val
		}
	}

	// Get the offset
	offset := 0
	if len(r.URL.Query().Get("offset")) > 0 {
		val, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err == nil {
			offset = val
		}
	}

	// Get the filter
	// Get the data from the database
	filter := ""
	if len(r.URL.Query().Get("filter")) > 0 {
		filter = r.URL.Query().Get("filter")
		err := Vault.Mongo.C("CompassCalResults").Find(bson.M{"SerialNumber": &bson.RegEx{Pattern: filter}}).Sort("-Created").All(&compassCalData.CompassCals)
		CheckError(err)
	} else {
		err := Vault.Mongo.C("CompassCalResults").Find(bson.M{}).Skip(offset).Limit(limit).Sort("-Created").All(&compassCalData.CompassCals)
		CheckError(err)
	}

	// Remove NaN from data
	// for i := range compassCalData.CompassCals {
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreAccelStdDevErr) {
	// 		fmt.Printf("NaN found at CalScoreAccelStdDevErr: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreStdDevErr) {
	// 		fmt.Printf("NaN found at CalScoreStdDevErr: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreXAccelCoverage) {
	// 		fmt.Printf("NaN found at CalScoreXAccelCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreXCoverage) {
	// 		fmt.Printf("NaN found at CalScoreXCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreYAccelCoverage) {
	// 		fmt.Printf("NaN found at CalScoreYAccelCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreYCoverage) {
	// 		fmt.Printf("NaN found at CalScoreYCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreZAccelCoverage) {
	// 		fmt.Printf("NaN found at CalScoreZAccelCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CalScoreZCoverage) {
	// 		fmt.Printf("NaN found at CalScoreZCoverage: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// 	if math.IsNaN(compassCalData.CompassCals[i].CompasscalBeam1Error) {
	// 		fmt.Printf("NaN found at CompasscalBeam1Error: %s %s\n", i, compassCalData.CompassCals[i].SerialNumber)
	// 	}
	// }
	//fmt.Printf("%v", compassCalData.CompassCal)
	// b, err := json.Marshal(compassCalData.CompassCal)
	// b = bytes.Replace(b, []byte(":nan"), []byte(":0"), -1)
	// err = json.Unmarshal(b, &compassCalData.CompassCal)
	// CheckError(err)

	fmt.Println("Number of CompassCals: ", len(compassCalData.CompassCals))
	fmt.Printf("limit: %s\n", r.URL.Query().Get("limit"))
	fmt.Printf("offset: %s\n", r.URL.Query().Get("offset"))
	fmt.Printf("filter: %s\n", r.URL.Query().Get("filter"))

	// Set data type and OK status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(compassCalData); err != nil {
		CheckError(err)
		//panic(err)
	}
}

// Set the Compass Cal Selected value.  This will invert the value that is in the database.
func vaultAPICompassCalSelectGetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			compassCal := getCompassCalResultsID(id)
			compassCal.IsSelected = !compassCal.IsSelected // Invert the value

			// Pass the data back to the database
			updateCompassCal(compassCal)

			fmt.Printf("given CompassCal: %v\n", compassCal)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(compassCal); err != nil {
				panic(err)
			}
		}
	case "POST":
		{

		}
	default:
		{

		}

	}
}

// Edit the Compass Cal data from the vault.
func vaultAPICompassCalEditHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameters.
	id := bone.GetValue(r, "id")

	switch r.Method {
	case "GET":
		{
			// ID is the ADCP serial number

			compassCal := getCompassCalResultsID(id) // Get the data from the database

			fmt.Printf("Edit compassCal: %v\n", compassCal)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			// Pass back as a JSON
			if err := json.NewEncoder(w).Encode(compassCal); err != nil {
				panic(err)
			}
		}
	case "POST":
		{
			// ID is the database ID

			// Verify the data exist
			if r.Body == nil {
				http.Error(w, "Send a request body", 400)
				return
			}

			defer r.Body.Close()

			// Read in the data
			fmt.Println("response Headers:", r.Header)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("response Body:", string(body))

			// Convert to JSON
			var compassCal CompassCal
			err := json.Unmarshal(body, &compassCal)
			if err != nil {
				fmt.Println("Error with unmarsharl: ", err)
			}

			fmt.Printf("POST CompassCal: %v\n", compassCal)

			// Store the new data to the database
			updateCompassCal(&compassCal)

			// Set data type and OK status
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	default:
		{

		}

	}
}

