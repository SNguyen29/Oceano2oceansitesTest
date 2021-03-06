// GetProfileNumber_test.go
//Test for GetProfileNumber for Seabird constructor

package seabird

import "testing"

import (
	"fmt"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
)

//function for testing GetProfileNumber
func TestGetProfile(t *testing.T){
// variable for test

fileconfigTest := "../configfile/configtoml.toml"
var cfg toml.Configtoml
var ncTest lib.Nc
	
cfg = toml.InitToml(fileconfigTest)

ncTest.Extras_s = make(map[string]string)

TestFile := "../data/csp00201.cnv"
Profile := GetProfileNumber(&ncTest,cfg,TestFile)
fmt.Println("Profile CTD Number : ",Profile)

}