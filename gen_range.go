package main

import (
	"time"
)

type rangeFile []string

// genRandoArchivos crea un listado de nombres de archivos, correspondiendo
// cada uno a una fecha dentro del rango indicado por la fecha de inicio y la
// fecha actual.
// (gen_range.go)
func genRangoArchivos(sfechaInicio string, maxFechas int) rangeFile {
	var archivos rangeFile

	f, _ := time.Parse("060102", sfechaInicio)
	ahora := time.Now().UTC()

	limite := int(ahora.Sub(f).Hours() / 24) // Días entre la fecha de ahora UTC y la fecha inicial

	if maxFechas > 0 {
		limite = maxFechas - 1
	}

	for i := 0; i <= limite; i++ {
		fi := f.AddDate(00, 00, i)
		sf := fi.Format("060102")
		archivo := "u_ex" + sf + ".log"
		archivos = append(archivos, archivo)
	}
	return archivos
}
