package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/angelcoto/logiss/proc"
	"github.com/angelcoto/logiss/util"
)

func run(parms parametros) error {
	inicio := time.Now()
	rangoArchivos := genRangoArchivos(parms.fechaInicial, parms.dias)

	archivos, err := refreshCache(parms.yamlCfg.origen, parms.yamlCfg.destino, rangoArchivos)
	if err != nil {
		util.LogError(err)
		return err
	}

	if err := proc.ProcArchivos(archivos, parms.yamlCfg.csvPath, parms.yamlCfg.exclUrsNull); err != nil {
		util.LogError(err)
		return err
	}

	util.LogMensaje(fmt.Sprintf("Proceso finalizado en %v", time.Since(inicio)))

	return nil

}

// runOnce ejecuta el proceso una vez.
// (run.go)
func runOnce(parms parametros) error {
	err := run(parms)
	return err
}

// runForever ejecuta el proceso de forma infinita, teniendo un tiempo de espera
// de n minutos entre cada ejecución.
func runForever(parms parametros) error {
	var err error
	for {
		if err = run(parms); err != nil {
			break
		}

		// Se ejecuta Garbage Collector para liberar memoria luego del procesamiento
		// de archivos, pues en función de la cantidad de archivos procesados la
		// cantidad de memoria utilizada puede ser significativa.
		runtime.GC()

		time.Sleep(time.Minute * time.Duration(parms.yamlCfg.espera))
	}
	return err
}
