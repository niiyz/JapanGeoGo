package main

import (
	"fmt"
	"github.com/niiyz/JapanCityGeoJson/geo"
	"io/ioutil"
	"os"
)

// Data All Pref GeoJson
const GEO_JSON_JAPAN string = "./japan.json"

// City
func makeCityGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_CITY, raw)

	// Loop
	for city, fts := range cityMap {
		// Pref Name
		pref := fts[0].GetPref()
		fmt.Println(pref, city, cap(fts))
		// Save Dir
		dir := "geojson/" + pref
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + city + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

// County
func makeCountyGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_COUNTY, raw)

	// Loop
	for county, fts := range cityMap {
		// Pref Name
		pref := fts[0].GetPref()
		fmt.Println(county, cap(fts))
		// Save Dir
		dir := "geojson/" + pref
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + county + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

// Pref
func makePrefGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_PREF, raw)

	// Loop
	for pref, fts := range cityMap {
		fmt.Println(pref, cap(fts))
		// Save Dir
		dir := "geojson/47都道府県"
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + pref + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

func reset() {
	os.RemoveAll("geojson")
}

func main() {
	raw, err := ioutil.ReadFile(GEO_JSON_JAPAN)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	reset()
	// City
	makeCityGeoJson(raw)
	// Pref
	makePrefGeoJson(raw)
	// County
	makeCountyGeoJson(raw)
}

// Check Dir Exist
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
