package main

import (
	"errors"
	"log"
	"os"

	"github.com/juanjoseamaker/sharefile/app"
)

func main() {
	config, err := parseCommandLine(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if config.RunServerOrClient {
		err = app.RunServer(config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = app.RunClient(config)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseCommandLine(args []string) (*app.Config, error) {
	config := new(app.Config)

	if len(args) < 3 {
		return nil, errors.New("invalid arguments")
	}

	switch args[1] {
	case "receive":
		config.RunServerOrClient = true
	case "send":
		config.RunServerOrClient = false

		if len(args) != 4 {
			return nil, errors.New("invalid arguments")
		}

		config.ClientPath = args[3]
	}

	config.ServerAddr = args[2]

	return config, nil
}
