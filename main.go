package main

import (
	"fmt"
	"os"

	//	_ "net/http/pprof"

	"github.com/angelcoto/logiss/parm"
	"github.com/angelcoto/logiss/util"
)

func main() {
	/*
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	*/
	fmt.Printf("logiss %s.  Copyright (c) 2024 Ángel Coto.  MIT License.\n\n", "v1.4.1")

	// Estructura con parámetros de funcionamiento
	var parms parm.Parametros

	// Carga parámetros
	if err := parms.LoadParms(); err != nil {
		util.LogError(err)
		os.Exit(0)
	}

	switch parms.Continuo {
	case false:
		runOnce(parms)
	case true:
		runForever(parms)
	}

}
