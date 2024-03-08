package main

import (
	"fmt"
	"path/filepath"
)

// refreshCache actualiza los archivos logs desde el directorio origen al directorio caché.
// Solo se actualizan los archivos que han cambiado en el directorio origen.
// (refresh_cache.go)
func refreshCache(dirOrigen string, dirDestino string, archivos rangeFile) ([]string, error) {
	logMensaje("Inicia actualización de directorio caché desde directorio origen")
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
					logError(err)
					continue
				}
			}
		}

		bWritten, err := cpFile(fDes, fOri)
		if err != nil {
			logError(err)
			continue
		}
		archivosListos = append(archivosListos, fDes)
		logMensaje((fmt.Sprint("Actualizado: ", fDes, " - ", bWritten, "bytes")))

	}

	return archivosListos, nil
}
