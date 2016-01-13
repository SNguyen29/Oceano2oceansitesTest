//Reflect.go
//Function for who use the struct RoscopAtrribute to initialise variable nectcdf
package roscop

import (
		"reflect"
		"github.com/fhs/go-netcdf/netcdf"
		"strings"
		)

func Reflectroscop(r RoscopAttribute,m netcdf.Var){
		
		val := reflect.ValueOf(r)

		for j:=0;j<val.NumField();j++{
			a := m.Attr(val.Type().Field(j).Name)
			switch{
				case strings.EqualFold("string",val.Type().Field(j).Type.String()):
					a.WriteBytes([]byte(val.Field(j).String()))
				case strings.EqualFold("float64",val.Type().Field(j).Type.String()):
					a.WriteFloat64s([]float64{val.Field(j).Float()})
				}		
		}
		
	
}