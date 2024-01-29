package main

import (
	"io"
	"os"
)

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
