//roscop_test.go

package roscop

import "testing"
import "fmt"

func TestRoscop(t *testing.T) {

	// initialize new roscop object from file
	roscop := NewRoscop("code_roscop.csv")

// loop over each physicalParameter and display attribute values
	for _, physicalParameter := range roscop.GetPhysicalParameters() {
	fmt.Println("%s => ", physicalParameter)

		for _, attributeName := range roscop.GetAttributes(physicalParameter) {
			value := roscop.GetAttributesValue(physicalParameter, attributeName)
			fmt.Println("%s: %v (%T), ", attributeName, value, value)
		}
	
	}
}