package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func borraTmp(tmp string) error {
	err := filepath.Walk(tmp, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err := os.Remove(ruta)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// tranfAarchivos copia los archivos logs desde el directorio dirOrigen al directorio dirDestino.
// Previo a la transferencia limpia el contenido del directorio temporal (dirDestino).
func tranfArchivos(dirOrigen string, dirDestino string, archivos rangeFile) (rangeFile, error) {
	err := borraTmp(dirDestino)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n* ETAPA 1: transferencia de logs desde %s\n", dirOrigen)
	var arcCopiados rangeFile
	for _, archivo := range archivos {
		fOri := filepath.Join(dirOrigen, archivo)
		fDes := filepath.Join(dirDestino, archivo)

		bWritten, err := cpFile(fDes, fOri)
		if err != nil {
			printError(err)
		} else {
			fmt.Println("Transferido:", fDes, "-", bWritten)
			arcCopiados = append(arcCopiados, fDes)

		}

	}

	return arcCopiados, nil
}
