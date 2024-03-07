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
		printError(err)
		os.Exit(0)
	}

	inicio := time.Now()

	rangoArchivos := genRangoArchivos(fecha, maxDias)

	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n", "v1.1.0")

	archivos, err := tranfArchivos(conf.origen, conf.destino, rangoArchivos)
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	if err := procArchivos(archivos, conf.csvPath, conf.exclUrsNull); err != nil {
		printError(err)
	}

	fmt.Printf("\nFIN DE PROCESO: %v\n", time.Since(inicio))

}
