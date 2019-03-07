package main

import (
	//"log"
	re "main/modules/RESTSITE"
	"main/modules/config"

	//"fmt"
	"os"
)

//export
//var s mongo.SessionGame
func main() {
	//mongo.InitiateSession()
	config.Config_init("/modules/config/config.json")
	port := os.Getenv("PORT")
	tlsos := os.Getenv("TLSUSE")
	tls := false
	var ch chan bool
	if tlsos != "" {
		tls = true
	}
	port = ":80"
	//port = "localhost:8001"
	go re.GoServerListen(port, tls)
	//var guessColor string
	<-ch
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
