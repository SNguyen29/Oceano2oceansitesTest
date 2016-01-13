// AnalyzeConstructor_test.go
package analyze

import "testing"

import (
	"fmt"
)

//function for testing ReadGlobal
func  TestAnalyzeConstructor(t *testing.T) {

TestFile := []string{"data/csp00101.cnv"}
	
Construct := AnalyzeConstructor(TestFile)

fmt.Println("Numero Constructeur = ",Construct)

}
