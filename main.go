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
	mongo        = flag.String("mongo", "mongodb://RTI:32768/Vault", "mongoDB server address.")
)

func main() {

	flag.Parse()
	//go readUDP()

	ip, _ := GetLocalIP()
	fmt.Printf("Webserver: %s:%s\n", ip, *addr)
	fmt.Printf("Exerntal IP: %s\n", GetExternalIP())

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
	mux.HandleFunc("/so", http.HandlerFunc(salesOrderHandler))
	mux.HandleFunc("/so/add", http.HandlerFunc(salesOrderAddHandler))
	mux.HandleFunc("/so/update/:id", http.HandlerFunc(salesOrderUpdateHandler))
	mux.HandleFunc("/rma", http.HandlerFunc(rmaHandler))
	mux.HandleFunc("/rma/add", http.HandlerFunc(rmaAddHandler))
	mux.HandleFunc("/rma/update/:id", http.HandlerFunc(rmaUpdateHandler))
	mux.HandleFunc("/rma/report/:id", http.HandlerFunc(rmaReportHandler))
	mux.HandleFunc("/product", http.HandlerFunc(productListHandler))
	mux.HandleFunc("/product/add", http.HandlerFunc(productAddHandler))
	mux.HandleFunc("/product/update/:id", http.HandlerFunc(productUpdateHandler))
	mux.HandleFunc("/adcp_checklist/add", http.HandlerFunc(adcpChecklistAddHandler))
	mux.HandleFunc("/adcp_checklist/update/:id", http.HandlerFunc(adcpChecklistUpdateHandler))

	// HTTP server
	if err := http.ListenAndServe(*addr, mux); err != nil {
		fmt.Printf("Error trying to bind to port: %v, so exiting...", err)
		log.Fatal("Error ListenAndServe:", err)
	}

}
