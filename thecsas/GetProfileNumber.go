// GetProfileNumber.go
//Function for get the profil number of a data file for THECSAS Constructor
package thecsas

import (
	"fmt"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
)

func GetProfileNumber(nc *lib.Nc,profile int) float64 {
	
		var fprofile = float64(profile)
		
		
		nc.Extras_s[fmt.Sprintf("PRFL_NAME:%d", profile)] = fmt.Sprintf("%d",profile)

		return fprofile

}