package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/cbhakar/mphasis/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	startApp()
}

func startApp() {

	r := mux.NewRouter()
	r.HandleFunc("/image", api.AddImage).Methods("POST")
	r.HandleFunc("/image", api.GetImages).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("error starting server")
	}
	log.Println("server started at port : 8080")

	go HandleOSSignals(func() {
		err := api.Stop()
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("closing app")
}

func HandleOSSignals(fn func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals,  syscall.SIGINT, syscall.SIGTERM)

	for sig := range signals {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			fn()
		}
	}
}
