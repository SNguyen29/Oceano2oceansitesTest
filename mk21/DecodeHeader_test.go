//DecodeHeader_test.go
//Test for DecodeHeader for MK21 constructor

package mk21

import "testing"
import "fmt"
import "regexp"
import "github.com/SNguyen29/Oceano2oceansitesTest/lib"
import "github.com/SNguyen29/Oceano2oceansitesTest/toml"

func TestDecodeHeader(t *testing.T){
	
	// variable for test

var cfg toml.Configtoml
var fileconfig string
var ncTest lib.Nc
var optDebug *bool
fileconfigTest := "../configfile/configtoml.toml"	

cfg = toml.InitToml(fileconfigTest)

fmt.Println("fileconfig ",fileconfig)

	ncTest.Dimensions = make(map[string]int)
	ncTest.Attributes = make(map[string]string)
	ncTest.Extras_f = make(map[string]float64)
	ncTest.Extras_s = make(map[string]string)
	ncTest.Variables_1D = make(map[string]interface{})

	// initialize map entry from nil interface to empty slice of float64
	ncTest.Variables_1D["TIME"] = []float64{}
	ncTest.Variables_1D["LATITUDE"] = []float64{}
	ncTest.Variables_1D["LONGITUDE"] = []float64{}

var profileTest float64 = 00001

//var StringTest string = "Date        = 2015/07/24"
//var StringTest string = "Start_Time  = 06:11:23"
//var StringTest string = "Latitude    = -10.9877"
var StringTest string = "Longitude     :  24 12.78809W"

temp := regexp.MustCompile(cfg.Mk21.Longitude)

fmt.Println(cfg.Mk21.Longitude)
fmt.Println(StringTest)

if temp.MatchString(StringTest){
	fmt.Println("same")
	}else{
		fmt.Println("not same")
		}
DecodeHeader(&ncTest,cfg,StringTest,profileTest,optDebug)	
	
}
