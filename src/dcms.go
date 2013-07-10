package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"resty"
	"runtime/debug"
	"site"
	"storage"
)

var config struct {
	SqlHostname   string
	SqlUser       string
	SqlPassword   string
	SqlDatabase   string
	ListenAddress string
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "The configuration file.")
	flag.Parse()

	configErr := ReadConfiguration(configFile)
	if configErr != nil {
		flag.Usage()
		os.Exit(1)
	}

	storage.Init()

	resty.Cms()
	site.Site()

	log.Println("starting to listen on " + config.ListenAddress)
	listenErr := http.ListenAndServe(config.ListenAddress, nil)
	if listenErr != nil {
		log.Fatal(listenErr)
	} 
}

func ReadConfiguration(configFile string) error {
	file, readErr := ioutil.ReadFile(configFile)
	if readErr != nil {
		return readErr
	}
	jsonErr := json.Unmarshal(file, &config)
	Fatal(jsonErr)
	return nil
}

func Fatal(err error) {
	if err != nil {
		log.Fatalf("%v", err)
		debug.PrintStack()
		os.Exit(-1)
	}
}
