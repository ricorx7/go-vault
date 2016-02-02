package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Adcp holds the ADCP information.
type Adcp struct {
	ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	SerialNumber          string        `bson:"SerialNumber" json:"SerialNumber"`
	Customer              string        `bson:"Customer" json:"Customer"`
	OrderNumber           string        `bson:"OrderNumber" json:"OrderNumber"`
	DepthRating           string        `bson:"DepthRating" json:"DepthRating"`
	HeadType              string        `bson:"HeadType" json:"HeadType"`
	Hardware              string        `bson:"Hardware" json:"Hardware"`
	ConnectorType         string        `bson:"ConnectorType" json:"ConnectorType"`
	Frequency             string        `bson:"Frequency" json:"Frequency"`
	Firmware              string        `bson:"Firmware" json:"Firmware"`
	Software              string        `bson:"Software" json:"Software"`
	TemperaturePresent    bool          `bson:"TemperaturePresent" json:"TemperaturePresent"`
	PressureSensorPresent bool          `bson:"PressureSensorPresent" json:"PressureSensorPresent"`
	PressureSensorRating  string        `bson:"PressureSensorRating" json:"PressureSensorRating"`
	RecorderSize          string        `bson:"RecorderSize" json:"RecorderSize"`
	RecorderFormated      bool          `bson:"RecorderFormated" json:"RecorderFormated"`
	SystemType            string        `bson:"SystemType" json:"SystemType"`
	Application           string        `bson:"Application" json:"Application"`
	Created               time.Time     `bson:"created"`
	Modified              time.Time     `bson:"modified"`
	Name                  string        `bson:"name"`
	User                  bson.ObjectId `bson:"user"`
}

// CompassCal holds the CompassCal information.
type CompassCal struct {
	ID                     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IsSelected             bool          `bson:"IsSelected" json:"IsSelected"`
	Created                time.Time     `bson:"Created" json:"Created"`
	UserName               string        `bson:"UserName" json:"UserName"`
	SerialNumber           string        `bson:"SerialNumber" json:"SerialNumber"`
	Firmware               string        `bson:"Firmware" json:"Firmware"`
	CalScoreStdDevErr      float32       `bson:"IsSelected" json:"IsSelected"`
	CalScoreXCoverage      float32       `bson:"CalScore_xCoverage" json:"CalScore_xCoverage"`
	CalScoreYCoverage      float32       `bson:"CalScore_yCoverage" json:"CalScore_yCoverage"`
	CalScoreZCoverage      float32       `bson:"CalScore_zCoverage" json:"CalScore_zCoverage"`
	CalScoreAccelStdDevErr float32       `bson:"CalScore_accelStdDevErr" json:"CalScore_accelStdDevErr"`
	CalScoreXAccelCoverage float32       `bson:"CalScore_xAccelCoverage" json:"CalScore_xAccelCoverage"`
	CalScoreYAccelCoverage float32       `bson:"CalScore_yAccelCoverage" json:"CalScore_yAccelCoverage"`
	CalScoreZAccelCoverage float32       `bson:"CalScore_zAccelCoverage" json:"CalScore_zAccelCoverage"`
	Point1PreHdg           float32       `bson:"Point1_Pre_Hdg" json:"Point1_Pre_Hdg"`
	Point1PrePtch          float32       `bson:"Point1_Pre_Ptch" json:"Point1_Pre_Ptch"`
	Point1PreRoll          float32       `bson:"Point1_Pre_Roll" json:"Point1_Pre_Roll"`
	Point2PreHdg           float32       `bson:"Point2_Pre_Hdg" json:"Point2_Pre_Hdg"`
	Point2PrePtch          float32       `bson:"Point2_Pre_Ptch" json:"Point2_Pre_Ptch"`
	Point2PreRoll          float32       `bson:"Point2_Pre_Roll" json:"Point2_Pre_Roll"`
	Point3PreHdg           float32       `bson:"Point3_Pre_Hdg" json:"Point3_Pre_Hdg"`
	Point3PrePtch          float32       `bson:"Point3_Pre_Ptch" json:"Point3_Pre_Ptch"`
	Point3PreRoll          float32       `bson:"Point3_Pre_Roll" json:"Point3_Pre_Roll"`
	Point4PreHdg           float32       `bson:"Point4_Pre_Hdg" json:"Point4_Pre_Hdg"`
	Point4PrePtch          float32       `bson:"Point4_Pre_Ptch" json:"Point4_Pre_Ptch"`
	Point4PreRoll          float32       `bson:"Point4_Pre_Roll" json:"Point4_Pre_Roll"`
	Point1PostHdg          float32       `bson:"Point1_Post_Hdg" json:"Point1_Post_Hdg"`
	Point1PostPtch         float32       `bson:"Point1_Post_Ptch" json:"Point1_Post_Ptch"`
	Point1PostRoll         float32       `bson:"Point1_Post_Roll" json:"Point1_Post_Roll"`
	Point2PostHdg          float32       `bson:"Point2_Post_Hdg" json:"Point2_Post_Hdg"`
	Point2PostPtch         float32       `bson:"Point2_Post_Ptch" json:"Point2_Post_Ptch"`
	Point2PostRoll         float32       `bson:"Point2_Post_Roll" json:"Point2_Post_Roll"`
	Point3PostHdg          float32       `bson:"Point3_Post_Hdg" json:"Point3_Post_Hdg"`
	Point3PostPtch         float32       `bson:"Point3_Post_Ptch" json:"Point3_Post_Ptch"`
	Point3PostRoll         float32       `bson:"Point3_Post_Roll" json:"Point3_Post_Roll"`
	Point4PostHdg          float32       `bson:"Point4_Post_Hdg" json:"Point4_Post_Hdg"`
	Point4PostPtch         float32       `bson:"Point4_Post_Ptch" json:"Point4_Post_Ptch"`
	Point4PostRoll         float32       `bson:"Point4_Post_Roll" json:"Point4_Post_Roll"`
	CompasscalBeam1Error   float32
	CompasscalBeam2Error   float32
	CompasscalBeam3Error   float32
	CompasscalBeam4Error   float32
	LastModified           float32 `bson:"LastModified" json:"LastModified"`
}

// VaultDb holds the vault database.
type VaultDb struct {
	Mongo *mgo.Database // Mongo DB server connection
}

// Vault is the vault connection.
var Vault = &VaultDb{}

// DbConnect will connect the to the database.
func DbConnect(server string) {
	fmt.Println("Connect to MongoDB: ", server)
	session, err := mgo.Dial(server)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Vault")
		panic(err)
	}
	//defer session.Close()

	// Set the connection
	Vault.Mongo = session.DB("Vault")

}

// Find the ADCP from the database
func getAdcp(serialNum string) *Adcp {
	fmt.Println("Get ADCP - SerialNum: ", serialNum)

	var data Adcp
	err := Vault.Mongo.C("adcps").Find(bson.M{"SerialNumber": serialNum}).One(&data)
	if err != nil {
		fmt.Printf("Can't find document %v\n", err)
	}
	fmt.Println("SerialNum: ", data.SerialNumber)
	fmt.Println("Customer: ", data.Customer)
	fmt.Println("ID:", data.ID)

	return &data
}

func updateAdcp(adcp *Adcp) {
	fmt.Println("UpdateAdcp - ID", adcp.ID)

	//err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp._id}, bson.M{"$inc": bson.M{"Customer": adcp.Customer}})
	err := Vault.Mongo.C("adcps").Update(bson.M{"_id": adcp.ID}, bson.M{"$set": bson.M{"Customer": adcp.Customer,
		"OrderNumber":           adcp.OrderNumber,
		"Application":           adcp.Application,
		"ConnectorType":         adcp.ConnectorType,
		"DepthRating":           adcp.DepthRating,
		"Firmware":              adcp.Firmware,
		"Hardware":              adcp.Hardware,
		"HeadType":              adcp.HeadType,
		"Modified":              adcp.Modified,
		"PressureSensorPresent": adcp.PressureSensorPresent,
		"PressureSensorRating":  adcp.PressureSensorRating,
		"RecorderFormated":      adcp.RecorderFormated,
		"RecorderSize":          adcp.RecorderSize,
		"Software":              adcp.Software,
		"SystemType":            adcp.SystemType,
		"TemperaturePresent":    adcp.TemperaturePresent}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
	}
}

// Find the CompassCal from the database
func getCompassCal(serialNum string) *[]CompassCal {
	var data []CompassCal
	err := Vault.Mongo.C("CompassCalResults").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	if err != nil {
		fmt.Printf("Can't find CompassCal data %v\n", err)
	}
	return &data
}

// Find the CompassCal from the database
func getCompassCalSelected(serialNum string) *[]CompassCal {
	fmt.Println("getcompassCalSelected", serialNum)

	var data []CompassCal
	err := Vault.Mongo.C("CompassCalResults").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	if err != nil {
		fmt.Printf("Can't find CompassCal data %v\n", err)
	}
	return &data
}
