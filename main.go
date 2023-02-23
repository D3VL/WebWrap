package main

import (
	"os"
	"os/signal"
	"embed"
	"net"
	"strconv"
)

// import internal packages
import (
    "D3VL/WebWrap/packages/logging"
    "D3VL/WebWrap/packages/open-browser"
    "D3VL/WebWrap/packages/server"
)

//go:embed static
var embeddedFS embed.FS

func findAvailablePort() (int, error) {
	address, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}


func main() {
	// listen for interrupt signal
	interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

	 
    log.EnableVerbose()
	log.Debug("WebWrap Initializing...")

	log.Info("Starting...")

	// generate a random port number
	portInt, err := findAvailablePort()
	if err != nil {
		log.Error("Error finding an available port: " + err.Error())
		return
	}

	port := strconv.Itoa(portInt)
	
	go func() {
        server.Start(embeddedFS, port);
    }()
	
	browser.Open("http://127.0.0.1:" + port)

	log.Info("When you are done, close this window to exit")

	for {
		select {
            case <-interrupt:
                log.Debug("Process interrupted")
                return
		}
	}
}