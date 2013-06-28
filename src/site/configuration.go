package site

import (
	"encoding/json"
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"log"
	"resty"
	"time"
)

var SiteConfiguration struct {
	Channels map[string]Channel
}

func WatchConfiguration(configFile string) error {
	for {
		readErr := readConfiguration(configFile)
		if readErr != nil {
			log.Printf("Error reading configuration: %v", readErr)
		}
		time.Sleep(time.Second)
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}
		err = watcher.Watch(configFile)
		if err != nil {
			return err
		}
		select {
		case <-watcher.Event:
		case err := <-watcher.Error:
			log.Printf("Error while watching configuration: %v", err)
		}
		watcher.Close()
	}
	return nil
}

func readConfiguration(configFile string) error {
	log.Printf("Loading configuration...")
	resty.NotificationChannel <- resty.Notification{"configuration", "Configuration reloaded"}
	file, readErr := ioutil.ReadFile(configFile)
	if readErr != nil {
		return readErr
	}
	jsonErr := json.Unmarshal(file, &SiteConfiguration)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
