package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SnrTestData holds the SNR Test data.
type SnrTestData struct {
	SnrTests []SnrTestResults // SNR Tests
	Filter   string           // Serial Number filter
}

// Update the SnrTest in the database.
func updateSnrTest(wt *SnrTestResults) {
	fmt.Println("updateSnrTest - ID", wt.ID)

	err := Vault.Mongo.C("SnrTestResults").Update(bson.M{"_id": wt.ID}, bson.M{"$set": bson.M{
		"SerialNumber":              wt.SerialNumber,
		"IsSelected":                wt.IsSelected,
		"TestOrientation":           wt.TestOrientation,
		"Beam0NoiseFloor":           wt.Beam0NoiseFloor,
		"Beam1NoiseFloor":           wt.Beam1NoiseFloor,
		"Beam2NoiseFloor":           wt.Beam2NoiseFloor,
		"Beam3NoiseFloor":           wt.Beam3NoiseFloor,
		"Beam0SignalLake":           wt.Beam0SignalLake,
		"Beam1SignalLake":           wt.Beam1SignalLake,
		"Beam2SignalLake":           wt.Beam2SignalLake,
		"Beam3SignalLake":           wt.Beam3SignalLake,
		"Beam0SignalOcean":          wt.Beam0SignalOcean,
		"Beam1SignalOcean":          wt.Beam1SignalOcean,
		"Beam2SignalOcean":          wt.Beam2SignalOcean,
		"Beam3SignalOcean":          wt.Beam3SignalOcean,
		"Beam0SnrLake":              wt.Beam0SnrLake,
		"Beam1SnrLake":              wt.Beam1SnrLake,
		"Beam2SnrLake":              wt.Beam2SnrLake,
		"Beam3SnrLake":              wt.Beam3SnrLake,
		"Beam0SnrOcean":             wt.Beam0SnrOcean,
		"Beam1SnrOcean":             wt.Beam1SnrOcean,
		"Beam2SnrOcean":             wt.Beam2SnrOcean,
		"Beam3SnrOcean":             wt.Beam3SnrOcean,
		"GlitchCountBeam0":          wt.GlitchCountBeam0,
		"GlitchCountBeam1":          wt.GlitchCountBeam1,
		"GlitchCountBeam2":          wt.GlitchCountBeam2,
		"GlitchCountBeam3":          wt.GlitchCountBeam3,
		"BottomTrackAmplitudeBeam0": wt.BottomTrackAmplitudeBeam0,
		"BottomTrackAmplitudeBeam1": wt.BottomTrackAmplitudeBeam1,
		"BottomTrackAmplitudeBeam2": wt.BottomTrackAmplitudeBeam2,
		"BottomTrackAmplitudeBeam3": wt.BottomTrackAmplitudeBeam3,
		"PlotReport":                wt.PlotReport,
		"Notes":                     wt.Notes,
		"Modified":                  time.Now()}})
	if err != nil {
		fmt.Printf("Can't update TankTest %v\n", err)
	}
}
