package main

import (
	"errors"
	"time"
)

type rangeFile []string

// validaFecha verifica que el valor ingresado sea una fecha válida.
func validaFecha(sfecha string) error {

	fecha, err := time.Parse("060102", sfecha)
	if err != nil {
		return errors.New("la fecha debe estar en formato 'aammdd'")
	}

	ahora := time.Now().UTC()
	if ahora.Before(fecha) {
		return errors.New("la fecha es mayor a la fecha de ahora")
	}
	return err

}

// genRandoArchivos crea un listado de nombres de archivos, correspondiendo
// cada uno a una fecha dentro del rango indicado por la fecha de inicio y la
// fecha actual.
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
