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
		number	int
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
		case strings.ContainsAny(cfg.Instrument.Seabird,str) : 
			result.Name = "Seabird"
			result.number = 1
		}
	}
	return result
}

