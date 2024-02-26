package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"todos/api"
	"todos/auth-service"
	"todos/evnt-service"
	"todos/rest-service"
	"todos/website"
)

var (
	addr = flag.String("http", "0.0.0.0:8000", "http address for listening")
)

func main() {
	flag.Parse()
	log.Printf("Listening for HTTP traffic at http://%s\n", *addr)
	http.Handle("/", website.Connect("http://"+*addr))
	http.Handle("/auth/", http.StripPrefix("/auth", auth.Service()))
	http.Handle("/todo/", http.StripPrefix("/todo", rest.Service()))
	http.Handle("/evnt/", http.StripPrefix("/evnt", evnt.Service()))
	go setup(api.NewClient("http://"+*addr, ""))
	exitOnErr(http.ListenAndServe(*addr, nil))
}

func setup(client *api.Client) {
	time.Sleep(500 * time.Millisecond)
	token, err := client.Signup("connor@testing.com", "password 123")
	exitOnErr(err)
	client = api.NewClient(client.Host, token)
	client.CreateTodo("User Login through [Provider]")
	client.CreateTodo("Deploy Runtime via [Platform]")
	client.CreateTodo("Cluster Events via [Broker]")
	client.CreateTodo("Persist Todos with [Database]")
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
