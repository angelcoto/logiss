package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

// rec define la estructura de la línea de log que interesa generar.  No está
// constituida por todas las líneas del log original, sino que solo las que se
// consideran de interés.  En la estrucura se incorpora fechaLoc (fecha en hora
// local), que no corresponde a ningún campo de log original.
// (proc_files.go)
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

// creaLinea es un método asociado al tipo rec, que procesa la línea del
// log original (lineaPartida) para asignar los valores de todos los
// campos de rec.
// (proc_files.go)
func (r *rec) creaLinea(lineaPartida []string) error {
	// Selección de los campos que conforman la línea de log de salida
	r.fechaOri = lineaPartida[0] + " " + lineaPartida[1]
	fechaOri, err := time.Parse("2006-01-02 15:04:05", r.fechaOri)
	if err != nil {
		return err
	}
	r.fechaLoc = fechaOri.Local().Format("2006-01-02 15:04:05")
	r.metodo = lineaPartida[3]
	r.uriStem = lineaPartida[4]
	r.puerto = lineaPartida[6]
	r.usuario = lineaPartida[7]
	r.ipC = lineaPartida[8]
	r.referer = lineaPartida[10]
	r.status = lineaPartida[11]
	r.tiempo = lineaPartida[14]

	return nil
}

type lineasLog []rec

// log2csv agrega las entradas del archivo log al archivo csv.  Devuelve
// la cantidad de líneas escritas y error en caso que no sea posible
// escribir en el archivo.
func (entradas lineasLog) log2csv(csv string, exclUsrNull bool) (int, error) {

	// Abre el archivo en modo append (agregar)
	archivo, err := os.OpenFile(csv, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		logError(err)
		return 0, err
	}
	defer archivo.Close()

	// Crea un escritor que apunta al archivo
	writer := bufio.NewWriter(archivo)

	contador := 0
	for _, entrada := range entradas {
		// En el arvhivo CSV no se graban las lineas que no tienen usuario
		if entrada.usuario == "-" && exclUsrNull {
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

		// Escribe línea en el buffer
		if _, err := writer.WriteString(linea); err != nil {
			return 0, err
		}
		contador++
	}

	writer.Flush()
	return contador, nil
}

// procArchivo revisa cada línea del archivo log para seleccionar
// los campos de interés.  Se desestiman las líneas de comentario, así como
// las líneas que contienen uri_stem declarados en lista de exclusión.
// Se devuelve error en caso que no se pueda procesar el archivo.
// (proc_files.go)
func procArchivo(archivo, csvPath string, exclUsrNull bool) error {

	f, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer f.Close()

	var linea rec
	var lineaPartida []string
	var contador int32
	var logTransformado lineasLog

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

		// se crea la línea de interés que será agregada al log de salida (logTransformado)
		if err := linea.creaLinea(lineaPartida); err != nil {
			logError(err)
			continue
		}

		logTransformado = append(logTransformado, linea)
	}

	lineasInsertadas, err := logTransformado.log2csv(csvPath, exclUsrNull)
	if err != nil {
		return err
	}

	logMensaje(fmt.Sprint("Procesado: ", archivo, " - ", lineasInsertadas, " líneas en archivo CSV."))
	return nil
}

// procArchivos invoca el procesamiento de cada uno de los archivos log listados
// listados en la variable archivos.  Realiza una un backup del archivo csv y
// coloca la línea de encabezado en el archivo.  Error es devuelto en caso que
// el proceso no pueda realizarse.
// (proc_files.go)
func procArchivos(archivos rangeFile, csvPath string, exclUsrNull bool) error {
	logMensaje("Inicia procesamiento de archivos en directorio cache")

	csvPathBk := csvPath + ".bak"

	if err := os.Rename(csvPath, csvPathBk); err != nil {
		logError(err)
	}

	// Abre el archivo csv para agregar el encabezado
	archivo, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	// Escribe el encabezado
	encabezado := "fecha_ori,fecha,metodo,uri_stem,puerto,usuario,ip_c,referer,status,tiempo"
	if _, err := archivo.WriteString(encabezado + "\n"); err != nil {
		archivo.Close()
		return err
	}
	archivo.Close()

	// Procesa todos los archivos transferidos
	for _, archivo := range archivos {
		err := procArchivo(archivo, csvPath, exclUsrNull)
		if err != nil {
			logError(err)
		}
	}

	return nil
}
