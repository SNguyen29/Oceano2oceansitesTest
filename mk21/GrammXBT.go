//GrammXBT.go
//File Grammar for XBT instrument

package mk21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"regexp"
	"strings"
	"Oceano2oceansitesTest/config"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)


//function
// read .cnv files and return dimensions
func firstPassXBT(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string) (int, int) {
	
	regIsHeader := regexp.MustCompile(cfg.Mk21.Header)
	
	//variable init
	var depth float64 = 0
	var	maxDepth float64 = 0
	var cpt int = 0
	var minDepth float64 = 0
	var line int = 0
	var maxLine int = 0

	fmt.Fprintf(lib.Echo, "First pass: ")
	// loop over each files passed throw command line
	for _, file := range files {
		var Etat bool = false
		fmt.Println(file)
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
			
			if Etat {
				values := strings.Fields(str)
				// read the depth
				if depth, err = strconv.ParseFloat(values[m.Map_var["DEPTH"]], 64); err != nil {
					log.Fatal(err)
				} else {
				}
			}
			if match {
				Etat = true
			}
				
			if depth >= 0 && cpt == 0 {
				minDepth = depth
				cpt++
				}
			if depth > maxDepth {
				maxDepth = depth
				line = line + 1
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		//fmt.Fprintf(lib.Debug, "Read %s size: %d max pres: %4.f\n", file, line, maxPres)

		if line > maxLine {
			maxLine = line
		}
		// store the maximum pressure and maximum depth value per cast
		nc.Extras_f[fmt.Sprintf("MINDEPTH:%d", int(profile))] = minDepth
		nc.Extras_f[fmt.Sprintf("DEPTH:%d", int(profile))] = math.Floor(maxDepth)
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = "UNKNOW"
		nc.Variables_1D["TYPECAST"] = append(nc.Variables_1D["TYPECAST"].([]float64), 0)
		
		// reset value for next loop
		maxDepth = 0
		depth = 0
		line = 0
	}

	//fmt.Fprintf(lib.Echo, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	//fmt.Fprintf(lib.Debug, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	fmt.Fprintf(lib.Debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine+1
}

func secondPassXBT(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string,optDebug *bool) {

	regIsHeader := regexp.MustCompile(cfg.Mk21.Header)

	fmt.Fprintf(lib.Echo, "Second pass ...\n")	
	
	// initialize profile 
	var nbProfile int = 0
	

	// loop over each files passed throw command line
	for _, file := range files {
		var Etat bool = true
		var line int = 0

		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()
		// fmt.Printf("Read %s\n", file)

		profile := GetProfileNumber(nc,cfg,file)
		nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64),profile)
		scanner := bufio.NewScanner(fid)
		downcast := true
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if Etat {
				DecodeHeader(nc,cfg,str, profile,optDebug)
			} else {
				// fill map data with information contain in read line str
				DecodeData(nc,m,str, profile, file, line)
				
				if downcast {
					// fill 2D slice
					for _, key := range m.Hdr {
						if key != "PRFL" {
							lib.SetData(nc.Variables_2D[key],nbProfile,line,config.GetData(m.Data[key]))
							//fmt.Println("Line: ", line, "key: ", key, " data: ", m.Data[key])
						}
						
					}
					// exit loop if reach maximum pressure for the profile
					if m.Data["DEPTH"] == nc.Extras_f[fmt.Sprintf("DEPTH:%d", int(profile))] {
						downcast = false
					}
					
				} else {
					// store last julian day for end profile
					//nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = m.Data["ETDD"].(float64)
					//fmt.Println(presMax)
				}
				line++
			}
			if match{
				Etat = false
				}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// increment sclice index
		nbProfile += 1

		// store last julian day for end profile
		//nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = m.Data["ETDD"].(float64)
		//fmt.Println(presMax)
		value := nc.Extras_s["DATE"] +" "+ nc.Extras_s["HEURE"]
		value = lib.ConvertDate(value)
		var t = lib.NewTimeFromString("Jan 02 2006 15:04:05", value)
		v := t.Time2JulianDec()
		nc.Variables_1D["TIME"] = append(nc.Variables_1D["TIME"].([]float64),v)
	}
	
	fmt.Fprintln(lib.Debug, nc.Variables_1D["PROFILE"])
}