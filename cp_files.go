package main

import (
	"fmt"
	"path/filepath"
)

// tranfAarchivos copia los archivos logs desde el directorio dirOrigen al directorio dirDestino.
// Previo a la transferencia limpia el contenido del directorio temporal (dirDestino).
func tranfArchivos(dirOrigen string, dirDestino string, archivos rangeFile) ([]string, error) {
	/* 	err := borraTmp(dirDestino)
	   	if err != nil {
	   		return nil, err
	   	}
	*/
	fmt.Printf("\n* ETAPA 1: actualizando directorio caché desde directorio origen %s\n", dirOrigen)
	var archivosListos []string
	for _, archivo := range archivos {
		fOri := filepath.Join(dirOrigen, archivo)
		fDes := filepath.Join(dirDestino, archivo)

		if existeArchivo(fDes) {
			if iguales, _ := tamanosIguales(fOri, fDes); iguales {
				archivosListos = append(archivosListos, fDes)
				continue
			} else {
				if err := borraArchivo(fDes); err != nil {
					printError(err)
					continue
				}
			}
		}

		bWritten, err := cpFile(fDes, fOri)
		if err != nil {
			printError(err)
			continue
		}
		archivosListos = append(archivosListos, fDes)
		fmt.Println("Transferido:", fDes, "-", bWritten, "bytes")

	}

	return archivosListos, nil
}
