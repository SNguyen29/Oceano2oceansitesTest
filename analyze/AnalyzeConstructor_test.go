// AnalyzeConstructor_test.go
package analyze

import "testing"

import (
	"fmt"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
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

TestFile2 := []string{"../data/csp00101.btl"}
	
Construct2 := AnalyzeConstructor(cfg,TestFile2)

fmt.Println("Numero Constructeur = ",Construct2)

TestFile3 := []string{"../data/csp00201.lad"}
	
Construct3 := AnalyzeConstructor(cfg,TestFile3)

fmt.Println("Numero Constructeur = ",Construct3)

TestFile4 := []string{"../data/20150718-114047-AT_COLCOR.COLCOR"}
	
Construct4 := AnalyzeConstructor(cfg,TestFile4)

fmt.Println("Numero Constructeur = ",Construct4)

TestFile5 := []string{"../data/T7_00001.EDF"}
	
Construct5 := AnalyzeConstructor(cfg,TestFile5)

fmt.Println("Numero Constructeur = ",Construct5)

}
