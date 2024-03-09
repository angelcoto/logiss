package main

import (
	"flag"
	"time"

	"github.com/angelcoto/logiss/util"
	"github.com/spf13/viper"
)

type minutos uint16

// cfg es la estructura para almacenar los valores definidos en
// el archivo de configuración yaml.
// (load_parms.go)
type cfg struct {
	origen      string
	destino     string
	csvPath     string
	exclUrsNull bool
	espera      minutos
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

	c.origen = viper.GetString("dirs.origen")
	c.destino = viper.GetString("dirs.destino")
	c.csvPath = viper.GetString("dirs.csv_path")
	c.exclUrsNull = viper.GetBool("mode.excluir_usuarios_nulos")
	c.espera = minutos(viper.GetUint16("mode.espera_ejecucion_continua"))

	return nil
}

// parametros es la estructura para almacenar todos los parámetros de
// funcionamiento, los cuales están conformados por los ingresados por
// línea de comando y por los contenidos en la estructura cfg.
// (load_parms.go)
type parametros struct {
	fechaInicial string
	dias         int
	continuo     bool
	yamlCfg      cfg
}

// loadParms llena los campos de la estructura parametros, lo cual incluye
// los ingresados desde línea de comando y los definidos en el archivo
// de configuración "conf.yaml"
// (load_parms.go)
func (p *parametros) loadParms() error {
	ahora := time.Now().UTC()
	fDefault := ahora.Format("060102")

	fechaPtr := flag.String("f", fDefault, "Fecha del primer log a procesar (formato 'aammdd')")
	maxDiasPtr := flag.Int("d", 0, "Cantidad de días a procesar desde la fecha definida")
	modoContinuo := flag.Bool("c", false, "Ejecución contínua")

	flag.Parse()

	if err := p.yamlCfg.loadCfg(); err != nil {
		return err
	}

	p.continuo = *modoContinuo

	if err := util.ValidaFecha(*fechaPtr); err != nil {
		return err
	}
	p.fechaInicial = *fechaPtr

	if *fechaPtr == fDefault {
		p.dias = 0
	} else {
		p.dias = *maxDiasPtr
	}

	return nil
}
