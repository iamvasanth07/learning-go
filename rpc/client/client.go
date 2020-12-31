package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//Item is...
type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error", err)
	}

	a := Item{"First", "First A"}
	b := Item{"Second", "Second B"}
	c := Item{"Third", "Third C"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

	client.Call("API.GetDB", "", &db)

	fmt.Println("Database :", db)

}
