package main

import (
	"fmt"
	"time"
)

// runOnce ejecuta el proceso completo una vez.
// (run.go)
func runOnce(conf cfg, fecha string, maxDias int) error {
	inicio := time.Now()
	rangoArchivos := genRangoArchivos(fecha, maxDias)

	archivos, err := refreshCache(conf.origen, conf.destino, rangoArchivos)
	if err != nil {
		logError(err)
		return err
	}

	if err := procArchivos(archivos, conf.csvPath, conf.exclUrsNull); err != nil {
		logError(err)
		return err
	}

	logMensaje(fmt.Sprintf("Proceso finalizado en %v", time.Since(inicio)))

	return nil

}
