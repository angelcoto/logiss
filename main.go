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

	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n", "v1.0.0")

	arcTransferidos, err := tranfArchivos(conf.origen, conf.tmp, rangoArchivos)
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	err = procArchivos(arcTransferidos, conf.csvPath, conf.exclUrsNull)
	if err != nil {
		printError(err)
	}

	fmt.Printf("\nFIN DE PROCESO: %v\n", time.Since(inicio))

}
