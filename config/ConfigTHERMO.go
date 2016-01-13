//ConfigTHERMO.go
//File for config a instrument THERMO

package config

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
	"Oceano2oceansitesTest/lib"
)

type thermo struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
}

func GetConfigTHERMO(nc *lib.Nc,m *Map,configFile string,cfg Config,Type string) {

}