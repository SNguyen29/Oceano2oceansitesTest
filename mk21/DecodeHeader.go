//DecoderHeader.go
//Function for decode the header of a data file from MK21 constructor

package mk21

import (
	"fmt"
	"regexp"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
)

// parse header line from .cnv and extract correct information
// use regular expression
// to parse time with non standard format, see:
// http://golang.org/src/time/format.go

func  DecodeHeader(nc *lib.Nc,cfg toml.Configtoml,str string, profile float64,optDebug *bool) {
	
	regDate := regexp.MustCompile(cfg.Mk21.Date)
	regTime := regexp.MustCompile(cfg.Mk21.Time)
	regLatitude := regexp.MustCompile(cfg.Mk21.Latitude)
	regLongitude := regexp.MustCompile(cfg.Mk21.Longitude)

	switch {
		
		case regDate.MatchString(str) :
			res := regDate.FindStringSubmatch(str)
			month := res[1]
			day := res[2]
			year := res[3]	
			value := (day+"/"+month+"/"+year)
			nc.Extras_s["DATE"] = value
		
		case regTime.MatchString(str) :
			res := regTime.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintf(lib.Debug, "%s -> ", value)
			nc.Extras_s["HEURE"] = value
	
		case regLatitude.MatchString(str) :	
			if v, err := lib.Position3Decimal(str); err == nil {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), 1e36)
			}
					
		case regLongitude.MatchString(str) :
			if v, err := lib.Position3Decimal(str); err == nil {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), 1e36)
			}
	}
}