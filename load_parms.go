package main

import (
	"flag"
	"time"

	"github.com/spf13/viper"
)

type cfg struct {
	origen  string
	tmp     string
	archive string
	csvPath string
}

func (c *cfg) loadCfg(cfgFile string) error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	c.origen = viper.GetString("dirs.origen")
	c.tmp = viper.GetString("dirs.tmp")
	c.archive = viper.GetString("dirs.archive")
	c.csvPath = viper.GetString("dirs.csvPath")

	return nil

}

func loadParms() (cfg, string, int, error) {
	ahora := time.Now().UTC()
	fDefault := ahora.Format("060102")

	fechaPtr := flag.String("f", fDefault, "Fecha del primer log a cargar en formato 'aammdd'")
	maxDiasPtr := flag.Int("l", 0, "Límite de fechas a procesar")

	flag.Parse()

	var c cfg
	err := c.loadCfg("conf.yaml")
	if err != nil {
		return cfg{}, "", 0, err
	}

	err = validaFecha(*fechaPtr)
	if err != nil {
		return cfg{}, "", 0, err
	}

	if *fechaPtr == fDefault {
		*maxDiasPtr = 0
	}

	return c, *fechaPtr, *maxDiasPtr, err
}
