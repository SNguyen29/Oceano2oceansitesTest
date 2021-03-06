//GrammXBT_test.go
//Test for Grammar XBT

package mk21

import "testing"
import "fmt"
import "github.com/SNguyen29/Oceano2oceansitesTest/toml"
import "github.com/SNguyen29/Oceano2oceansitesTest/config"
import "github.com/SNguyen29/Oceano2oceansitesTest/lib"

func TestFirstPass(t *testing.T){
	
	var ncTest lib.Nc
	var m config.Map
	var cfg toml.Configtoml	
	filetoml := "../configfile/configtoml.toml"
	cfg = toml.InitToml(filetoml)

	m = config.InitMap()
	
	ncTest.Extras_s = make(map[string]string)
	ncTest.Extras_f = make(map[string]float64)
	
	TestFile := []string{"../data/T7_00001.EDF"}
	
	time,depth := firstPassXBT(&ncTest,&m,cfg,TestFile)
	
	fmt.Println("Time : ",time)
	fmt.Println("Depth : ",depth)
	
	}
	
func TestSecondPass(t *testing.T){
	
	var ncTest lib.Nc
	var m config.Map
	var cfg toml.Configtoml	
	var optDebug *bool
	filetoml := "../configfile/configtoml.toml"
	cfg = toml.InitToml(filetoml)

	m = config.InitMap()
	
	ncTest.Extras_s = make(map[string]string)
	ncTest.Extras_f = make(map[string]float64)
	ncTest.Variables_1D = make(map[string]interface{})
	ncTest.Variables_1D["TIME"] = []float64{}
	ncTest.Variables_1D["LATITUDE"] = []float64{}
	ncTest.Variables_1D["LONGITUDE"] = []float64{}
	ncTest.Variables_1D["PROFILE"] = []float64{}
	
	TestFile := []string{"../data/T7_00001.EDF"}
	
	secondPassXBT(&ncTest,&m,cfg,TestFile,optDebug)
	
	fmt.Println("Time :",ncTest.Variables_1D["TIME"])
	fmt.Println("Latitude :",ncTest.Variables_1D["LATITUDE"])
	fmt.Println("Longitude :",ncTest.Variables_1D["LONGITUDE"])
	
	}