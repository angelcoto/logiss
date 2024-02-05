package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func csvLog(entradas log, csv string) (int, error) {

	// Abre el archivo en modo append (agregar)
	archivo, err := os.OpenFile(csv, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return 0, err
	}
	defer archivo.Close()

	// Crea un escritor que apunta al archivo
	writer := bufio.NewWriter(archivo)

	contador := 0
	for _, entrada := range entradas {
		// En el arvhivo CSV no se graban las lineas que no tienen usuario
		if entrada.usuario == "-" {
			continue
		}

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
			return 0, err
		}
		contador++
	}

	writer.Flush()
	return contador, nil
}

func procArchivo(archivo, csvPath string) error {

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

		exclusiones := []string{"/bundles/modernizr", "/Content/css", "/bundles/jquery", "/bundles/bootstrap"}
		if slices.Contains(exclusiones, lineaPartida[4]) {
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

		// usuariols -
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

	lineasInsertadas, err := csvLog(lineasLog, csvPath)
	if err != nil {
		return err
	}

	fmt.Println("Procesado:", archivo, "->", lineasInsertadas, "líneas insertadas en archivo CSV.")
	return nil
}

func procArchivos(archivos rangeFile, csvPath string) error {
	fmt.Printf("\n* ETAPA 2: procesamiento de archivos en directorio temporal\n")

	csvPathBk := csvPath + ".bak"
	err := os.Rename(csvPath, csvPathBk)
	if err != nil {
		printError(err)
	}

	// Abre el archivo csv para agregar el encabezado
	archivo, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	// Escribe el encabezado
	encabezado := "fecha_ori,fecha,metodo,uri_stem,puerto,usuario,ip_c,referer,status,tiempo"
	_, err = archivo.WriteString(encabezado + "\n")
	if err != nil {
		archivo.Close()
		return err
	}
	archivo.Close()

	// Procesa todos los archivos transferidos
	for _, archivo := range archivos {
		err := procArchivo(archivo, csvPath)
		if err != nil {
			printError(err)
		}
	}

	return nil
}
