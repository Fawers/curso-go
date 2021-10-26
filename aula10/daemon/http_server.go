package daemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (d *Daemon) startHTTPServer(comm chan<- int) {
	http.HandleFunc("/status", d.endpointStatus)
	http.HandleFunc("/services", d.endpointServices)
	log.Println("[daemon] Starting server...")

	err := http.ListenAndServe(d.HTTPEndpoint, nil)
	if err != nil {
		log.Printf("[daemon] Error initializing the server: %s\n", err)
		comm <- 1
	}

	comm <- 0
}

func (d *Daemon) endpointStatus(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "OK")
}

func (d *Daemon) endpointServices(rw http.ResponseWriter, req *http.Request) {
	services := d.store.List()

	payload, err := json.Marshal(services)
	if err != nil {
		log.Println("[daemon] unable to serialize services into json")
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, "fail")
		return
	}

	fmt.Fprintln(rw, string(payload))
}
