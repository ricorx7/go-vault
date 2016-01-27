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
