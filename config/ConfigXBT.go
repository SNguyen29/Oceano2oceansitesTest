//ConfigXBT.go
//File for config a instrument XBT

package config

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
	"Oceano2oceansitesTest/lib"
)

type xbt struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
	
}

func GetConfigXBT(nc *lib.Nc,m *Map,configFile string,cfg Config,Type string) {

}
