package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Carga parámetros de funcionamiento
	conf, fecha, maxDias, err := loadParms()
	if err != nil {
		logError(err)
		os.Exit(0)
	}

	inicio := time.Now()

	rangoArchivos := genRangoArchivos(fecha, maxDias)

	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n\n", "v1.1.2")

	archivos, err := refreshCache(conf.origen, conf.destino, rangoArchivos)
	if err != nil {
		logError(err)
		os.Exit(0)
	}

	if err := procArchivos(archivos, conf.csvPath, conf.exclUrsNull); err != nil {
		logError(err)
	}

	logMensaje(fmt.Sprintf("Proceso finalizado en %v", time.Since(inicio)))

}
