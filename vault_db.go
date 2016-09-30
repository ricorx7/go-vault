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
	RecorderFormatted     bool          `bson:"RecorderFormatted" json:"RecorderFormatted"`
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
	CalScoreStdDevErr      float64       `bson:"CalScoreStdDevErr" json:"CalScoreStdDevErr"`
	CalScoreXCoverage      float64       `bson:"CalScore_xCoverage" json:"CalScore_xCoverage"`
	CalScoreYCoverage      float64       `bson:"CalScore_yCoverage" json:"CalScore_yCoverage"`
	CalScoreZCoverage      float64       `bson:"CalScore_zCoverage" json:"CalScore_zCoverage"`
	CalScoreAccelStdDevErr float64       `bson:"CalScore_accelStdDevErr" json:"CalScore_accelStdDevErr"`
	CalScoreXAccelCoverage float64       `bson:"CalScore_xAccelCoverage" json:"CalScore_xAccelCoverage"`
	CalScoreYAccelCoverage float64       `bson:"CalScore_yAccelCoverage" json:"CalScore_yAccelCoverage"`
	CalScoreZAccelCoverage float64       `bson:"CalScore_zAccelCoverage" json:"CalScore_zAccelCoverage"`
	Point1PreHdg           float64       `bson:"Point1_Pre_Hdg" json:"Point1_Pre_Hdg"`
	Point1PrePtch          float64       `bson:"Point1_Pre_Ptch" json:"Point1_Pre_Ptch"`
	Point1PreRoll          float64       `bson:"Point1_Pre_Roll" json:"Point1_Pre_Roll"`
	Point2PreHdg           float64       `bson:"Point2_Pre_Hdg" json:"Point2_Pre_Hdg"`
	Point2PrePtch          float64       `bson:"Point2_Pre_Ptch" json:"Point2_Pre_Ptch"`
	Point2PreRoll          float64       `bson:"Point2_Pre_Roll" json:"Point2_Pre_Roll"`
	Point3PreHdg           float64       `bson:"Point3_Pre_Hdg" json:"Point3_Pre_Hdg"`
	Point3PrePtch          float64       `bson:"Point3_Pre_Ptch" json:"Point3_Pre_Ptch"`
	Point3PreRoll          float64       `bson:"Point3_Pre_Roll" json:"Point3_Pre_Roll"`
	Point4PreHdg           float64       `bson:"Point4_Pre_Hdg" json:"Point4_Pre_Hdg"`
	Point4PrePtch          float64       `bson:"Point4_Pre_Ptch" json:"Point4_Pre_Ptch"`
	Point4PreRoll          float64       `bson:"Point4_Pre_Roll" json:"Point4_Pre_Roll"`
	Point1PostHdg          float64       `bson:"Point1_Post_Hdg" json:"Point1_Post_Hdg"`
	Point1PostPtch         float64       `bson:"Point1_Post_Ptch" json:"Point1_Post_Ptch"`
	Point1PostRoll         float64       `bson:"Point1_Post_Roll" json:"Point1_Post_Roll"`
	Point2PostHdg          float64       `bson:"Point2_Post_Hdg" json:"Point2_Post_Hdg"`
	Point2PostPtch         float64       `bson:"Point2_Post_Ptch" json:"Point2_Post_Ptch"`
	Point2PostRoll         float64       `bson:"Point2_Post_Roll" json:"Point2_Post_Roll"`
	Point3PostHdg          float64       `bson:"Point3_Post_Hdg" json:"Point3_Post_Hdg"`
	Point3PostPtch         float64       `bson:"Point3_Post_Ptch" json:"Point3_Post_Ptch"`
	Point3PostRoll         float64       `bson:"Point3_Post_Roll" json:"Point3_Post_Roll"`
	Point4PostHdg          float64       `bson:"Point4_Post_Hdg" json:"Point4_Post_Hdg"`
	Point4PostPtch         float64       `bson:"Point4_Post_Ptch" json:"Point4_Post_Ptch"`
	Point4PostRoll         float64       `bson:"Point4_Post_Roll" json:"Point4_Post_Roll"`
	LastModified           float64       `bson:"LastModified" json:"LastModified"`
	Notes                  string        `bson:"Notes" json:"Notes"`
	CompasscalBeam1Error   float64
	CompasscalBeam2Error   float64
	CompasscalBeam3Error   float64
	CompasscalBeam4Error   float64
}

// TankTestResults holds the Tank Test information.
type TankTestResults struct {
	ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IsSelected           bool          `bson:"IsSelected" json:"IsSelected"`
	TankTestType         string        `bson:"TankTestType" json:"TankTestType"`
	Created              time.Time     `bson:"Created" json:"Created"`
	Modified             time.Time     `bson:"Modified" json:"Modified"`
	UserName             string        `bson:"UserName" json:"UserName"`
	SerialNumber         string        `bson:"SerialNumber" json:"SerialNumber"`
	Firmware             string        `bson:"Firmware" json:"Firmware"`
	SubsystemDescStr     string        `bson:"SubsystemDescStr" json:"SubsystemDescStr"`
	SubsystemCode        int           `bson:"SubsystemCode" json:"SubsystemCode"`
	SubsystemCepoIndex   int           `bson:"SubsystemCepoIndex" json:"SubsystemCepoIndex"`
	SubsystemConfigIndex int           `bson:"SubsystemConfigIndex" json:"SubsystemConfigIndex"`
	TestOrientation      int           `bson:"TestOrientation" json:"TestOrientation"`
	Beam0NoiseFloor      float64       `bson:"Beam0NoiseFloor" json:"Beam0NoiseFloor"`
	Beam1NoiseFloor      float64       `bson:"Beam1NoiseFloor" json:"Beam1NoiseFloor"`
	Beam2NoiseFloor      float64       `bson:"Beam2NoiseFloor" json:"Beam2NoiseFloor"`
	Beam3NoiseFloor      float64       `bson:"Beam3NoiseFloor" json:"Beam3NoiseFloor"`
	Beam0Signal1mTank    float64       `bson:"Beam0Signal1mTank" json:"Beam0Signal1mTank"`
	Beam1Signal1mTank    float64       `bson:"Beam1Signal1mTank" json:"Beam1Signal1mTank"`
	Beam2Signal1mTank    float64       `bson:"Beam2Signal1mTank" json:"Beam2Signal1mTank"`
	Beam3Signal1mTank    float64       `bson:"Beam3Signal1mTank" json:"Beam3Signal1mTank"`
	Beam0SignalTank      float64       `bson:"Beam0SignalTank" json:"Beam0SignalTank"`
	Beam1SignalTank      float64       `bson:"Beam1SignalTank" json:"Beam1SignalTank"`
	Beam2SignalTank      float64       `bson:"Beam2SignalTank" json:"Beam2SignalTank"`
	Beam3SignalTank      float64       `bson:"Beam3SignalTank" json:"Beam3SignalTank"`
	Beam0SignalLake      float64       `bson:"Beam0SignalLake" json:"Beam0SignalLake"`
	Beam1SignalLake      float64       `bson:"Beam1SignalLake" json:"Beam1SignalLake"`
	Beam2SignalLake      float64       `bson:"Beam2SignalLake" json:"Beam2SignalLake"`
	Beam3SignalLake      float64       `bson:"Beam3SignalLake" json:"Beam3SignalLake"`
	Beam0SignalOcean     float64       `bson:"Beam0SignalOcean" json:"Beam0SignalOcean"`
	Beam1SignalOcean     float64       `bson:"Beam1SignalOcean" json:"Beam1SignalOcean"`
	Beam2SignalOcean     float64       `bson:"Beam2SignalOcean" json:"Beam2SignalOcean"`
	Beam3SignalOcean     float64       `bson:"Beam3SignalOcean" json:"Beam3SignalOcean"`
	Beam0SnrTank         float64       `bson:"Beam0SnrTank" json:"Beam0SnrTank"`
	Beam1SnrTank         float64       `bson:"Beam1SnrTank" json:"Beam1SnrTank"`
	Beam2SnrTank         float64       `bson:"Beam2SnrTank" json:"Beam2SnrTank"`
	Beam3SnrTank         float64       `bson:"Beam3SnrTank" json:"Beam3SnrTank"`
	Beam0SnrLake         float64       `bson:"Beam0SnrLake" json:"Beam0SnrLake"`
	Beam1SnrLake         float64       `bson:"Beam1SnrLake" json:"Beam1SnrLake"`
	Beam2SnrLake         float64       `bson:"Beam2SnrLake" json:"Beam2SnrLake"`
	Beam3SnrLake         float64       `bson:"Beam3SnrLake" json:"Beam3SnrLake"`
	Beam0SnrOcean        float64       `bson:"Beam0SnrOcean" json:"Beam0SnrOcean"`
	Beam1SnrOcean        float64       `bson:"Beam1SnrOcean" json:"Beam1SnrOcean"`
	Beam2SnrOcean        float64       `bson:"Beam2SnrOcean" json:"Beam2SnrOcean"`
	Beam3SnrOcean        float64       `bson:"Beam3SnrOcean" json:"Beam3SnrOcean"`
	GpsDistance          string        `bson:"GpsDistance" json:"GpsDistance"`
	GpsDirection         string        `bson:"GpsDirection" json:"GpsDirection"`
	BtDistance           string        `bson:"BtDistance" json:"BtDistance"`
	BtDirection          string        `bson:"BtDirection" json:"BtDirection"`
	DistanceError        string        `bson:"DistanceError" json:"DistanceError"`
	DirectionError       string        `bson:"DirectionError" json:"DirectionError"`
	ProfileRangeBeam0    float64       `bson:"ProfileRangeBeam0" json:"ProfileRangeBeam0"`
	ProfileRangeBeam1    float64       `bson:"ProfileRangeBeam1" json:"ProfileRangeBeam1"`
	ProfileRangeBeam2    float64       `bson:"ProfileRangeBeam2" json:"ProfileRangeBeam2"`
	ProfileRangeBeam3    float64       `bson:"ProfileRangeBeam3" json:"ProfileRangeBeam3"`
	GlitchCountBeam0     float64       `bson:"GlitchCountBeam0" json:"GlitchCountBeam0"`
	GlitchCountBeam1     float64       `bson:"GlitchCountBeam1" json:"GlitchCountBeam1"`
	GlitchCountBeam2     float64       `bson:"GlitchCountBeam2" json:"GlitchCountBeam2"`
	GlitchCountBeam3     float64       `bson:"GlitchCountBeam3" json:"GlitchCountBeam3"`
	PlotReport           string        `bson:"PlotReport" json:"PlotReport"`
	Notes                string        `bson:"Notes" json:"Notes"`
}

// WaterTestResults holds the Water Test information.
type WaterTestResults struct {
	ID                        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IsSelected                bool          `bson:"IsSelected" json:"IsSelected"`
	Created                   time.Time     `bson:"Created" json:"Created"`
	Modified                  time.Time     `bson:"Modified" json:"Modified"`
	UserName                  string        `bson:"UserName" json:"UserName"`
	SerialNumber              string        `bson:"SerialNumber" json:"SerialNumber"`
	Firmware                  string        `bson:"Firmware" json:"Firmware"`
	SubsystemDescStr          string        `bson:"SubsystemDescStr" json:"SubsystemDescStr"`
	SubsystemCode             int           `bson:"SubsystemCode" json:"SubsystemCode"`
	SubsystemCepoIndex        int           `bson:"SubsystemCepoIndex" json:"SubsystemCepoIndex"`
	SubsystemConfigIndex      int           `bson:"SubsystemConfigIndex" json:"SubsystemConfigIndex"`
	TestOrientation           int           `bson:"TestOrientation" json:"TestOrientation"`
	Beam0NoiseFloor           float64       `bson:"Beam0NoiseFloor" json:"Beam0NoiseFloor"`
	Beam1NoiseFloor           float64       `bson:"Beam1NoiseFloor" json:"Beam1NoiseFloor"`
	Beam2NoiseFloor           float64       `bson:"Beam2NoiseFloor" json:"Beam2NoiseFloor"`
	Beam3NoiseFloor           float64       `bson:"Beam3NoiseFloor" json:"Beam3NoiseFloor"`
	Beam0SignalLake           float64       `bson:"Beam0SignalLake" json:"Beam0SignalLake"`
	Beam1SignalLake           float64       `bson:"Beam1SignalLake" json:"Beam1SignalLake"`
	Beam2SignalLake           float64       `bson:"Beam2SignalLake" json:"Beam2SignalLake"`
	Beam3SignalLake           float64       `bson:"Beam3SignalLake" json:"Beam3SignalLake"`
	Beam0SignalOcean          float64       `bson:"Beam0SignalOcean" json:"Beam0SignalOcean"`
	Beam1SignalOcean          float64       `bson:"Beam1SignalOcean" json:"Beam1SignalOcean"`
	Beam2SignalOcean          float64       `bson:"Beam2SignalOcean" json:"Beam2SignalOcean"`
	Beam3SignalOcean          float64       `bson:"Beam3SignalOcean" json:"Beam3SignalOcean"`
	Beam0SnrLake              float64       `bson:"Beam0SnrLake" json:"Beam0SnrLake"`
	Beam1SnrLake              float64       `bson:"Beam1SnrLake" json:"Beam1SnrLake"`
	Beam2SnrLake              float64       `bson:"Beam2SnrLake" json:"Beam2SnrLake"`
	Beam3SnrLake              float64       `bson:"Beam3SnrLake" json:"Beam3SnrLake"`
	Beam0SnrOcean             float64       `bson:"Beam0SnrOcean" json:"Beam0SnrOcean"`
	Beam1SnrOcean             float64       `bson:"Beam1SnrOcean" json:"Beam1SnrOcean"`
	Beam2SnrOcean             float64       `bson:"Beam2SnrOcean" json:"Beam2SnrOcean"`
	Beam3SnrOcean             float64       `bson:"Beam3SnrOcean" json:"Beam3SnrOcean"`
	GpsDistance               string        `bson:"GpsDistance" json:"GpsDistance"`
	GpsDirection              string        `bson:"GpsDirection" json:"GpsDirection"`
	BtDistance                string        `bson:"BtDistance" json:"BtDistance"`
	BtDirection               string        `bson:"BtDirection" json:"BtDirection"`
	DistanceError             string        `bson:"DistanceError" json:"DistanceError"`
	DirectionError            string        `bson:"DirectionError" json:"DirectionError"`
	ProfileRangeBeam0         float64       `bson:"ProfileRangeBeam0" json:"ProfileRangeBeam0"`
	ProfileRangeBeam1         float64       `bson:"ProfileRangeBeam1" json:"ProfileRangeBeam1"`
	ProfileRangeBeam2         float64       `bson:"ProfileRangeBeam2" json:"ProfileRangeBeam2"`
	ProfileRangeBeam3         float64       `bson:"ProfileRangeBeam3" json:"ProfileRangeBeam3"`
	GlitchCountBeam0          float64       `bson:"GlitchCountBeam0" json:"GlitchCountBeam0"`
	GlitchCountBeam1          float64       `bson:"GlitchCountBeam1" json:"GlitchCountBeam1"`
	GlitchCountBeam2          float64       `bson:"GlitchCountBeam2" json:"GlitchCountBeam2"`
	GlitchCountBeam3          float64       `bson:"GlitchCountBeam3" json:"GlitchCountBeam3"`
	BottomTrackAmplitudeBeam0 float64       `bson:"BottomTrackAmplitudeBeam0" json:"BottomTrackAmplitudeBeam0"`
	BottomTrackAmplitudeBeam1 float64       `bson:"BottomTrackAmplitudeBeam1" json:"BottomTrackAmplitudeBeam1"`
	BottomTrackAmplitudeBeam2 float64       `bson:"BottomTrackAmplitudeBeam2" json:"BottomTrackAmplitudeBeam2"`
	BottomTrackAmplitudeBeam3 float64       `bson:"BottomTrackAmplitudeBeam3" json:"BottomTrackAmplitudeBeam3"`
	PlotReport                string        `bson:"PlotReport" json:"PlotReport"`
	Notes                     string        `bson:"Notes" json:"Notes"`
}

// SnrTestResults holds the SNR Test information.
type SnrTestResults struct {
	ID                        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IsSelected                bool          `bson:"IsSelected" json:"IsSelected"`
	Created                   time.Time     `bson:"Created" json:"Created"`
	UserName                  string        `bson:"UserName" json:"UserName"`
	SerialNumber              string        `bson:"SerialNumber" json:"SerialNumber"`
	Firmware                  string        `bson:"Firmware" json:"Firmware"`
	SubsystemDescStr          string        `bson:"SubsystemDescStr" json:"SubsystemDescStr"`
	SubsystemCode             int           `bson:"SubsystemCode" json:"SubsystemCode"`
	SubsystemCepoIndex        int           `bson:"SubsystemCepoIndex" json:"SubsystemCepoIndex"`
	SubsystemConfigIndex      int           `bson:"SubsystemConfigIndex" json:"SubsystemConfigIndex"`
	TestOrientation           int           `bson:"TestOrientation" json:"TestOrientation"`
	Beam0NoiseFloor           float64       `bson:"Beam0NoiseFloor" json:"Beam0NoiseFloor"`
	Beam1NoiseFloor           float64       `bson:"Beam1NoiseFloor" json:"Beam1NoiseFloor"`
	Beam2NoiseFloor           float64       `bson:"Beam2NoiseFloor" json:"Beam2NoiseFloor"`
	Beam3NoiseFloor           float64       `bson:"Beam3NoiseFloor" json:"Beam3NoiseFloor"`
	Beam0SignalLake           float64       `bson:"Beam0SignalLake" json:"Beam0SignalLake"`
	Beam1SignalLake           float64       `bson:"Beam1SignalLake" json:"Beam1SignalLake"`
	Beam2SignalLake           float64       `bson:"Beam2SignalLake" json:"Beam2SignalLake"`
	Beam3SignalLake           float64       `bson:"Beam3SignalLake" json:"Beam3SignalLake"`
	Beam0SignalOcean          float64       `bson:"Beam0SignalOcean" json:"Beam0SignalOcean"`
	Beam1SignalOcean          float64       `bson:"Beam1SignalOcean" json:"Beam1SignalOcean"`
	Beam2SignalOcean          float64       `bson:"Beam2SignalOcean" json:"Beam2SignalOcean"`
	Beam3SignalOcean          float64       `bson:"Beam3SignalOcean" json:"Beam3SignalOcean"`
	Beam0SnrLake              float64       `bson:"Beam0SnrLake" json:"Beam0SnrLake"`
	Beam1SnrLake              float64       `bson:"Beam1SnrLake" json:"Beam1SnrLake"`
	Beam2SnrLake              float64       `bson:"Beam2SnrLake" json:"Beam2SnrLake"`
	Beam3SnrLake              float64       `bson:"Beam3SnrLake" json:"Beam3SnrLake"`
	Beam0SnrOcean             float64       `bson:"Beam0SnrOcean" json:"Beam0SnrOcean"`
	Beam1SnrOcean             float64       `bson:"Beam1SnrOcean" json:"Beam1SnrOcean"`
	Beam2SnrOcean             float64       `bson:"Beam2SnrOcean" json:"Beam2SnrOcean"`
	Beam3SnrOcean             float64       `bson:"Beam3SnrOcean" json:"Beam3SnrOcean"`
	GlitchCountBeam0          float64       `bson:"GlitchCountBeam0" json:"GlitchCountBeam0"`
	GlitchCountBeam1          float64       `bson:"GlitchCountBeam1" json:"GlitchCountBeam1"`
	GlitchCountBeam2          float64       `bson:"GlitchCountBeam2" json:"GlitchCountBeam2"`
	GlitchCountBeam3          float64       `bson:"GlitchCountBeam3" json:"GlitchCountBeam3"`
	BottomTrackAmplitudeBeam0 float64       `bson:"BottomTrackAmplitudeBeam0" json:"BottomTrackAmplitudeBeam0"`
	BottomTrackAmplitudeBeam1 float64       `bson:"BottomTrackAmplitudeBeam1" json:"BottomTrackAmplitudeBeam1"`
	BottomTrackAmplitudeBeam2 float64       `bson:"BottomTrackAmplitudeBeam2" json:"BottomTrackAmplitudeBeam2"`
	BottomTrackAmplitudeBeam3 float64       `bson:"BottomTrackAmplitudeBeam3" json:"BottomTrackAmplitudeBeam3"`
	PlotReport                string        `bson:"PlotReport" json:"PlotReport"`
	Notes                     string        `bson:"Notes" json:"Notes"`
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
		fmt.Printf("Can't find ADCP %v\n", err)
	}
	fmt.Println("SerialNum: ", data.SerialNumber)
	fmt.Println("Customer: ", data.Customer)
	fmt.Println("ID:", data.ID)

	return &data
}

// Find all the ADCPs with the serial number partial given.  This will filter the ADCPs down.
func getAdcpContain(serialNumPartial string) *[]Adcp {
	var adcps []Adcp
	err := Vault.Mongo.C("adcps").Find(bson.M{"SerialNumber": bson.M{"$regex": serialNumPartial}}).Sort("-created").All(&adcps)
	if err != nil {
		fmt.Printf("Can't find ADCP Partials %v\n", err)
	}

	return &adcps
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
		"RecorderFormatted":     adcp.RecorderFormatted,
		"RecorderSize":          adcp.RecorderSize,
		"Software":              adcp.Software,
		"SystemType":            adcp.SystemType,
		"TemperaturePresent":    adcp.TemperaturePresent}})
	if err != nil {
		fmt.Printf("Can't update ADCP %v\n", err)
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

// Find the CompassCal from the database.  List only the selected.
func getCompassCalSelected(serialNum string) *[]CompassCal {
	fmt.Println("getcompassCalSelected", serialNum)
	var data []CompassCal
	err := Vault.Mongo.C("CompassCalResults").Find(bson.M{"SerialNumber": serialNum, "IsSelected": true}).All(&data)
	if err != nil {
		fmt.Printf("Can't find CompassCal data %v\n", err)
	}
	fmt.Printf("getCompassCalSelected: %s : Count[%d]\n", serialNum, len(data))
	return &data
}

// Find the CompassCal data from the database based off the ID
func getCompassCalResultsID(id string) *CompassCal {
	var data CompassCal

	err := Vault.Mongo.C("CompassCalResults").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)
	if err != nil {
		fmt.Printf("Can't find from ID CompassCal data %v\n", err)
	}
	return &data
}

// Find the TankTestResults from the database
func getTankTestResults(serialNum string) *[]TankTestResults {
	var data []TankTestResults
	err := Vault.Mongo.C("TankTestResults").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	if err != nil {
		fmt.Printf("Can't find TankTest data %v\n", err)
	}
	return &data
}

// Find the TankTestResults from the database.  List only the selected.
func getTankTestResultsSelected(serialNum string) *[]TankTestResults {
	fmt.Println("getTankTestResultsSelected", serialNum)
	var data []TankTestResults
	err := Vault.Mongo.C("TankTestResults").Find(bson.M{"SerialNumber": serialNum, "IsSelected": true}).All(&data)
	if err != nil {
		fmt.Printf("Can't find TankTest data %v\n", err)
	}
	fmt.Printf("getTankTestResultsSelected: %s : Count[%d]\n", serialNum, len(data))
	return &data
}

// Find the TankTestResults from the database.  List only the selected with the given tank test type.
func getTankTestResultsSelectedType(serialNum string, testType string) *[]TankTestResults {
	fmt.Println("getTankTestResultsSelected", serialNum)
	var data []TankTestResults
	err := Vault.Mongo.C("TankTestResults").Find(bson.M{"SerialNumber": serialNum, "IsSelected": true, "TankTestType": testType}).All(&data)
	if err != nil {
		fmt.Printf("Can't find TankTest data %v\n", err)
	}
	fmt.Printf("getTankTestResultsSelected: %s : Count[%d]\n", serialNum, len(data))
	return &data
}

// Find the TankTestResults from the database based off the ID
func getTankTestResultsID(id string) *TankTestResults {
	var data TankTestResults

	err := Vault.Mongo.C("TankTestResults").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)
	if err != nil {
		fmt.Printf("Can't find from ID TankTest data %v\n", err)
	}
	return &data
}

// Find the WaterTestResults from the database based off the ID
func getWaterTestResultsID(id string) *WaterTestResults {
	var data WaterTestResults

	err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)
	if err != nil {
		fmt.Printf("Can't find from ID WaterTest data %v\n", err)
	}
	return &data
}

// Find the WaterTestResults from the database
func getWaterTestResults(serialNum string) *[]WaterTestResults {
	var data []WaterTestResults
	err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	if err != nil {
		fmt.Printf("Can't find WaterTest data %v\n", err)
	}
	return &data
}

// Find the WaterTestResults from the database.  List only the selected.
func getWaterTestResultsSelected(serialNum string) *[]WaterTestResults {
	fmt.Println("getWaterTestResultsSelected", serialNum)
	var data []WaterTestResults
	err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": serialNum, "IsSelected": true}).All(&data)
	if err != nil {
		fmt.Printf("Can't find WaterTest data %v\n", err)
	}
	fmt.Printf("getWaterTestResultsSelected: %s : Count[%d]\n", serialNum, len(data))
	return &data
}

// Find all the Water Tests with the serial number partial given.  This will filter the Water Tests down.
func getWaterTestContain(serialNumPartial string) *[]WaterTestResults {
	var waterTests []WaterTestResults
	err := Vault.Mongo.C("WaterTestResults").Find(bson.M{"SerialNumber": bson.M{"$regex": serialNumPartial}}).Sort("-Created", "IsSelected").All(&waterTests)
	if err != nil {
		fmt.Printf("Can't find Water Test Partials %v\n", err)
	}

	return &waterTests
}

// Find the SnrTestResults from the database
func getSnrTestResults(serialNum string) *[]SnrTestResults {
	var data []SnrTestResults
	err := Vault.Mongo.C("SnrTestResults").Find(bson.M{"SerialNumber": serialNum}).All(&data)
	if err != nil {
		fmt.Printf("Can't find SnrTest data %v\n", err)
	}
	return &data
}

// Find the SnrTestResults from the database.  List only the selected.
func getSnrTestResultsSelected(serialNum string) *[]SnrTestResults {
	fmt.Println("getSnrTestResultsSelected", serialNum)
	var data []SnrTestResults
	err := Vault.Mongo.C("SnrTestResults").Find(bson.M{"SerialNumber": serialNum, "IsSelected": true}).All(&data)
	if err != nil {
		fmt.Printf("Can't find SnrTest data %v\n", err)
	}
	fmt.Printf("getSnrTestResultsSelected: %s : Count[%d]\n", serialNum, len(data))
	return &data
}

// Find the SnrTestResults from the database based off the ID
func getSnrTestResultsID(id string) *SnrTestResults {
	var data SnrTestResults

	err := Vault.Mongo.C("SnrTestResults").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)
	if err != nil {
		fmt.Printf("Can't find from ID SnrTest data %v\n", err)
	}
	return &data
}
