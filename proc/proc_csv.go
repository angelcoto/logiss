package proc

import (
	"bufio"
	"fmt"
	"os"

	"github.com/angelcoto/logiss/util"
)

type lineasLog []rec

// Crea un cache para la indagación del DNS
var cacheIP util.IPCache = make(util.IPCache)

// log2csv agrega las entradas del archivo log al archivo csv.  Devuelve
// la cantidad de líneas escritas y error en caso que no sea posible
// escribir en el archivo.
func (entradas lineasLog) log2csv(csv, sep string, exclUsrNull bool) (int, error) {

	// Abre el archivo en modo append (agregar)
	archivo, err := os.OpenFile(csv, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, fmt.Errorf("no fue posible abrir archivo CSV en log2csv: %v", err)
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

		linea := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%sh: %s%s%s%s%s%s%s%s%s%s%s%s%s\n",
			entrada.fechaLoc, sep,
			entrada.metodo, sep,
			entrada.uriStem, sep,
			entrada.puerto, sep,
			entrada.usuario, sep,
			util.Hostname(entrada.ipC, cacheIP), sep,
			entrada.referer, sep,
			entrada.status, sep,
			entrada.tiempoRespSec, sep,
			entrada.dia, sep,
			entrada.hora, sep,
			entrada.minuto)

		// Escribe línea en el buffer
		if _, err := writer.WriteString(linea); err != nil {
			return 0, fmt.Errorf("no fue posible escribir línea en archivo CSV en log2csv: %v", err)
		}
		contador++
	}

	writer.Flush()
	return contador, nil
}
