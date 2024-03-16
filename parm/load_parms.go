package parm

import (
	"flag"
	"fmt"
	"slices"
	"time"

	"github.com/angelcoto/logiss/util"
	"github.com/spf13/viper"
)

type minutos uint16

// cfg es la estructura para almacenar los valores definidos en
// el archivo de configuración yaml.
// (load_parms.go)
type cfg struct {
	Origen       string
	Destino      string
	CsvPath      string
	CsvSeparador string
	Espera       minutos
	ExclUrsNull  bool
	ExclIP       []string
	ExclUri      []string
}

// loadCfg obtiene parámetros de configuración desde el archivo conf.yaml.
// (load_parms.go)
func (c *cfg) loadCfg() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	c.Origen = viper.GetString("dirs.origen")
	c.Destino = viper.GetString("dirs.destino")
	c.CsvPath = viper.GetString("dirs.csv_path")
	c.Espera = minutos(viper.GetUint16("mode.espera_ejecucion_continua"))
	c.ExclUrsNull = viper.GetBool("excl.excluir_usuarios_nulos")
	c.ExclIP = viper.GetStringSlice("excl.ip_excluidas")
	c.ExclUri = viper.GetStringSlice("excl.uri_excluidas")
	c.CsvSeparador = viper.GetString("csv.separador")

	// Valida el separador definido en el yaml
	if err := func() error {
		var separadores = []string{"coma", "tab", ";", "|"}
		if !slices.Contains(separadores, c.CsvSeparador) {
			return (fmt.Errorf("el separador definido en archivo de configuración no es válido"))
		}
		return nil
	}(); err != nil {
		return err
	}

	switch c.CsvSeparador {
	case "tab":
		c.CsvSeparador = "\t"
	case "coma":
		c.CsvSeparador = ","
	}

	return nil
}

// parametros es la estructura para almacenar todos los parámetros de
// funcionamiento, los cuales están conformados por los ingresados por
// línea de comando y por los contenidos en la estructura cfg.
// (load_parms.go)
type Parametros struct {
	FechaInicial string
	Dias         int
	Continuo     bool
	YamlCfg      cfg
}

// loadParms llena los campos de la estructura parametros, lo cual incluye
// los ingresados desde línea de comando y los definidos en el archivo
// de configuración "conf.yaml"
// (load_parms.go)
func (p *Parametros) LoadParms() error {
	ahora := time.Now().UTC()
	fDefault := ahora.Format("060102")

	fechaPtr := flag.String("f", fDefault, "Fecha del primer log a procesar (formato 'aammdd')")
	maxDiasPtr := flag.Int("d", 0, "Cantidad de días a procesar desde la fecha definida")
	modoContinuo := flag.Bool("c", false, "Ejecución contínua")

	flag.Parse()

	if err := p.YamlCfg.loadCfg(); err != nil {
		return err
	}

	p.Continuo = *modoContinuo

	if err := util.ValidaFecha(*fechaPtr); err != nil {
		return err
	}
	p.FechaInicial = *fechaPtr

	if *fechaPtr == fDefault {
		p.Dias = 0
	} else {
		p.Dias = *maxDiasPtr
	}

	return nil
}
