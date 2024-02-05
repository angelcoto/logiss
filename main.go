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
	dirs, fecha, maxDias, err := loadParms()
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	inicio := time.Now()

	rangoArchivos := genRangoArchivos(fecha, maxDias)

	arcTransferidos, err := tranfArchivos(dirs.origen, dirs.tmp, rangoArchivos)
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	procArchivos(arcTransferidos, dirs.csvPath)

	fmt.Printf("\nFIN DE PROCESO: %v\n", time.Since(inicio))

}
