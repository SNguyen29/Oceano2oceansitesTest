//GrammTHERMO.go
//File Grammar for THERMO instrument

package thecsas

import (
	"bufio"
	"fmt"
	//"strings"
	"log"
	"os"
	"regexp"
	"Oceano2oceansitesTest/config"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/toml"
)


//function
// read .cnv files and return dimensions
func firstPassTHERMO(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string) (int, int) {
	
	regdata:= regexp.MustCompile(cfg.Thecsas.Data)
	
	//variable init
	var nbProfile int = 0
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
		Profile := GetProfileNumber(nc,nbProfile)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regdata.MatchString(str)
			if match {
				line++
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}

		if line > maxLine {
			maxLine = line
		}
		nc.Extras_f[fmt.Sprintf("DEPTH:%d", int(Profile))] = float64(maxLine)
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(Profile))] = "UNKNOW"
		nc.Variables_1D["TYPECAST"] = append(nc.Variables_1D["TYPECAST"].([]float64), 0)
		
		line = 0
		nbProfile++
	}
	fmt.Fprintf(lib.Debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine
}

func secondPassTHERMO(nc *lib.Nc,m *config.Map,cfg toml.Configtoml,files []string,optDebug *bool) {

	regdata := regexp.MustCompile(cfg.Thecsas.Data)

	fmt.Fprintf(lib.Echo, "Second pass ...\n")	
	
	// initialize profile 
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
		// increment slice index
		Profile := GetProfileNumber(nc,nbProfile)
		nc.Variables_1D["PROFILE"] = append(nc.Variables_1D["PROFILE"].([]float64),Profile)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regdata.MatchString(str)
			if match {
				// fill map data with information contain in read line str
				DecodeData(nc,m,str, Profile, file, line)
					// fill 2D slice
					for _, key := range m.Hdr {
						if key != "PRFL" {
							//fmt.Println("Line: ", line, "key: ", key, " data: ", m.Data[key])
							lib.SetData(nc.Variables_2D[key],nbProfile,line,config.GetData(m.Data[key]))
						}
					}
					
				line++
			}
			
		}
	
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		nbProfile += 1
		
		var t = lib.NewTimeFromString("Jan 02 2006 15:04:05", nc.Extras_s[fmt.Sprintf("Starttime:%d", int(Profile))])
		v := t.Time2JulianDec()
		nc.Variables_1D["TIME"] = append(nc.Variables_1D["TIME"].([]float64),v)
		
		if v, err := lib.Position2Decimal(nc.Extras_s[fmt.Sprintf("Startlongpos:%d", int(Profile))]); err == nil {
				nc.Variables_1D["LONGITUDE"] = append(nc.Variables_1D["LONGITUDE"].([]float64), v)
			}
		if v, err := lib.Position2Decimal(nc.Extras_s[fmt.Sprintf("Startlatpos:%d", int(Profile))]); err == nil {
				nc.Variables_1D["LATITUDE"] = append(nc.Variables_1D["LATITUDE"].([]float64), v)
			}
	}
	
}