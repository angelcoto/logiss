package main

import (
	"fmt"
	"os"

	"github.com/angelcoto/logiss/util"
)

func main() {

	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n\n", "v1.2.0")

	// Estructura con parámetros de funcionamiento
	var parms parametros

	// Carga parámetros
	err := parms.loadParms()
	if err != nil {
		util.LogError(err)
		os.Exit(0)
	}

	switch parms.continuo {
	case false:
		if err := runOnce(parms); err != nil {
			util.LogError(err)
			os.Exit(0)

		}
	case true:
		if err := runForever(parms); err != nil {
			util.LogError(err)
			os.Exit(0)

		}
	}

}
