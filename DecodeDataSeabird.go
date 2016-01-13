// DecodeDataSeabird.go
//Function for decode the data for Seabird constructor

package main

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
func DecodeDataSeabird(nc *lib.Nc,m *config.Map,str string, profile float64, file string, line int) {

	// split the string str using whitespace characters
	values := strings.Fields(str)
	fmt.Println("Value : ",values)
	nb_value := len(values)
	fmt.Println("NbValues : ",nb_value)

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
	m.Data["PRFL"] = profile
}