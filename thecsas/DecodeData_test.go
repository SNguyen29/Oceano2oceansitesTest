// DecodeData_test.go
//Test for DecodeData for thecsas software
package thecsas

import "testing"

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
)


//function for testing Decodedata
func TestDecodeData(t *testing.T){
// variable for test

var ncTest lib.Nc
var m config.Map

m = config.InitMap()

ncTest.Extras_s = make(map[string]string)

fmt.Println("Debut fichier THERMO :")

TestFile := "../data/FileTestDecodeData5.TEST"

var profileTest float64 = 1

	var line int = 0

	fid, err := os.Open(TestFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		DecodeData(&ncTest,&m,str,profileTest,TestFile,line)
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
fmt.Println("Number of line : ",line)

}