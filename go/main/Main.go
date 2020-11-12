package main

import (
	"daily/go/controller"
	"log"
	"net/http"
)

func init() {
	period()
}

func period() {
	http.HandleFunc("/period/saveRecord", controller.SaveRecord)
	http.HandleFunc("/period/getLastRecord", controller.GetLastRecord)
}

func main() {

	startServer()

}

func startServer() {
	err := http.ListenAndServe(":8963", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	} else {
		log.Println("start server...")
	}
}
