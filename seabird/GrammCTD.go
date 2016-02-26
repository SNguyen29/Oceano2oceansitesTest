//GrammCTD.go
//File for with the regular expression for CTD type instrument and function for read CTD files

package seabird

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"regexp"
	"strings"
	"github.com/SNguyen29/Oceano2oceansitesTest/config"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
)


//function
// read .cnv files and return dimensions
func firstPassCTD(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string) (int, int) {
	
	regIsHeader := regexp.MustCompile(cfg.Seabird.Header)
	
	//variable init
	var	pres float64 = 0
	var depth float64 = 0
	var	maxDepth float64 = 0
	var maxPres	float64 = 0
	var maxPresAll float64 = 0
	var line int = 0
	var maxLine int = 0
	
	fmt.Fprintf(lib.Echo, "First pass: ")
	// loop over each files passed throw command line
	for _, file := range files {
		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()

		profile := GetProfileNumber(nc,cfg,file)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if !match {
				values := strings.Fields(str)
				// read the pressure
				if pres, err = strconv.ParseFloat(values[m.Map_var["PRES"]], 64); err != nil {
					log.Fatal(err)
				}
				// read the depth
				if depth, err = strconv.ParseFloat(values[m.Map_var["DEPTH"]], 64); err != nil {
					log.Fatal(err)
				} else {
					//p(math.Floor(depth))
				}
			}
			if pres > maxPres {
				maxPres = pres
				maxDepth = depth
				line = line + 1
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Fprintf(lib.Debug, "Read %s size: %d max pres: %4.f\n", file, line, maxPres)

		if line > maxLine {
			maxLine = line
		}
		// store the maximum pressure and maximum depth value per cast
		nc.Extras_f[fmt.Sprintf("PRES:%d", int(profile))] = maxPres
		nc.Extras_f[fmt.Sprintf("DEPTH:%d", int(profile))] = math.Floor(maxDepth)
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = "UNKNOW"
		if maxPres > maxPresAll {
			maxPresAll = maxPres
		}
		// reset value for next loop
		maxPres = 0
		maxDepth = 0
		pres = 0
		line = 0
	}

	fmt.Fprintf(lib.Echo, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	fmt.Fprintf(lib.Debug, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	fmt.Fprintf(lib.Debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine
}

func secondPassCTD(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string,optDebug *bool) {

	regIsHeader := regexp.MustCompile(cfg.Seabird.Header)

	fmt.Fprintf(lib.Echo, "Second pass ...\n")	
	
	// initialize profile and pressure max
	var nbProfile int = 0

	// loop over each files passed throw command line
	for _, file := range files {
		var line int = 0

		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()
		// fmt.Printf("Read %s\n", file)

		profile := GetProfileNumber(nc,cfg,file)
		scanner := bufio.NewScanner(fid)
		downcast := true
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if match {
				DecodeHeader(nc,cfg,str, profile,optDebug)
			} else {
				// fill map data with information contain in read line str
				DecodeData(nc,m,str, profile, file, line)

				if downcast {
					// fill 2D slice
					for _, key := range m.Hdr {
						if key != "PRFL" {
							//fmt.Println("Line: ", line, "key: ", key, " data: ", data[key])
							lib.SetData(nc.Variables_2D[key],nbProfile,line,config.GetData(m.Data[key]))
						}
					}
					// exit loop if reach maximum pressure for the profile
					if m.Data["PRES"] == nc.Extras_f[fmt.Sprintf("PRES:%d", int(profile))] {
						downcast = false
					}
				} else {
					// store last julian day for end profile
					nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = m.Data["ETDD"].(float64)
					//fmt.Println(presMax)
				}
				line++
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// increment sclice index
		nbProfile += 1

		// store last julian day for end profile
		nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = m.Data["ETDD"].(float64)
		//fmt.Println(presMax)
	}
	fmt.Fprintln(lib.Debug, nc.Variables_1D["PROFILE"])
}