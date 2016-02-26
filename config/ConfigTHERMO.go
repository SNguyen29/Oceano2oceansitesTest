//ConfigTHERMO.go
//File for config a instrument THERMO

package config

import (
	"fmt"
	"strconv"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/roscop"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
)

type thermo struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
}

func GetConfigTHERMO(nc *lib.Nc,m *Map,cfg toml.Configtoml,configFile string,Type string,optAll *bool) {
	//	var split, header, format string
	var split []string

	// define map from netcdf structure
	nc.Dimensions = make(map[string]int)
	nc.Attributes = make(map[string]string)
	nc.Extras_f = make(map[string]float64)
	nc.Extras_s = make(map[string]string)
	nc.Variables_1D = make(map[string]interface{})

	// initialize map entry from nil interface to empty slice of float64
	nc.Variables_1D["PROFILE"] = []float64{}
	nc.Variables_1D["TIME"] = []float64{}
	nc.Variables_1D["LATITUDE"] = []float64{}
	nc.Variables_1D["LONGITUDE"] = []float64{}
	nc.Variables_1D["TYPECAST"] = []float64{}
	nc.Roscop = roscop.NewRoscop(cfg.Roscopfile)

	// add some global attributes for profile, change in future
	nc.Attributes["data_type"] = Type

			split = cfg.Thermo.Split
		
		//		stationPrefixLength = cfg.Ctd.StationPrefixLength
		// TODOS: complete
		nc.Attributes["cycle_mesure"] = cfg.Cruise.CycleMesure
		nc.Attributes["plateforme"] = cfg.Cruise.Plateforme
		nc.Attributes["institute"] = cfg.Cruise.Institute
		nc.Attributes["pi"] = cfg.Cruise.Pi
		nc.Attributes["timezone"] = cfg.Cruise.Timezone
		nc.Attributes["begin_date"] = cfg.Cruise.BeginDate
		nc.Attributes["end_date"] = cfg.Cruise.EndDate
		nc.Attributes["creator"] = cfg.Cruise.Creator
		nc.Attributes["type_instrument"] = cfg.Thermo.TypeInstrument
		nc.Attributes["instrument_number"] = cfg.Thermo.InstrumentNumber

	

	// add specific column(s) to the first header line in ascii file
	
		// First column should be PRFL
		m.Hdr = append(m.Hdr, "PRFL")
	

	// fill map_var from split (read in .ini configuration file)
	// store the position (column) of each physical parameter
	var fields []string
	
		fields = split
		
	fmt.Fprintln(lib.Debug, "getConfig: ", fields)

	// construct header slice from split
	for i := 0; i < len(fields); i += 2 {
		if v, err := strconv.Atoi(fields[i+1]); err == nil {
			m.Map_var[fields[i]] = v - 1
			m.Hdr = append(m.Hdr, fields[i])
		}
	}
	
	fmt.Fprintln(lib.Debug, "getConfig: ", m.Hdr)

	// fill map_format from code_roscop
	for _, key := range m.Hdr {
		m.Map_format[key] = nc.Roscop.GetAttributesM(key,"format")
	}
}