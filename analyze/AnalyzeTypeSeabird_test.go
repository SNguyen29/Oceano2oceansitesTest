//AnalyzeType_test.go
package analyze

import "testing"

import (
	"fmt"
	"Oceano2oceansitesTest/toml"
	
)

//function for testing ReadGlobal
func  TestAnalyzeType(t *testing.T) {
	
	var cfg toml.Configtoml	
	filetoml := "../configfile/configtoml.toml"
	cfg = toml.InitToml(filetoml)	
	
TestFile := []string{"../data/csp00101.cnv"}
	
Type := AnalyzeType(cfg,TestFile)

fmt.Println(cfg.Instrument.Type)
fmt.Println(cfg.Instrument.Decodetype)
fmt.Println("Type = "+Type)

TestFile2 := []string{"../data/csp00101.btl"}
	
Type2 := AnalyzeType(cfg,TestFile2)

fmt.Println("Type = "+Type2)

TestFile3 := []string{"../data/csp00201.lad"}
	
Type3 := AnalyzeType(cfg,TestFile3)

fmt.Println("Type = "+Type3)


}
