//DecoderHeader.go
//Function for decode the header of a data file from thecsas software

package thecsas

import (
	"fmt"
	"strconv"
	//"strings"
	"regexp"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)

// parse header line from .cnv and extract correct information
// use regular expression
// to parse time with non standard format, see:
// http://golang.org/src/time/format.go

func  DecodeHeader(nc *lib.Nc,cfg toml.Configtoml,str string, profile float64,optDebug *bool) {
	
	regDate := regexp.MustCompile(cfg.Ifm.Date)
	regStartTime := regexp.MustCompile(cfg.Ifm.StartTime)
	regLatitude := regexp.MustCompile(cfg.Ifm.Latitude)
	regLongitude := regexp.MustCompile(cfg.Ifm.Longitude)

	switch {
		
		case regDate.MatchString(str) :
			res := regDate.FindStringSubmatch(str)
			year := res[1]
			month := res[2]
			day := res[3]	
			value := (day+"/"+month+"/"+year)
			nc.Extras_s["DATE"] = value
		
		case regStartTime.MatchString(str) :
			res := regStartTime.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintf(lib.Debug, "%s -> ", value)
			nc.Extras_s["HEURE"] = value
	
		case regLatitude.MatchString(str) :
			res := regLatitude.FindStringSubmatch(str)
			if v, err := strconv.ParseFloat(res[1],64); err == nil {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), 1e36)
			}
					
		case regLongitude.MatchString(str) :
			res := regLongitude.FindStringSubmatch(str)
			if v, err := strconv.ParseFloat(res[1],64); err == nil {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), 1e36)
			}
	}
}