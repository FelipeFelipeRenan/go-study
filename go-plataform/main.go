package main

import "plataform/logging"

func writeMessage(logger logging.Logger){
	logger.Info("Hello, Plataform")
}
func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}