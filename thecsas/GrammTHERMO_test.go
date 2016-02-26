//GrammTHERMO_test.go
//Test for Grammar THERMO

package thecsas

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
	
	TestFile := []string{"../data/20150718-114047-AT_COLCOR.COLCOR"}
	
	time,depth := firstPassTHERMO(&ncTest,&m,cfg,TestFile)
	
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
	
	profile := GetProfileNumber(&ncTest,0)
	
	TestFile := []string{"../data/20150718-114047-AT_COLCOR.COLCOR"}
	
	secondPassTHERMO(&ncTest,&m,cfg,TestFile,optDebug)
	
	fmt.Println("Time start :",ncTest.Extras_s[fmt.Sprintf("Starttime:%d", int(profile))])
	fmt.Println("Latitude start :",ncTest.Extras_s[fmt.Sprintf("Startlatpos:%d", int(profile))])
	fmt.Println("Longitude start :",ncTest.Extras_s[fmt.Sprintf("Startlongpos:%d", int(profile))])
	
	}