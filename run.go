package main

import (
	"fmt"
	"time"

	"github.com/angelcoto/logiss/parm"
	"github.com/angelcoto/logiss/proc"
	"github.com/angelcoto/logiss/util"
)

func run(parms parm.Parametros) error {
	inicio := time.Now()
	rangoArchivos := genRangoArchivos(parms.FechaInicial, parms.Dias)

	archivos, err := refreshCache(parms.YamlCfg.Origen, parms.YamlCfg.Destino, rangoArchivos)
	if err != nil {
		return err
	}

	if err := proc.ProcArchivos(archivos, parms); err != nil {
		return err
	}

	util.LogMensaje(fmt.Sprintf("Proceso finalizado en %v", time.Since(inicio)))

	return nil

}

// runOnce ejecuta el proceso una vez.
// (run.go)
func runOnce(parms parm.Parametros) {
	if err := run(parms); err != nil {
		util.LogError(err)
	}
}

// runForever ejecuta el proceso de forma infinita, teniendo un tiempo de espera
// de n minutos entre cada ejecución.
func runForever(parms parm.Parametros) {
	var err error
	for {
		if err = run(parms); err != nil {
			util.LogError(err)
		}

		time.Sleep(time.Minute * time.Duration(parms.YamlCfg.Espera))
	}
}
