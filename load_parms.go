package main

import (
	"flag"
	"time"

	"github.com/spf13/viper"
)

type cfg struct {
	origen      string
	destino     string
	csvPath     string
	exclUrsNull bool
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

	return nil
}

// loadParms obtiene parámetros de funcionamiento a través de flags de línea de comando.
// (load_parms.go)
func loadParms() (cfg, string, int, error) {
	ahora := time.Now().UTC()
	fDefault := ahora.Format("060102")

	fechaPtr := flag.String("f", fDefault, "Fecha del primer log a cargar en formato 'aammdd'")
	maxDiasPtr := flag.Int("l", 0, "Límite de fechas a procesar")

	flag.Parse()

	var c cfg

	if err := c.loadCfg(); err != nil {
		return cfg{}, "", 0, err
	}

	if err := validaFecha(*fechaPtr); err != nil {
		return cfg{}, "", 0, err
	}

	if *fechaPtr == fDefault {
		*maxDiasPtr = 0
	}

	return c, *fechaPtr, *maxDiasPtr, nil
}
