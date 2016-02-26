// GetProfileNumber_test.go
//Test for GetProfileNumber for THECSAS constructor

package thecsas

import "testing"

import (
	"fmt"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
)

//function for testing GetProfileNumber
func TestGetProfile(t *testing.T){
// variable for test

var ncTest lib.Nc

ncTest.Extras_s = make(map[string]string)

Profile := GetProfileNumber(&ncTest,1)
fmt.Println("Profile THERMO Number : ",Profile)

}