package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func printError(err error) {
	fmt.Println("* Error:", err)
}

func main() {

	ahora := time.Now().UTC()
	fDefault := ahora.Format("060102")

	fechaPtr := flag.String("f", fDefault, "Fecha del primer log a cargar (formato 'aammdd')")
	maxDiasPtr := flag.Int("l", 0, "Límite de fechas a procesar")

	flag.Parse()

	err := validaFecha(*fechaPtr)
	if err != nil {
		printError(err)
		os.Exit(0)
	}

	fmt.Println(*fechaPtr)

	rangoArchivos := genRangoArchivos(*fechaPtr, *maxDiasPtr)
	fmt.Printf("%v\n", rangoArchivos)
}
