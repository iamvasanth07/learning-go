package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

//Item is...
type Item struct {
	Title string
	Body  string
}

var database []Item

//API type...
type API int

//AddItem is ...
func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

//GetDB is ...
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Fatal("Exiting due to manual interuption...")
		os.Exit(1)
	}()

	if err != nil {
		log.Fatal("Error regestering rpc server:", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Listener error:", err)
	}

	defer listener.Close()

	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Serving Error :", err)
	}

}
