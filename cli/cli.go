package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	explorer "github.com/andy5090/nomadcoin/explorer/templates"
	"github.com/andy5090/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Nomad Coin\n\n")
	fmt.Printf("Please use the following flags:\n")
	fmt.Printf("-port=4000:	Set the PORT of the server\n")
	fmt.Printf("-mode=rest:	Choose between 'html and 'rest'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of ther server")
	mode := flag.String("mode", "rest", "Choose between 'html and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
