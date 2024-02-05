package main

import (
	"fmt"
	"os"
	"time"
)

func printError(err error) {
	fmt.Println("* Error:", err)
}

func main() {

	// Carga parámetros de funcionamiento
	conf, fecha, maxDias, err := loadParms()
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	inicio := time.Now()

	rangoArchivos := genRangoArchivos(fecha, maxDias)

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
