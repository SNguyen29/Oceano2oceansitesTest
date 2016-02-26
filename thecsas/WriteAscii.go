// writeAscii.go
//func writes Ascii file for thecsas software

package thecsas

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"

)

const (
	codeForProfile = -1
)

func WriteAscii(nc *lib.Nc,cfg toml.Configtoml,map_format map[string]string, hdr []string, inst string,prefixAll string) {
	// define 2 files, profiles header and data
	var asciiFilename string
	var lat2 string
	var latdecimal float64
	var longdecimal float64
	var long2 string

	// build filenames
	str := nc.Attributes["cycle_mesure"]
	str = strings.Replace(str, "\r", "", -1)
	headerFilename := fmt.Sprintf(cfg.Outputfile+"%s."+inst, strings.ToLower(str))
	asciiFilename = fmt.Sprintf(cfg.Outputfile+"%s%s_"+inst, strings.ToLower(str), prefixAll)
	//fmt.Println(headerFilename)
	//fmt.Println(asciiFilename)

	// open header file for writing result
	fid_hdr, err := os.Create(headerFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fid_hdr.Close()

	// use buffered mode for writing
	fbuf_hdr := bufio.NewWriter(fid_hdr)

	// open ASCII file for writing result
	fid_ascii, err := os.Create(asciiFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fid_ascii.Close()

	// use buffered mode for writing
	fbuf_ascii := bufio.NewWriter(fid_ascii)

	// write header to string
	str = fmt.Sprintf("%s  %s  %s  %s %s  %s\n",
		nc.Attributes["cycle_mesure"],
		nc.Attributes["plateforme"],
		nc.Attributes["institute"],
		nc.Attributes["type_instrument"],
		nc.Attributes["instrument_number"],
		nc.Attributes["pi"])

	// write first line header on header file and ascii file
	fmt.Fprintf(fbuf_hdr, str)
	fmt.Fprintf(fbuf_ascii, str)

	// display on screen
	fmt.Printf("%s", str)

	// write physical parameters in second line
	str = ""
	for _, key := range hdr {
		fmt.Fprintf(fbuf_ascii, "%s   ", key)
		fmt.Fprintf(lib.Debug, "%s   ", key)
	}
	// append new line
	//fmt.Fprintln(fbuf_ascii, "\n")

	// write second line header on ascii file
	fmt.Fprintln(fbuf_ascii, str)

	// display on screen
	fmt.Printf("%s", str)

	// get data (slices) from nc struct
	len_1D := nc.Dimensions["TIME"]
	len_2D := nc.Dimensions["DEPTH"]
	time := nc.Variables_1D["TIME"].([]float64)
	lat := nc.Variables_1D["LATITUDE"].([]float64)
	lon := nc.Variables_1D["LONGITUDE"].([]float64)
	profile := nc.Variables_1D["PROFILE"].([]float64)

	// loop over each profile
	for x := 0; x < len_1D; x++ {
		str = ""
		// write profile informations to ASCII data file with DEPTH = -1
		t1 := lib.NewTimeFromJulian(time[x])
		var t = lib.NewTimeFromString("Jan 02 2006 15:04:05", nc.Extras_s[fmt.Sprintf("Stopttime:%d", int(profile[x]))])
		t2 := lib.NewTimeFromJulian(t.Time2JulianDec())
		// TODOS: adapt profile format to stationPrefixLength
		fmt.Fprintf(fbuf_ascii, "%05.0f %4d %f %f %f %s",
			profile[x],
			codeForProfile,
			t1.JulianDayOfYear(),
			lat[x],
			lon[x],
			t1.Format("20060102150405"))

		if longdecimal, err := lib.Position2Decimal(nc.Extras_s[fmt.Sprintf("Stoplongpos:%d", int(profile[x]))]); err == nil {
				long2 = lib.DecimalPosition2String(longdecimal,0)
			}
		if latdecimal, err := lib.Position2Decimal(nc.Extras_s[fmt.Sprintf("Stoplatpos:%d", int(profile[x]))]); err == nil {
				lat2 = lib.DecimalPosition2String(latdecimal,0)
			}		
		
		// write profile informations to header file
		str = fmt.Sprintf("%05.0f %s %s %s %s %s %s %s %s\n",
			profile[x],
			t1.Format("02/01/2006 15:04:05"),
			t2.Format("02/01/2006 15:04:05"),
			lib.DecimalPosition2String(lat[x], 0),
			lib.DecimalPosition2String(lon[x], 0),
			lat2,
			long2,
			nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile[x]))],	
			nc.Extras_s[fmt.Sprintf("PRFL_NAME:%d", int(profile[x]))]+cfg.Thermo.CruisePrefix)

		// write profile information to header file
		fmt.Fprintf(fbuf_hdr, str)

		// display on screen
		fmt.Printf("%s", str)

		// fill last header columns with 1e36
		for i := 0; i < len(hdr)-6; i++ {
			fmt.Fprintf(fbuf_ascii, " %g", 1e36)
		}
		fmt.Fprintln(fbuf_ascii) // add newline

		// loop over each level
		for y := 0; y < len_2D; y++ {
			// goto next profile when max depth reach
			if lib.GetData(nc.Variables_2D["LAT"])[x][y] ==
				latdecimal && lib.GetData(nc.Variables_2D["LONG"])[x][y] == longdecimal {
				continue
			}
			fmt.Fprintf(fbuf_ascii, "%05.0f ", profile[x])
			// loop over each physical parameter (key) in the rigth order
			for _, key := range hdr {
				// if key not in map, goto next key
				if _, ok := nc.Variables_2D[key]; ok {
					// fill 2D slice
					data := lib.GetData(nc.Variables_2D[key])[x][y]
					// print data with it's format, change format for FillValue
					if data == 1e36 {
						fmt.Fprintf(fbuf_ascii, "%g ", data)
					} else {
						if strings.ContainsAny(map_format[key],"lf"){
								res := strings.Split(map_format[key],"l")
								map_format[key] = strings.Join(res,"")
							}
						fmt.Fprintf(fbuf_ascii, map_format[key], data)
					}
				}
			}
			fmt.Fprintf(fbuf_ascii, "\n")

		}
		fbuf_ascii.Flush()
		fbuf_hdr.Flush()
	}
}
