package main

import (
	"fmt"
	"path/filepath"

	"github.com/angelcoto/logiss/util"
)

// refreshCache actualiza los archivos logs desde el directorio origen al directorio caché.
// Solo se actualizan los archivos que han cambiado en el directorio origen.
// (refresh_cache.go)
func refreshCache(dirOrigen string, dirDestino string, archivos []string) ([]string, error) {
	util.LogMensaje("Inicia actualización de directorio caché desde directorio origen")
	var archivosListos []string
	for _, archivo := range archivos {
		fOri := filepath.Join(dirOrigen, archivo)
		fDes := filepath.Join(dirDestino, archivo)

		if util.ExisteArchivo(fDes) {
			if !util.ExisteArchivo(fOri) {
				util.LogError(fmt.Errorf("no se puede acceder archivo %s", fOri))
				archivosListos = append(archivosListos, fDes)
				continue
			}

			if iguales, _ := util.TamanosIguales(fOri, fDes); iguales {
				archivosListos = append(archivosListos, fDes)
				continue
			} else {
				if err := util.BorraArchivo(fDes); err != nil {
					util.LogError(err)
					continue
				}
			}
		}

		bWritten, err := util.CpFile(fDes, fOri)
		if err != nil {
			util.LogError(err)
			continue
		}
		archivosListos = append(archivosListos, fDes)
		util.LogMensaje((fmt.Sprint("Actualizado: ", fDes, " - ", bWritten, " bytes")))

	}

	return archivosListos, nil
}
