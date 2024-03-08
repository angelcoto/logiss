package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n\n", "v1.1.3")

	// Carga parámetros de funcionamiento
	conf, fecha, maxDias, err := loadParms()
	if err != nil {
		logError(err)
		os.Exit(0)
	}

	// Ejecución de una vez
	if err := runOnce(conf, fecha, maxDias); err != nil {
		logError(err)
		os.Exit(0)

	}

}
