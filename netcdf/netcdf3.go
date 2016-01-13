package netcdf

import (
	"fmt"
	"github.com/fhs/go-netcdf/netcdf"
	"log"
	"strings"
	"io"
	"io/ioutil"
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
	"Oceano2oceansitesTest/toml"
	r "Oceano2oceansitesTest/roscop"
)

// use for debug mode
var debug io.Writer = ioutil.Discard
// use for echo mode
var echo io.Writer = ioutil.Discard

// creates the NetCDF file following nc structure.
//func WriteNetcdf(any interface{}) error {
func WriteNetcdf(nc *lib.Nc,m *config.Map, cfg toml.Configtoml,ncType string,prefixAll string) {

	// build filename
	filename := fmt.Sprintf(cfg.Outputfile+"OS_%s%s_%s.nc",
		strings.ToUpper(nc.Attributes["cycle_mesure"]),
		strings.ToUpper(prefixAll),
		ncType)
	//fmt.Println(filename)

	// get roscop definition file for variables attributes
	var roscop = nc.Roscop
	//	for k, v := range roscop {
	//		fmt.Printf("%s: ", k)
	//		fmt.Println(v)
	//	}
	//	os.Exit(0)

	fmt.Fprintf(echo, "writing netCDF: %s\n", filename)

	// get variables_1D size
	len_1D := nc.Dimensions["TIME"]
	len_2D := nc.Dimensions["DEPTH"]

	// Create a new NetCDF 3 file. The dataset is returned.
	ds, err := netcdf.CreateFile(filename, netcdf.CLOBBER)
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	// Add the dimensions for our data to the dataset
	dim_1D := make([]netcdf.Dim, 1)
	dim_2D := make([]netcdf.Dim, 2)

	// dimensions for ROSCOP paremeters as DEPTH, PRES, TEMP, PSAL, etc
	dim_2D[0], err = ds.AddDim("TIME", uint64(len_1D))
	if err != nil {
		log.Fatal(err)
	}
	dim_2D[1], err = ds.AddDim("DEPTH", uint64(len_2D))
	if err != nil {
		log.Fatal(err)
	}
	// dimension for PROFILE, LATITUDE, LONGITUDE and BATH
	dim_1D[0] = dim_2D[0]

	// Add the variable to the dataset that will store our data
	map_1D := make(map[string]netcdf.Var)
	for key, _ := range nc.Variables_1D {
		v, err := ds.AddVar(key, netcdf.DOUBLE, dim_1D)
		if err != nil {
			log.Fatal(err)
		}
		map_1D[key] = v
		
		//initialise struct roscop with netcdf variable in reflect ways without knowing the form of the roscop struct
		r.Reflectroscop(roscop[key],map_1D[key])
		
	}

	map_2D := make(map[string]netcdf.Var)

	// use the order list gave by split or splitAll (config file) because
	// the iteration order is not specified and is not guaranteed to be
	// the same from one iteration to the next in golang
	// for key, _ := range nc.Variables_2D {
	for _, key := range m.Hdr {
		// remove PRFL from the key list
		if key == "PRFL" {
			continue
		}
		v, err := ds.AddVar(key, netcdf.DOUBLE, dim_2D)
		if err != nil {
			log.Fatal(err)
		}
		map_2D[key] = v
		
		//initialise struct roscop with netcdf variable in reflect ways without knowing the form of the roscop struct
		r.Reflectroscop(roscop[key],map_2D[key])
	
	}

	// defines global attributes
	for key, value := range nc.Attributes {
		a := ds.Attr(key)
		a.WriteBytes([]byte(value))
	}

	// leave define mode in NetCDF3
	ds.EndDef()

	// Create the data with the above dimensions and write it to the file.
	for key, value := range nc.Variables_1D {

		v := value.([]float64)
		fmt.Fprintf(echo, "writing %s: %d\n", key, len(v))
		err = map_1D[key].WriteFloat64s(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	// write data 2D (value.data) to netcdf variables
	// for key, value := range nc.Variables_2D {
	for _, key := range m.Hdr {
		// remove PRFL from the key list
		if key == "PRFL" {
			continue
		}
		value := nc.Variables_2D[key]
		i := 0
		ht := len(lib.GetData(value))
		wd := len(lib.GetData(value)[0])
		fmt.Fprintf(echo, "writing %s: %d x %d\n", key, ht, wd)
		fmt.Fprintf(debug, "writing %s: %d x %d\n", key, ht, wd)
		// Write<type> netcdf methods need []<type>, [][]data will be flatten
		gopher := make([]float64, ht*wd)
		for x := 0; x < ht; x++ {
			for y := 0; y < wd; y++ {
				gopher[i] = lib.GetData(value)[x][y]
				i++
			}
		}
		err = map_2D[key].WriteFloat64s(gopher)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprintf(echo, "writing %s done ...\n", filename)
	//return nil
}
