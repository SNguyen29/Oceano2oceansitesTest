// GetProfileNumber_test.go
package seabird

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
var file string
	
file,cfg = toml.InitToml(fileconfigTest)

fmt.Println("fileconfig : ",file)

ncTest.Extras_s = make(map[string]string)

TestFile := "../data/csp00101.cnv"
Profile := GetProfileNumber(&ncTest,cfg,TestFile)
fmt.Println("Profile CTD Number : ",Profile)

TestFile2 := "../data/csp00101.btl"
Profile2 := GetProfileNumber(&ncTest,cfg,TestFile2)
fmt.Println("Profile BTL Number : ",Profile2)

}