package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

var config Config

func main() {
	config, _ = parseConfig("config.json")

	r := mux.NewRouter()
	r.HandleFunc("/volume/{vol}", VolumeHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

func getConn() net.Conn {
	address := fmt.Sprintf("%s:%d", config.DenonIP, config.DenonPort)
	conn, _ := net.Dial("tcp", address)
	return conn
}

func execCommand(command string) {
	conn := getConn()
	sCmd := fmt.Sprintf("%s\r", command)
	conn.Write([]byte(sCmd))
	conn.Close()
}

func VolumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vol := vars["vol"]

	execCommand(fmt.Sprintf("MV%s", vol))
}
