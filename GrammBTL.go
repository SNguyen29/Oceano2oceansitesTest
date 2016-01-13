//GrammBTL.go
//File for with the regular expression for BTL type instrument and function for read BTL files

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
)


//regular expressions
var regIsHour = regexp.MustCompile(`^\s+(\d+:\d+:\d+)`)
var regIsDate = regexp.MustCompile(`^\s+\d+\s+(\w{3}\s+\d{2}\s+\d{4})`)
var regIsMontDayYear = regexp.MustCompile(`^\s+\d+\s+(\w{3})\s+(\d{2})\s+(\d{4})`)
var regIsHeaderBtl = regexp.MustCompile(`^[*#]|^\s+\w+`)

// read .btl files and return dimensions
func firstPassBTL(nc *lib.Nc,m *config.Map,files []string) (int, int) {

	var line int = 0
	var maxLine int = 0
	var bottle float64 = 0
	var maxBottle float64 = 0
	var maxBottleAll float64 = 0

	fmt.Fprintf(echo, "First pass: ")
	// loop over each files passed throw command line
	for _, file := range files {
		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()

		profile := GetProfileNumber(nc,file)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeaderBtl.MatchString(str)
			if !match {
				p(str)
				values := strings.Fields(str)
				p("BOTL", m.Map_var["BOTL"])
				p(values[m.Map_var["BOTL"]])
				if bottle, err = strconv.ParseFloat(values[m.Map_var["BOTL"]], 64); err != nil {
					log.Fatal(err)
				}
				fmt.Fprintln(debug, values)

			}
			if bottle > maxBottle {
				maxBottle = bottle
				line = line + 1
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Fprintf(debug, "Read %s size: %d max pres: %4.f\n", file, line, maxBottle)

		if line > maxLine {
			maxLine = line
		}
		// store the maximum pressure value
		nc.Extras_f[fmt.Sprintf("BOTL:%d", int(profile))] = maxBottle
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = "n/a"
		if maxBottle > maxBottleAll {
			maxBottleAll = maxBottle
		}
		// reset value for next loop
		maxBottle = 0
		bottle = 0
		line = 0
	}
	
	fmt.Fprintf(echo, "First pass: %d files read, maximum bottle found: %4.0f db\n", len(files), maxBottle)
	fmt.Fprintf(debug, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxBottle)
	fmt.Fprintf(debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine
}

// read .cnv files and extract data
func secondPassBTL(nc *lib.Nc,m *config.Map,files []string) {

	regIsHeader := regexp.MustCompile(cfg.Seabird.Header)
	
	var month, day, year string

	fmt.Fprintf(echo, "Second pass ...\n")
	// loop over each files passed throw command line
	for _, file := range files {
		var line int = 0

		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()
		// fmt.Printf("Read %s\n", file)

		profile := GetProfileNumber(nc,file)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if match {
				DecodeHeaderSeabird(nc,str, profile)
			} else {
				match = regIsMontDayYear.MatchString(str)
				if match {
					res := regIsMontDayYear.FindStringSubmatch(str)
					month, day, year = res[1], res[2], res[3]
					//date = res[1]
					//f("Date -> %s/%s/%s\n", month, day, year)
				}
				match = regIsHour.MatchString(str)
				if match {
					res := regIsHour.FindStringSubmatch(str)
					time := res[1]
					//f("Time -> %s\n", time)
					// create new Time object, see tools.go
					var t = NewTimeFromString("Jan 02 2006 15:04:05 UTC",
						fmt.Sprintf("%s %s %s %s", month, day, year, time))
					//					v := t.Time2JulianDec()
					//					t1 := NewTimeFromJulian(v)
					y, _ := strconv.ParseFloat(year, 64)
					t2 := NewTimeFromJulianDay(y, t)
					nc.Variables_1D["TIME"] = append(nc.Variables_1D["TIME"].([]float64),
						t2.JulianDayOfYear())
					//p(t2.JulianDayOfYear())
				}
			}
			line++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}