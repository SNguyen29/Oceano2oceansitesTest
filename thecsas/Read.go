//Read.go
//Func Read for thecsas software

package thecsas

import (
	"Oceano2oceansitesTest/lib"
	"Oceano2oceansitesTest/config"
	"Oceano2oceansitesTest/analyze"
	"Oceano2oceansitesTest/toml"
	"Oceano2oceansitesTest/netcdf"
	)


// read cnv files in two pass, the first to get dimensions
// second to get data
func Read(nc *lib.Nc, m *config.Map,filestruct analyze.Structfile,cfg toml.Configtoml,files []string,optCfgfile string,optAll *bool,optDebug *bool,prefixAll string) {
	
	switch{
		case filestruct.Instrument == cfg.Instrument.Type[3] :
		
			config.GetConfigTHERMO(nc,m,cfg,optCfgfile,filestruct.TypeInstrument,optAll)
			
			// first pass, return dimensions fron cnv files
			nc.Dimensions["TIME"], nc.Dimensions["DEPTH"] = firstPassTHERMO(nc,m,cfg,files)
		
			// initialize 2D data
			nc.Variables_2D = make(lib.AllData_2D)
			for i, _ := range m.Map_var {
				nc.Variables_2D.NewData_2D(i, nc.Dimensions["TIME"], nc.Dimensions["DEPTH"])
			}
		
			// second pass, read files again, extract data and fill slices
			secondPassTHERMO(nc,m,cfg,files,optDebug)
			// write ASCII file
			WriteAscii(nc,cfg,m.Map_format, m.Hdr,filestruct.Instrument,prefixAll)
		
			// write netcdf file
			//if err := nc.WriteNetcdf(); err != nil {
			//log.Fatal(err)
			//}
			netcdf.WriteNetcdf(nc,m,cfg,filestruct.Instrument,prefixAll)
			}
			
}