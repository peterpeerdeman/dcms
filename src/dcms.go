package main

import (
	"encoding/json"
	"flag"
	"github.com/alleveenstra/godb"
	"github.com/ziutek/mymysql/autorc"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"site"
	"resty"
	"net/http"
)

var config struct {
	SqlHostname       string
	SqlUser           string
	SqlPassword       string
	SqlDatabase       string
	ListenAddress     string
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

	if 1 == 2 {
		godb.Database = autorc.New("tcp", "", config.SqlHostname, config.SqlUser, config.SqlPassword, config.SqlDatabase)
		godb.Database.Register("SET NAMES utf8")
	}

	resty.Cms()
	site.Site()

	http.ListenAndServe(config.ListenAddress, nil)

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
