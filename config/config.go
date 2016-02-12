//config.go
//File with struct config who need to be change when a new type of instrument is add

package config

import (
	"io/ioutil"
	"io"
	)

// use for debug mode
var debug io.Writer = ioutil.Discard

// use for echo mode
var echo io.Writer = ioutil.Discard


type Map struct {
	Map_var 	map[string]int
	Map_format 	map[string]string
	Data		map[string]interface{}
	Hdr 		[]string
	}


type Config struct {
	Global struct {
		Author string
		Debug  bool
		Echo   bool
	}
	Cruise struct {
		CycleMesure string
		Plateforme  string
		Callsign    string
		Institute   string
		Pi          string
		Timezone    string
		BeginDate   string
		EndDate     string
		Creator     string
	}
	Ctd ctd
	Btl btl
	Xbt xbt
	Thermo thermo
	Ladcp ladcp
	Sadcp sadcp
	//add new type of instrument
}

//function for init struct Map
func InitMap() Map{
	
	var m Map
	
	// Create an empty map.
	m.Map_var = map[string]int{}
 	m.Map_format = map[string]string{}
	m.Data = make(map[string]interface{})
	
	return m
	}
	
//function to get Data from struct Map
func GetData(data interface{})float64{
		return data.(float64)
	}

func GetPhysicalParametersList(m *Map) []string {
	return m.Hdr
}