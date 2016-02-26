// GetProfileNumber_test.go
//Test for GetProfileNumber for MK21 constructor

package mk21

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

TestFile := "../data/T7_00001.EDF"
Profile := GetProfileNumber(&ncTest,cfg,TestFile)
fmt.Println("Profile XBT Number : ",Profile)

}