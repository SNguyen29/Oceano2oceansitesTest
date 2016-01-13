// GetProfileNumber_test.go
package main

import "testing"

import (
	"fmt"
)

//function for testing GetProfileNumber
func TestGetProfile(t *testing.T){
// variable for test

var ncTest Nc
ncTest.TestInitNC()

TestFile := "data/csp00101.cnv"
Profile := ncTest.GetProfileNumber(TestFile)
fmt.Println("Profile CTD Number : ",Profile)

TestFile2 := "data/csp00101.btl"
Profile2 := ncTest.GetProfileNumber(TestFile2)
fmt.Println("Profile BTL Number : ",Profile2)

}