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
		// convert types from code_roscop structure to native netcdf types
		var netcdfType netcdf.Type
		
		/*//remove LATITUDE and LONGITUDE from the key list if thermo file
		if ncType=="THERMO"{
			if key == "LATITUDE" {
				continue
			}
			if key == "LONGITUDE" {
				continue
			}
		}*/

		pa := roscop.GetAttributesStringValue(key, "types")
		switch pa {
		case "int32":
			netcdfType = netcdf.INT
		case "float32":
			netcdfType = netcdf.FLOAT
		case "float64":
			netcdfType = netcdf.DOUBLE
		default:
			log.Fatal(fmt.Sprintf("Error: key: %s, Value: [%s], check roscop file\n", key, pa)) // wrong type, check code_roscop file
		}
		// add variables
		v, err := ds.AddVar(key, netcdfType, dim_1D)
		if err != nil {
			log.Fatal(err)
		}
		map_1D[key] = v

		// define variable attributes with the right type
		// for an physical parameter, get a slice of attributes name
		for _, name := range roscop.GetAttributes(key) {
			// for each attribute, get the value
			value := roscop.GetAttributesValue(key, name)
			// add new attribute to the variable v
			a := v.Attr(name)
			// value is an interface{}, need type assertion
			switch value.(type) {
			case string:
				err = a.WriteBytes([]byte(value.(string)))
			case int32:
				err = a.WriteInt32s([]int32{value.(int32)})
			case float32:
				err = a.WriteFloat32s([]float32{value.(float32)})
			case float64:
				err = a.WriteFloat64s([]float64{value.(float64)})
			default:
				log.Fatal(fmt.Sprintf("netcdf: create 1D attribute error, %v=%v:%v (%T)",
					key, name, value, value)) // wrong type, check code_roscop file
			}
			if err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, v, v))
			}
			fmt.Fprintf(debug, "%s: %s=%v (%T)\n", key, name, value, value)
		}
	}

	// Add the variable to the dataset that will store our data
	map_2D := make(map[string]netcdf.Var)

	// use the order list gave by split or splitAll (config file) because
	// the iteration order is not specified and is not guaranteed to be
	// the same from one iteration to the next in golang
	// for key, _ := range nc.Variables_2D {
	for _, key := range config.GetPhysicalParametersList(m) {
		// remove PRFL from the key list
		if key == "PRFL" {
			continue
		}
		
		/*if key == "SSJT" {
			continue
		}*/
	
		// convert types from code_roscop structure to native netcdf types
		var netcdfType netcdf.Type

		pa := roscop.GetAttributesStringValue(key, "types")
		switch pa {
		case "int32":
			netcdfType = netcdf.INT
		case "float32":
			netcdfType = netcdf.FLOAT
		case "float64":
			netcdfType = netcdf.DOUBLE
		default:
			log.Fatal(fmt.Sprintf("Error: key: %s, Value: [%s], check roscop file\n", key, pa)) // wrong type, check code_roscop file
		}
		v, err := ds.AddVar(key, netcdfType, dim_2D)
		if err != nil {
			log.Fatal(err) 
		}
		map_2D[key] = v

		// define variable attributes with the right type
		// for an physical parameter, get a slice of attributes name
		for _, name := range roscop.GetAttributes(key) {
			// for each attribute, get the value
			value := roscop.GetAttributesValue(key, name)
			// add new attribute to the variable v
			a := v.Attr(name)
			// value is an interface{}, need type assertion
			switch value.(type) {
			case string:
				err = a.WriteBytes([]byte(value.(string)))
			case int32:
				err = a.WriteInt32s([]int32{value.(int32)})
			case float32:
				err = a.WriteFloat32s([]float32{value.(float32)})
			case float64:
				err = a.WriteFloat64s([]float64{value.(float64)})
			default:
				log.Fatal(fmt.Sprintf("netcdf: create 1D attribute error: key: %s, Value: [%s], \n", key, value))
			}
			if err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, v, v))
			}
			fmt.Fprintf(debug, "%s: %s=%v (%T)\n", key, name, value, value)
		}
	}

	// defines global attributes
	for key, value := range nc.Attributes {
		a := ds.Attr(key)
		err = a.WriteBytes([]byte(value))
		if err != nil {
			log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, value, value))
		}
	}

	// leave define mode in NetCDF3
	ds.EndDef()

	// Create the data with the above dimensions and type,
	// write them to the file.
	for key, value := range nc.Variables_1D {
		// convert types from code_roscop structure to native netcdf types
		switch roscop.GetAttributesStringValue(key, "types") {
		case "int32":
			length := len(value.([]float64))
			v := make([]int32, length)
			fmt.Fprintf(echo, "writing %s: %d\n", key, len(v))
			fmt.Fprintf(debug, "writing %s: %d\n", key, len(v))
			for i := 0; i < length; i++ {
				v[i] = int32(value.([]float64)[i])
			}
			if err := map_1D[key].WriteInt32s(v); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, v, v))
			}
		case "float32":
			length := len(value.([]float64))
			v := make([]float32, length)
			fmt.Fprintf(echo, "writing %s: %d\n", key, len(v))
			fmt.Fprintf(debug, "writing %s: %d\n", key, len(v))
			for i := 0; i < length; i++ {
				v[i] = float32(value.([]float64)[i])
			}
			if err := map_1D[key].WriteFloat32s(v); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, v, v))
			}
		case "float64":
			length := len(value.([]float64))
			v := make([]float64, length)
			fmt.Fprintf(echo, "writing %s: %d\n", key, len(v))
			fmt.Fprintf(debug, "writing %s: %d\n", key, len(v))
			for i := 0; i < length; i++ {
				v[i] = float64(value.([]float64)[i])
			}
			if err := map_1D[key].WriteFloat64s(v); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: %v (%T)", err, key, v, v))
			}
		default:
			log.Fatal(fmt.Sprintf("%s, %v", err, key))
		}

	}

	// write data 2D (value.data) to netcdf variables
	// for key, value := range nc.Variables_2D {
	for _, key := range config.GetPhysicalParametersList(m) {
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
		switch roscop.GetAttributesStringValue(key, "types") {
		case "int32":
			v := make([]int32, ht*wd)
			for i := 0; i < ht*wd; i++ {
				v[i] = int32(gopher[i])
			}
			if err := map_2D[key].WriteInt32s(v); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: (%T)", err, key, v))
			}
		case "float32":
			v := make([]float32, ht*wd)
			for i := 0; i < ht*wd; i++ {
				v[i] = float32(gopher[i])
			}
			if err := map_2D[key].WriteFloat32s(v); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: (%T)", err, key, v))
			}
		case "float64":
			if err := map_2D[key].WriteFloat64s(gopher); err != nil {
				log.Fatal(fmt.Sprintf("%s, %v: (%T)", err, key, gopher))
			}
		default:
			log.Fatal(fmt.Sprintf("%s, %v", err, key))
		}
	}
	fmt.Fprintf(echo, "writing %s done ...\n", filename)
}