//lib.go

package lib

import "github.com/SNguyen29/Oceano2oceansitesTest/roscop"

type Data_2D struct {
	data [][]float64
}

type AllData_2D map[string]Data_2D

type Nc struct {
	Dimensions   map[string]int
	Variables_1D map[string]interface{}
	Variables_2D AllData_2D
	Attributes   map[string]string
	Extras_f     map[string]float64 // used to store max of profiles value
	Extras_s     map[string]string  // used to store max of profiles type
	Roscop       roscop.Roscop
}

// initialize a slice with 2 dimensions to store data
// It should be notice that this table has two dimensions allows to write
// data straightforward, it will then be flatten to write netcdf file
func (mp AllData_2D) NewData_2D(name string, width, height int) *AllData_2D {
	mt := new(Data_2D)
	mt.data = make([][]float64, width)
	for i := range mt.data {
		mt.data[i] = make([]float64, height)
		for j := range mt.data[i] {
			mt.data[i][j] = 1e36
		}
	}
	mp[name] = *mt
	return &mp
}

//function to get the data of a Data_2D 
func GetData(data Data_2D) [][]float64{
	return data.data
	}

//function to set a Data_2D with float64	
func SetData(data Data_2D,x int,y int,f float64){
	data.data[x][y] = f
	}
