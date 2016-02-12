// GetProfileNumber_test.go
//Test for GetProfileNumber for IFM constructor

package ifm

import "testing"

import (
	"fmt"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)

//function for testing GetProfileNumber
func TestGetProfile(t *testing.T){
// variable for test

fileconfigTest := "../configfile/configtoml.toml"
var cfg toml.Configtoml
var ncTest lib.Nc
	
cfg = toml.InitToml(fileconfigTest)

ncTest.Extras_s = make(map[string]string)

TestFile := "../data/csp00201.lad"
Profile := GetProfileNumber(&ncTest,cfg,TestFile)
fmt.Println("Profile LADCP Number : ",Profile)

}