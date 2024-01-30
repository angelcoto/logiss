package main

import (
	"fmt"
	"path/filepath"
)

func cpArchivos(dirOrigen string, dirDestino string, archivos rangeFile) {
	fmt.Println("Etapa 1: transferencia de logs desde", dirOrigen)
	for _, archivo := range archivos {
		fOri := filepath.Join(dirOrigen, archivo)
		fDes := filepath.Join(dirDestino, archivo)

		bWritten, err := cpFile(fDes, fOri)
		if err != nil {
			printError(err)
		} else {
			fmt.Println("Transferido:", fDes, "-", bWritten)

		}

	}
}
