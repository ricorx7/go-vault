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

	//reactServ := http.FileServer(http.Dir("react"))
	//http.Handle("/react", reactServ)

	// Router
	mux := bone.New()
	mux.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("libs"))))       // External libs
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images")))) // Image folder

	//reactServ := http.FileServer(http.Dir("react"))
	//mux.Handle("/react", reactServ)
	mux.Handle("/react/", http.StripPrefix("/react/", http.FileServer(http.Dir("react")))) // Image folder
	mux.HandleFunc("/api", http.HandlerFunc(apiHandler))

	mux.HandleFunc("/", http.HandlerFunc(adcpListHandler))
	mux.HandleFunc("/adcp", http.HandlerFunc(adcpListHandler))
	mux.HandleFunc("/adcp/update/:id", http.HandlerFunc(adcpUpdateHandler))
	mux.HandleFunc("/adcp/cert/:id", http.HandlerFunc(adcpCertHandler))
	mux.HandleFunc("/adcp/add", http.HandlerFunc(adcpAddHandler))
	mux.HandleFunc("/adcp/wt", http.HandlerFunc(watertestListHandler))
	mux.HandleFunc("/adcp/wt/update/:id", http.HandlerFunc(watertestUpdateHandler))
	mux.HandleFunc("/adcp/wt/add", http.HandlerFunc(watertestAddHandler))
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
	mux.HandleFunc("/vault/tank", http.HandlerFunc(vaultAPITankHandler))
	mux.HandleFunc("/vault/tank/:id", http.HandlerFunc(vaultAPITankSerialHandler))
	mux.HandleFunc("/vault/tank/selected/:id", http.HandlerFunc(vaultAPITankSelectedSerialHandler))
	mux.HandleFunc("/vault/tank/selected/moving/:id", http.HandlerFunc(vaultAPITankSelectedSerialMovingHandler))
	mux.HandleFunc("/vault/tank/selected/noise/:id", http.HandlerFunc(vaultAPITankSelectedSerialNoiseHandler))
	mux.HandleFunc("/vault/tank/selected/ringing/:id", http.HandlerFunc(vaultAPITankSelectedSerialRingingHandler))
	mux.HandleFunc("/react1", http.HandlerFunc(reactHandler))

	// HTTP server
	if err := http.ListenAndServe(*addr, mux); err != nil {
		fmt.Printf("Error trying to bind to port: %v, so exiting...", err)
		log.Fatal("Error ListenAndServe:", err)
	}

}
