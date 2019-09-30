package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("log_file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}

	log.SetOutput(file)

	// its possible to set a formatter
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	log.Println("Something Happened!!")

	log.Trace("trace level")
	log.Info("info level")
	log.Warning("warning level")
	log.Error("errror occured")
	log.Fatal("fatal error occured")

	fmt.Println("THIS LINE WIL NOT BE EXECUTED AS THE PROCCESS RETURNED DUE TO A FATAL ERROR")
}
