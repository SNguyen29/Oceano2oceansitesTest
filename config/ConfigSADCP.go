//ConfigSADCP.go
//File for config a instrument SADCP

package config

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
)

type sadcp struct {

	TypeInstrument      string
	InstrumentNumber    string	
	
}

func GetConfigSADCP(nc *lib.Nc,m *Map,configFile string,cfg Config,Type string) {

}