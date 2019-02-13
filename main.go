package main

import (
	//"log"
	re "main/modules/REST"
	"main/modules/timeoutdrop"

	//	res "RESTSITE"
	//"fmt"
	"os"
)

//export

//var s mongo.SessionGame
func main() {
	timeoutdrop.Initialize()
	//mongo.InitiateSession()

	port := os.Getenv("PORT")
	tlsos := os.Getenv("TLSUSE")
	tls := false
	var ch chan bool
	if tlsos != "" {
		tls = true
	}
	port = ":8000"
	//port1 := ":7000"
	port = "localhost:8000"
	go re.GoServerListen(port, tls)
	<-ch
	//go res.GoServerListen(port1, tls)
	// var guessColor string
	// for {
	// 	if _, err := fmt.Scanf("%s", &guessColor); err != nil {
	// 		fmt.Printf("%s\n", err)
	// 		return
	// 	}
	// 	if "exit" == guessColor {
	// 		os.Exit(0)
	// 		return
	// 	}
	// 	if "drop" == guessColor {
	// 		re.DropBase()
	// 		return
	// 	}
	// }
}
