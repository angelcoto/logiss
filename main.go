package main

import (
	"fmt"
	"os"
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

	rangoArchivos := genRangoArchivos(fecha, maxDias)

	cpArchivos(dirs.origen, dirs.tmp, rangoArchivos)
}
