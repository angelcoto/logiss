package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type rec struct {
	fechaOri string
	fechaLoc string
	metodo   string
	uriStem  string
	puerto   string
	usuario  string
	ipC      string
	referer  string
	status   string
	tiempo   string
}

type log []rec

func csvLog(entradas log) error {
	csv := "360.csv"

	// Abre el archivo en modo append (agregar)
	archivo, err := os.OpenFile(csv, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return err
	}
	defer archivo.Close()

	// Crea un escritor que apunta al archivo
	writer := bufio.NewWriter(archivo)

	for _, entrada := range entradas {
		linea := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			entrada.fechaOri,
			entrada.fechaLoc,
			entrada.metodo,
			entrada.uriStem,
			entrada.puerto,
			entrada.usuario,
			entrada.ipC,
			entrada.referer,
			entrada.status,
			entrada.tiempo)
		_, err := writer.WriteString(linea)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	fmt.Println("Se insertaron", len(entradas), "lineas")
	return nil
}

func procArchivo(archivo string) error {

	fmt.Println("Procesando archivo", archivo)

	f, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer f.Close()

	var linea rec
	var lineaPartida []string
	var contador int32
	var lineasLog log

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		contador++
		lineaPartida = strings.Split(scanner.Text(), " ")

		// Se omiten las líneas que comienzan con "#" por ser comentarios
		if string(lineaPartida[0][0]) == "#" {
			continue
		}

		// fecha_ori
		linea.fechaOri = lineaPartida[0] + " " + lineaPartida[1]

		// fecha
		fechaOri, err := time.Parse("2006-01-02 15:04:05", linea.fechaOri)
		if err != nil {
			fmt.Println(err, "en línea", contador)
			continue
		}
		linea.fechaLoc = fechaOri.Local().Format("2006-01-02 15:04:05")

		// metodo
		linea.metodo = lineaPartida[3]

		// uri_stem
		linea.uriStem = lineaPartida[4]

		// puerto
		linea.puerto = lineaPartida[6]

		// usuario
		linea.usuario = lineaPartida[7]

		// ip_c
		linea.ipC = lineaPartida[8]

		// referer
		linea.referer = lineaPartida[10]

		// status
		linea.status = lineaPartida[11]

		// tiempo
		linea.tiempo = lineaPartida[14]

		lineasLog = append(lineasLog, linea)

	}

	err = csvLog(lineasLog)
	if err != nil {
		return err
	}

	return nil
}

func procArchivos(archivos rangeFile) {
	fmt.Printf("\n* ETAPA 2: procesamiento de archivos en directorio temporal\n")

	for _, archivo := range archivos {
		fmt.Println(archivo)
		err := procArchivo(archivo)
		if err != nil {
			printError(err)
		}
	}
}
