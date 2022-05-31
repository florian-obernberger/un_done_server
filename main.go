package main

import (
	"fmt"
	"net/http"
	"os"
	"unDoneServer/api"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

const progName string = "undoneserv"

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	p := argparse.NewParser(progName, "A Server instance to sync TodoEntries with the UnDone App.")
	port := p.Int("p", "port", &argparse.Options{
		Help:    "The port on which the server will be running",
		Default: 8080,
	})

	if err := p.Parse(os.Args); err != nil {
		fmt.Print(p.Usage(err))
	}

	portString := fmt.Sprintf(":%d", *port)

	srv := api.InitServer()
	log.Infof("Starting Server on Port %d", *port)
	err := http.ListenAndServeTLS(portString, "cert/localhost.crt", "cert/localhost.decrypted.key", srv)
	log.Fatalf(err.Error())
}
