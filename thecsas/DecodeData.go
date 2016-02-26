//DecodeData.go
//Function for decode the data for thecsas software

package thecsas

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
)

// extract data from the line read in str with order gave by hash map_var
// values:  1318 81.583900 3.000 2.983 29.5431 29.5464 5 ...
// map_var: PRES:2 DEPTH:3 PSAL:21 DOX2:18 ...
func DecodeData(nc *lib.Nc,m *config.Map,str string, profile float64, file string, line int) {
	
	var timestop string
	var Poslatstop string
	var Poslongstop string
	
	// split the string str using coma characters
	values := strings.Split(str,",")
	nb_value := len(values)
	
	//extract time
	date := values[2]
	time := values[3]
	
	//extract pos
	lat_s := values[4]
	lat_deg := values[5]
	lat_min := values[6]
	long_s := values[7]
	long_deg := values[8]
	long_min := values[9]
	
	timestop = lib.ConvertDate(date+" "+time)
	Poslatstop = lat_deg + " " + lat_min + " " + lat_s
	Poslongstop = long_deg + " " + long_min + " " + long_s

	// for each physical parameter, extract its data from the rigth column
	// and save it in map data
	for key, column := range m.Map_var {
		
		if column > nb_value {
			log.Fatal(fmt.Sprintf("Error in func DecodeData() "+
				"configuration mismatch\nFound %d values, and we try to use column %d",
				nb_value, column))
		}
		
		if v, err := strconv.ParseFloat(values[column], 64); err == nil {
			m.Data[key] = v
		} else {
			log.Printf("Can't parse line: %d in file: %s\n", line, file)
			log.Fatal(err)
		}
	}
	if line == 0 {
		nc.Extras_s[fmt.Sprintf("Starttime:%d", int(profile))] = timestop
		nc.Extras_s[fmt.Sprintf("Startlatpos:%d", int(profile))] = Poslatstop
		nc.Extras_s[fmt.Sprintf("Startlongpos:%d", int(profile))] = Poslongstop
		}
			
	m.Data["PRFL"] = profile
	if v,err :=  lib.Position2Decimal(Poslatstop); err == nil {
	m.Data["LAT"] = v
	}
	if v,err := lib.Position2Decimal(Poslongstop); err == nil {
	m.Data["LONG"] = v
	}
	
	nc.Extras_s[fmt.Sprintf("Stopttime:%d", int(profile))] = timestop
	nc.Extras_s[fmt.Sprintf("Stoplatpos:%d", int(profile))] = Poslatstop
	nc.Extras_s[fmt.Sprintf("Stoplongpos:%d", int(profile))] = Poslongstop
}
	
	