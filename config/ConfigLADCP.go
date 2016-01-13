//ConfigLADCP.go
//File for config a instrument LADCP

package config

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
	"Oceano2oceansitesTest/lib"
)

type ladcp struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
	
}

func GetConfigLADCP(nc *lib.Nc,m *Map,configFile string,cfg Config,Type string) {

}