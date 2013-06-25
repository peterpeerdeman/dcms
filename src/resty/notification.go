package resty

import (
	"container/list"
	"go.net/websocket"
	"log"
	"time"
)

type Notification struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

var NotificationChannel = make(chan Notification)

func CreateNotificationServer() func(ws *websocket.Conn) {
	var observers = list.New()

	go pinger()
	go broadcaster(observers)

	return func(ws *websocket.Conn) {
		go writer(ws, observers)
		reader(ws)
	}
}

func reader(ws *websocket.Conn) {
	for {
		var notification Notification
		err := websocket.JSON.Receive(ws, &notification)
		if err != nil {
			log.Printf("reader() %v", err)
			return
		}
		NotificationChannel <- notification
	}
}

func writer(ws *websocket.Conn, observers *list.List) {
	myChannel := make(chan Notification)
	element := observers.PushBack(myChannel)
	defer observers.Remove(element)
	for {
		select {
		case notification, ok := <-myChannel:
			if !ok {
				log.Printf("Channel closed")
				return
			}
			err := websocket.JSON.Send(ws, notification)
			if err != nil {
				log.Printf("writer() %v", err)
				return
			}
		}
	}
}

func pinger() {
	for {
		var notification Notification
		notification.Type = "ping"
		NotificationChannel <- notification
		time.Sleep(5 * time.Second)
	}
}

func Message(msg string) {
	time.Sleep(8 * time.Second)
	var notification Notification
	notification.Type = "message"
	notification.Msg = msg
	NotificationChannel <- notification
}

func broadcaster(observers *list.List) {
	for {
		notification, ok := <-NotificationChannel
		if ok {
			for e := observers.Front(); e != nil; e = e.Next() {
				client := e.Value.(chan Notification)
				client <- notification
			}
		}
	}
}
