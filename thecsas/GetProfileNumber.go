// GetProfileNumberIFM.go
//Function for get the profil number of a data file for thecsas software
package thecsas

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)

func GetProfileNumber(nc *lib.Nc,cfg toml.Configtoml,str string) float64 {
	var value float64
	var err error
	var CruisePrefix string = cfg.Ladcp.CruisePrefix
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