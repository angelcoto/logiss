package proc

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/angelcoto/logiss/parm"
	"github.com/angelcoto/logiss/util"
)

const formatoFecha string = "2006-01-02 15:04:05"

// rec define la estructura de la línea de log que interesa generar.  No está
// constituida por todas las líneas del log original, sino que solo las que se
// consideran de interés.  En la estrucura se incorpora fechaLoc (fecha en hora
// local), que no corresponde a ningún campo de log original.
// (proc_files.go)
type rec struct {
	fechaOri      string
	fechaLoc      string
	metodo        string
	uriStem       string
	puerto        string
	usuario       string
	ipC           string
	referer       string
	status        string
	tiempoResp    string
	tiempoRespSec string
	dia           string
	hora          string
	minuto        string
}

// creaLinea es un método asociado al tipo rec, que procesa la línea del
// log original (lineaPartida) para asignar los valores de todos los
// campos de rec.
// (proc_files.go)
func (r *rec) creaLinea(lineaPartida []string) error {
	// Selección de los campos que conforman la línea de log de salida
	r.fechaOri = lineaPartida[0] + " " + lineaPartida[1]
	fechaOri, err := time.Parse(formatoFecha, r.fechaOri)
	if err != nil {
		return fmt.Errorf("no fue posible procesar supuesto campo de fecha %s en creaLinea: %v", r.fechaOri, err)
	}
	r.fechaLoc = fechaOri.Local().Format(formatoFecha)
	r.metodo = lineaPartida[3]
	r.uriStem = lineaPartida[4]
	r.puerto = lineaPartida[6]
	r.usuario = lineaPartida[7]
	r.ipC = lineaPartida[8]
	r.referer = lineaPartida[10]
	r.status = lineaPartida[11]
	r.tiempoResp = lineaPartida[14]

	t, err := strconv.ParseFloat(r.tiempoResp, 64)
	if err != nil {
		return fmt.Errorf("no fue posible procesar tiempo de respuesta %s", r.tiempoResp)
	}
	r.tiempoRespSec = fmt.Sprintf("%.6f", t/1000.0)

	fechaLoc, _ := time.Parse(formatoFecha, r.fechaLoc)
	r.dia = fechaLoc.Format("01/02")
	r.hora = fechaLoc.Format("15")
	r.minuto = fechaLoc.Format("15:04")

	return nil
}

// procArchivo revisa cada línea del archivo log para seleccionar
// los campos de interés.  Se desestiman las líneas de comentario, así como
// las líneas que contienen uri_stem declarados en lista de exclusión.
// Se devuelve error en caso que no se pueda procesar el archivo.
// (proc_files.go)
func procArchivo(archivo string, parms parm.Parametros) error {

	f, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("no fue posible abrir archivo log en procArchivo: %v", err)
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

		// Se excluyen las líneas de IP que no se quiere procesar
		if slices.Contains(parms.YamlCfg.ExclIP, lineaPartida[8]) {
			continue
		}

		// Se escluyen las líneas de las uri que no se quieren procesar
		if slices.Contains(parms.YamlCfg.ExclUri, lineaPartida[4]) {
			continue
		}

		// se crea la línea de interés que será agregada al log de salida (logTransformado)
		if err := linea.creaLinea(lineaPartida); err != nil {
			util.LogError(err)
			continue
		}

		logTransformado = append(logTransformado, linea)
	}

	lineasInsertadas, err := logTransformado.log2csv(parms.YamlCfg.CsvPath, parms.YamlCfg.CsvSeparador, parms.YamlCfg.ExclUrsNull)
	if err != nil {
		return err
	}

	util.LogMensaje(fmt.Sprint("Procesado: ", archivo, " - ", lineasInsertadas, " líneas en archivo CSV."))
	return nil
}
