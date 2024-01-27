package main

import (
	"errors"
	"time"
)

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

func genRangoArchivos(sfechaInicio string, maxFechas int) []string {
	var archivos []string

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
