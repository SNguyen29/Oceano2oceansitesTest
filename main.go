
package main

import (
	"fmt"
	"log"
	//"os"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
	"Oceano2oceansitesTest/toml"
	"io"
	"io/ioutil"
	"Oceano2oceansitesTest/analyze"
	"Oceano2oceansitesTest/seabird"
	
) 

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

var fileconfig string

// main body
func main() {
	
	//init variable cfg with config file in TOML
	fileconfig,cfg = toml.InitToml()
	//init variable m with empty map 
	m = config.InitMap()
	
	var files []string
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	files, optCfgfile := GetOptions()
	
	//analyse the file to know contructor, instrument and instrument type
	filestruct = analyze.AnalyzeFile(cfg,files)
	
	switch{
		
		//case constructor == seabird
		case filestruct.Constructeur.Number == 0 :
			seabird.ReadSeabird(&nc,&m,filestruct,cfg,files,optCfgfile,optAll,optDebug,prefixAll)
		}
	
}
