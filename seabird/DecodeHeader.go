// DecodeHeaderSeabird.go
//Function for decode the header of a data file from Seabird constructor

package seabird

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)

// parse header line from .cnv and extract correct information
// use regular expression
// to parse time with non standard format, see:
// http://golang.org/src/time/format.go

func  DecodeHeader(nc *lib.Nc,cfg toml.Configtoml,str string, profile float64,optDebug *bool) {
	
	regCruise := regexp.MustCompile(cfg.Seabird.Cruise)
	regShip := regexp.MustCompile(cfg.Seabird.Ship)
	regStation := regexp.MustCompile(cfg.Seabird.Station)
	regType := regexp.MustCompile(cfg.Seabird.Type)
	regOperator := regexp.MustCompile(cfg.Seabird.Operator)
	regBottomDepth := regexp.MustCompile(cfg.Seabird.BottomDepth)
	regDummyBottomDepth := regexp.MustCompile(cfg.Seabird.DummyBottomDepth)
	regSystemTime := regexp.MustCompile(cfg.Seabird.SystemTime)
	regNmeaLatitude := regexp.MustCompile(cfg.Seabird.Latitude)
	regNmeaLongitude := regexp.MustCompile(cfg.Seabird.Longitude)

	switch {
	// decode Systeme Upload Time		
		case regSystemTime.MatchString(str) : 
			res := regSystemTime.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintf(lib.Debug, "%s -> ", value)
			// create new Time object, see tools.go
			var t = lib.NewTimeFromString("Jan 02 2006 15:04:05", value)
			v := t.Time2JulianDec()
			nc.Variables_1D["TIME"] = append(nc.Variables_1D["TIME"].([]float64),v)
	
		case regNmeaLatitude.MatchString(str) :
			if v, err := lib.Position2Decimal(str); err == nil {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), v)
			} else {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), 1e36)
			}
			
			
		case regNmeaLongitude.MatchString(str) :
			if v, err := lib.Position2Decimal(str); err == nil {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), v)
				fmt.Println(v)
			} else {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), 1e36)
			}
			
			
		case regCruise.MatchString(str) :
			res := regCruise.FindStringSubmatch(str)
			value := res[1]
			fmt.Println(value)
			fmt.Fprintln(lib.Debug, value)
			nc.Attributes["cycle_mesure"] = value

		case regShip.MatchString(str) :
			res := regShip.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintln(lib.Debug, value)
			nc.Attributes["plateforme"] = value
			fmt.Println(value)
			
		case regStation.MatchString(str) :
			res := regStation.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintln(lib.Debug, v)
				// ch
				//			format := "%0" + cfg.Ctd.StationPrefixLength + ".0f"
				//			p := fmt.Sprintf(format, profile)
				//			//s := fmt.Sprintf(format, v)
				//			fmt.Println(p, v)
				//			if p != v {
				//				fmt.Printf("Warning: profile for header differ from file name: %s <=> %s\n", p, v)
				//			}
				nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64), profile)
			} else {
				nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64), 1e36)
			}
			fmt.Println(value)

		case regBottomDepth.MatchString(str) :
			res := regBottomDepth.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintf(lib.Debug, "Bath: %f\n", v)
				nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), v)
			} else {
				fmt.Fprintf(lib.Debug, "Bath: %f\n", v)
				nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), 1e36)
			}
			fmt.Println(value)
			
		case regDummyBottomDepth.MatchString(str) ://?
			nc.Variables_1D["BATH"] = append(nc.Variables_1D["BATH"].([]float64), 1e36)
			fmt.Fprintf(lib.Debug, "Bath: %g\n", 1e36)
			fmt.Println("1e36")

		case regOperator.MatchString(str) :
			res := regOperator.FindStringSubmatch(str)
			value := res[1]
			if *optDebug {
				fmt.Println(value)
			}
			fmt.Println(value)
	
	// TODOS: uncomment, add optionnal value from seabird header	
			
		case regType.MatchString(str) :
			res := regType.FindStringSubmatch(str)
			value := strings.ToUpper(res[1]) // convert to upper case
			var v float64
			switch value {
			case "PHY":
				v = float64(lib.PHY)
			case "GEO":
				v = float64(lib.GEO)
			case "BIO":
				v = float64(lib.BIO)
			default:
				v = float64(lib.UNKNOW)
			}
			fmt.Println(value, v)
			nc.Variables_1D["TYPECAST"] = append(nc.Variables_1D["TYPECAST"].([]float64), v)

			nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = value
		 
	
	}
}
