// DecodeDataSeabird_test.go
package seabird

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

fmt.Println("Debut fichier CTD :")

TestFile := "../data/FileTestDecodeData.cnv"

var profileTest float64 = 00101

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
fmt.Println("Debut fichier BTL :")
TestFile2 := "../data/FileTestDecodeData2.btl"

var profileTest2 float64 = 00101

	var line2 int = 0

	fid2, err2 := os.Open(TestFile2)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer fid2.Close()
	
	scanner2 := bufio.NewScanner(fid2)
	for scanner2.Scan() {
		str2 := scanner2.Text()
		DecodeData(&ncTest,&m,str2,profileTest2,TestFile2,line2)
		line2++
	}
	if err2 := scanner2.Err(); err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Number of line : ",line2)
}