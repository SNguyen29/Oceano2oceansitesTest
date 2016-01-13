//AnalyzeConstructor.go
//Analyze the constructor of the data files
package analyze

import (
	"bufio"
	"log"
	"os"
	"strings"
	"Oceano2oceansitesTest/toml"
)

// define constante for constructor type
type Constructor struct{
		Name	string
		Number	int
	}



// read all cnv files and return dimensions
func AnalyzeConstructor(cfg toml.Configtoml,files []string) Constructor {

	var result Constructor
	// open first file
	fid, err := os.Open(files[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		
		switch {
		case strings.ContainsAny(cfg.Instrument.Constructor[0],str) : 
			result.Name = cfg.Instrument.Constructor[0]
			result.Number = 0
		}
	}
	return result
}

