package main

import (
	"daily/go/controller"
	"daily/go/logger"
	"net/http"
)

func init() {
	period()
}

func period() {
	http.HandleFunc("/api/period/saveRecord", controller.SaveRecord)
	http.HandleFunc("/api/period/getLastRecord", controller.GetLastRecord)
	http.HandleFunc("/api/period/listRecord", controller.ListRecord)
}

func main() {

	startServer()
}

func startServer() {
	logger.Info.Println("start server...")
	err := http.ListenAndServe(":8963", nil)
	if err != nil {
		logger.Error.Println(err)
	}


}
