
package main

import (
	"fmt"
	"log"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/config"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
	"io"
	"io/ioutil"
	"github.com/SNguyen29/Oceano2oceansitesTest/analyze"
	"github.com/SNguyen29/Oceano2oceansitesTest/seabird"
	"github.com/SNguyen29/Oceano2oceansitesTest/ifm"
	"github.com/SNguyen29/Oceano2oceansitesTest/thecsas"
	"github.com/SNguyen29/Oceano2oceansitesTest/mk21"
	
) 

var filetoml = "configfile/configtoml.toml"

// file prefix for --all option: "-all" for all parameters, "" empty by default
var prefixAll = ""

// use for debug mode
var debug io.Writer = ioutil.Discard
// use for echo mode
var echo io.Writer = ioutil.Discard

// usefull macro
var p = fmt.Println
var f = fmt.Printf


var cfg toml.Configtoml
var filestruct analyze.Structfile

type AllData_2D lib.AllData_2D

var nc lib.Nc
var m config.Map

// main body
func main() {
	
	//init variable cfg with config file in TOML
	cfg = toml.InitToml(filetoml)
	//init variable m with empty map 
	m = config.InitMap()
	
	var files []string
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	files, optCfgfile := GetOptions(filetoml)
	
	//analyse the file to know contructor, instrument and instrument type
	filestruct = analyze.AnalyzeFile(cfg,files)
	
	switch{
		
		//case constructor == seabird
		case filestruct.Constructeur.Number == 0 :
			seabird.ReadSeabird(&nc,&m,filestruct,cfg,files,optCfgfile,optAll,optDebug,prefixAll)
		//case constructor == IFM-GEOMAR
		case filestruct.Constructeur.Number == 1 :
			ifm.Read(&nc,&m,filestruct,cfg,files,optCfgfile,optAll,optDebug,prefixAll)
		//case constructor == THECSAS
		case filestruct.Constructeur.Number == 2 :
			thecsas.Read(&nc,&m,filestruct,cfg,files,optCfgfile,optAll,optDebug,prefixAll)
		//case constructor == MK21
		case filestruct.Constructeur.Number == 3 :
			mk21.Read(&nc,&m,filestruct,cfg,files,optCfgfile,optAll,optDebug,prefixAll)
		}
	
}
