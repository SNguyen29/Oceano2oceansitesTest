// GetProfileNumber.go
//Function for get the profil number of a data file for Seabird Constructor
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"Oceano2oceansitesTest/lib"
)

func GetProfileNumber(nc *lib.Nc,str string) float64 {
	var value float64
	var err error
	var CruisePrefix string = cfg.Ctd.CruisePrefix
	if strings.ContainsAny(str,CruisePrefix) {
		res := strings.Split(str,CruisePrefix)
		res = strings.Split(res[1],".")
		if value, err = strconv.ParseFloat(res[0], 64); err == nil {
			// get profile name, eg: csp00101
			nc.Extras_s[fmt.Sprintf("PRFL_NAME:%d", int(value))] = res[0]
		} else {
			log.Fatal(err)
		}

	} else {
		log.Fatal("func GetProfileNumber", err)
	}
	return value

}
