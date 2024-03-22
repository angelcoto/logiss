package proc

import (
	"fmt"
	"os"

	"github.com/angelcoto/logiss/parm"
	"github.com/angelcoto/logiss/util"
)

// ProcArchivos invoca el procesamiento de cada uno de los archivos log listados
// listados en la variable archivos.  Realiza una un backup del archivo csv y
// coloca la línea de encabezado en el archivo.  Error es devuelto en caso que
// el proceso no pueda realizarse.
// (proc_files.go)
func ProcArchivos(archivos []string, parms parm.Parametros) error {
	util.LogMensaje("Inicia procesamiento de archivos en directorio cache")

	archivo, err := os.Create(parms.YamlCfg.CsvPath)
	if err != nil {
		return fmt.Errorf("no fue posible crear archivo CSV en ProcArchivos: %v", err)
	}

	// Escribe el encabezado
	encabezado := fmt.Sprintf("fecha%smetodo%suri_stem%spuerto%susuario%sip_c%sreferer%sstatus%ssecs%sdia%shora%sminuto",
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador,
		parms.YamlCfg.CsvSeparador)
	if _, err := archivo.WriteString(encabezado + "\n"); err != nil {
		archivo.Close()
		return fmt.Errorf("no fue posible escribir encabezado en archivo CSV en ProcArchivos: %v", err)
	}
	archivo.Close()

	// Procesa todos los archivos transferidos
	for _, archivo := range archivos {
		err := procArchivo(archivo, parms)
		if err != nil {
			util.LogError(err)
		}
	}

	return nil
}
