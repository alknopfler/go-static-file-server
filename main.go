package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	dT time.Duration

)

func main() {
	deadTime := flag.String("t", "10m", "Tiempo de duracion del enlace")
	dirToShare := flag.String( "d", "./", "Directorio que vamos a compartir")
	_ = flag.String("h", "", "Options available: -t=10m , -d='/home' ")

	flag.Parse()
	dT, _ := time.ParseDuration(*deadTime)

	http.Handle("/", http.FileServer(http.Dir(*dirToShare)))
	log.Println("Starting to share " + *dirToShare + " during " + *deadTime)

	go http.ListenAndServe(":8886", nil)

	select {
		case <-time.After(dT):
			log.Println("Run out of time to share...")
			os.Exit(0)
	}
	
}
