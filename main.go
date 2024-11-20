package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	flag.Parse()

	if _, err := os.Stat("./static"); err != nil {
		log.Fatalf("failed to stat static dir: %s", err)
	}

	var fn func() error

	switch flag.Arg(0) {
	case "update-talks":
		fn = updateTalks
	default:
		fn = serve
	}

	err := fn()
	if err != nil {
		log.Fatalf("error running handler: %s", err)
	}
}
