package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

///
/// Flags to set at startup
///
var (
	version      = "0.1"
	versionFloat = float32(0.1)
	addr         = flag.String("addr", ":8989", "http service address")
	port         = flag.String("port", "", "Serial COM Port")
	baud         = flag.String("baud", "115200", "Baud Rate")
	mongo        = flag.String("mongo", "mongodb://192.168.0.108:27017/Vault", "mongoDB server address.")
)

func main() {

	flag.Parse()
	//go readUDP()

	// Connect to DB
	go DbConnect(*mongo)

	// Router
	mux := bone.New()
	mux.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("libs"))))       // External libs
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images")))) // Image folder
	mux.HandleFunc("/adcp", http.HandlerFunc(adcpListHandler))
	mux.HandleFunc("/adcp/update/:id", http.HandlerFunc(adcpUpdateHandler))
	mux.HandleFunc("/adcp/cert/:id", http.HandlerFunc(adcpCertHandler))
	mux.HandleFunc("/adcp/add", http.HandlerFunc(adcpAddHandler))

	// HTTP server
	// http.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("libs")))) // External libs
	// http.HandleFunc("/adcp", adcpListHandler)                                            // List ADCP
	// http.HandleFunc("/adcp/update", adcpUpdateHandler)                                   // Update ADCP
	if err := http.ListenAndServe(*addr, mux); err != nil {
		fmt.Printf("Error trying to bind to port: %v, so exiting...", err)
		log.Fatal("Error ListenAndServe:", err)
	}

}
