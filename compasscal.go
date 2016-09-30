package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CompassCalData holds the Tank Test data.
type CompassCalData struct {
	CompassCals []CompassCal // CompassCal
	Filter      string       // Serial Number filter
}

// Update the CompassCal in the database.
func updateCompassCal(cc *CompassCal) {
	fmt.Println("updateCompassCal - ID", cc.ID)

	err := Vault.Mongo.C("CompassCalResults").Update(bson.M{"_id": cc.ID}, bson.M{"$set": bson.M{
		"SerialNumber":            cc.SerialNumber,
		"IsSelected":              cc.IsSelected,
		"UserName":                cc.UserName,
		"Firmware":                cc.Firmware,
		"CalScoreStdDevErr":       cc.CalScoreStdDevErr,
		"CalScore_XCoverage":      cc.CalScoreXCoverage,
		"CalScore_YCoverage":      cc.CalScoreYCoverage,
		"CalScore_ZCoverage":      cc.CalScoreZCoverage,
		"CalScore_AccelStdDevErr": cc.CalScoreAccelStdDevErr,
		"CalScore_XAccelCoverage": cc.CalScoreXAccelCoverage,
		"CalScore_YAccelCoverage": cc.CalScoreYAccelCoverage,
		"CalScore_ZAccelCoverage": cc.CalScoreZAccelCoverage,
		"Point1_Pre_Hdg":          cc.Point1PreHdg,
		"Point1_Pre_Ptch":         cc.Point1PrePtch,
		"Point1_Pre_Roll":         cc.Point1PreRoll,
		"Point2_Pre_Hdg":          cc.Point2PreHdg,
		"Point2_Pre_Ptch":         cc.Point2PrePtch,
		"Point2_Pre_Roll":         cc.Point2PreRoll,
		"Point3_Pre_Hdg":          cc.Point3PreHdg,
		"Point3_Pre_Ptch":         cc.Point3PrePtch,
		"Point3_Pre_Roll":         cc.Point3PreRoll,
		"Point4_Pre_Hdg":          cc.Point4PreHdg,
		"Point4_Pre_Ptch":         cc.Point4PrePtch,
		"Point4_Pre_Roll":         cc.Point4PreRoll,
		"Point1_Post_Hdg":         cc.Point1PostHdg,
		"Point1_Post_Ptch":        cc.Point1PostPtch,
		"Point1_Post_Roll":        cc.Point1PostRoll,
		"Point2_Post_Hdg":         cc.Point2PostHdg,
		"Point2_Post_Ptch":        cc.Point2PostPtch,
		"Point2_Post_Roll":        cc.Point2PostRoll,
		"Point3_Post_Hdg":         cc.Point3PostHdg,
		"Point3_Post_Ptch":        cc.Point3PostPtch,
		"Point3_Post_Roll":        cc.Point3PostRoll,
		"Point4_Post_Hdg":         cc.Point4PostHdg,
		"Point4_Post_Ptch":        cc.Point4PostPtch,
		"Point4_Post_Roll":        cc.Point4PostRoll,
		"Notes":                   cc.Notes,
		"Modified":                time.Now()}})

	if err != nil {
		fmt.Printf("Can't update CompassCal %v\n", err)
	}
}
