//AnalyzeType_test.go
package analyze

import "testing"

import (
	"fmt"
)

//function for testing ReadGlobal
func  TestAnalyzeTypeSeabird(t *testing.T) {
	
initToml()	
TestFile := []string{"data/csp00101.cnv"}
	
Type := AnalyzeTypeSeabird(TestFile)

fmt.Println(cfg.Instrument.Type)
fmt.Println(cfg.Instrument.Decodetype)
fmt.Println("Type = "+Type)

TestFile2 := []string{"data/csp00101.btl"}
	
Type2 := AnalyzeTypeSeabird(TestFile2)

fmt.Println("Type = "+Type2)


}
