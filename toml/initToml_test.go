//initToml_test.go

package toml

import "testing"
import "fmt"

func TestInitToml(t *testing.T){

fileconfigTest := "../configfile/configtoml.toml"	
var cfg Configtoml
	
cfg = InitToml(fileconfigTest)

if cfg.Progversion != "0.1.0"{
		t.Error("not the same")
}else {

	fmt.Println("version :",cfg.Progversion)
}

}