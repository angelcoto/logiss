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

/*
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
*/

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

// existeArchivo verifica si un archivo existe
func existeArchivo(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// tamanosIguales compara el tamaño de dos archivos y devuelve true en caso que los tamaños coincidan.
func tamanosIguales(f1, f2 string) (bool, error) {
	infoOri, err := os.Stat(f1)
	if err != nil {
		return false, err
	}

	infoDes, err := os.Stat(f2)
	if err != nil {
		return false, err
	}

	if infoOri.Size() == infoDes.Size() {
		return true, nil
	}

	return false, nil
}

// borraArchivo borra un archivo, devolviendo error en caso de falla
func borraArchivo(ruta string) error {
	err := os.Remove(ruta)
	if err != nil {
		return err
	}
	return nil
}
