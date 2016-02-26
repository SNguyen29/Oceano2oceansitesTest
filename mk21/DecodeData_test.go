// DecodeData_test.go
//Test for DecodeData for MK21 constructor
package mk21

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

fmt.Println("Debut fichier XBT :")

TestFile := "../data/FileTestDecodeData4.TEST"

var profileTest float64 = 00001

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