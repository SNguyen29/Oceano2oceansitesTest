//GrammBTL.go
//File for with the regular expression for BTL type instrument and function for read BTL files

package seabird

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"github.com/SNguyen29/Oceano2oceansitesTest/lib"
	"github.com/SNguyen29/Oceano2oceansitesTest/config"
	"github.com/SNguyen29/Oceano2oceansitesTest/toml"
)

// read .btl files and return dimensions
func firstPassBTL(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string) (int, int) {

	var line int = 0
	var maxLine int = 0
	var bottle float64 = 0
	var minBottle float64 = 0
	var maxBottle float64 = 0
	var cpt int = 0
	
	regheader := regexp.MustCompile(cfg.Seabird.Header)
	regheaderBTL := regexp.MustCompile(cfg.Seabird.HeaderBTL)
	regheaderBTL2 := regexp.MustCompile(cfg.Seabird.HeaderBTL2)

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
			match := regheader.MatchString(str)
			if !match {
				match := regheaderBTL.MatchString(str)
				match1 := regheaderBTL2.MatchString(str)
				if !match || !match1 {
					values := strings.Fields(str)
					if bottle, err = strconv.ParseFloat(values[m.Map_var["BOTL"]], 64); err != nil {
						log.Fatal(err)
						}
				}
			}
			if bottle >= 0 && cpt == 0 {
				minBottle = bottle
				cpt++
				}
			if bottle > maxBottle {
				line = line + 1
				maxBottle = bottle
				
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Fprintf(lib.Debug, "Read %s size: %d max pres: %4.f\n", file, line, maxBottle)

		if line > maxLine {
			maxLine = line
		}
		// store the maximum pressure value
		nc.Extras_f[fmt.Sprintf("MinBOTL:%d", int(profile))] = minBottle+1
		nc.Extras_f[fmt.Sprintf("BOTL:%d", int(profile))] = maxBottle
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = "UNKNOW"
		// reset value for next loop
		maxBottle = 0
		bottle = 0
		line = 0
	}
	
	fmt.Fprintf(lib.Echo, "First pass: %d files read, maximum bottle found: %4.0f db\n", len(files), maxBottle)
	fmt.Fprintf(lib.Debug, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxBottle)
	fmt.Fprintf(lib.Debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine
}

// read .cnv files and extract data
func secondPassBTL(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string,optDebug *bool) {

	regheader := regexp.MustCompile(cfg.Seabird.Header)
	regheaderBTL := regexp.MustCompile(cfg.Seabird.HeaderBTL)
	regheaderBTL2 := regexp.MustCompile(cfg.Seabird.HeaderBTL2)
	
	// initialize profile 
	var nbProfile int = 0

	fmt.Fprintf(lib.Echo, "Second pass ...\n")
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
		fmt.Println("start file : ",profile)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regheader.MatchString(str)
			if match {
				DecodeHeader(nc,cfg,str, profile,optDebug)
			} else {
				match := regheaderBTL.MatchString(str)
				match1 := regheaderBTL2.MatchString(str)
				if !match || !match1{
						fmt.Println("decodedata  :",line)
						DecodeData(nc,m,str, profile, file, line)
							// fill 2D slice
							for _, key := range m.Hdr {
								if key != "PRFL" {
									//fmt.Println("Line: ", line, "key: ", key, " data: ", m.Data[key])
									lib.SetData(nc.Variables_2D[key],nbProfile,line,config.GetData(m.Data[key]))
								}
								
							}
						line = line+1			
					}		
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		
	}
}