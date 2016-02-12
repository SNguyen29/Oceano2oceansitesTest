// AnalyzeConstructor_test.go
package analyze

import "testing"

import (
	"fmt"
	"Oceano2oceansitesTest/toml"
)

//function for testing ReadGlobal
func  TestAnalyzeConstructor(t *testing.T) {

	var cfg toml.Configtoml	
	filetoml := "../configfile/configtoml.toml"
	cfg = toml.InitToml(filetoml)
	//init variable m with empty map 

TestFile := []string{"../data/csp00101.cnv"}
	
Construct := AnalyzeConstructor(cfg,TestFile)

fmt.Println("Numero Constructeur = ",Construct)

TestFile2 := []string{"../data/csp00201.lad"}
	
Construct2 := AnalyzeConstructor(cfg,TestFile2)

fmt.Println("Numero Constructeur = ",Construct2)

}
