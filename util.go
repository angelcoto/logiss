package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// printError imprime en pantalla, en forma estandarizada, el mensaje de ero
// r recibido como parámetro.
// (util.go)
func printError(err error) {
	fmt.Println("* Error:", err)
}

// validaFecha verifica que el valor ingresado sea una fecha válida.
// (util.go)
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

// cpFile copia el archivo oriFileName en destFileName, devolviendo
// el número de bytes copiados y una variable de error.
// (util.go)
func cpFile(destFileName string, oriFileName string) (int64, error) {
	oriFile, err := os.Open(oriFileName)
	if err != nil {
		return 0, err
	}
	defer oriFile.Close()

	// Crear o truncar el archivo de destino
	destFile, err := os.Create(destFileName)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	// Copiar el contenido del archivo de origen al archivo de destino
	bytesWritten, err := io.Copy(destFile, oriFile)
	if err != nil {
		return 0, err
	}

	return bytesWritten, nil
}
