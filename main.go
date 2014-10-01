package main

import (
	"github.com/codegangsta/cli"
	"github.com/savaki/broadcast"
	"github.com/savaki/gremlind/executor"
	"github.com/savaki/gremlind/manifest"
	"log"
	"os"
)

const (
	FieldManifest    = "manifest"
	FieldLoggingPort = "logging-port"
)

func main() {
	app := cli.NewApp()
	app.Name = "gremlind"
	app.Usage = "gremlin daemon to run inside a container"
	app.Flags = []cli.Flag{
		cli.StringFlag{FieldManifest, "/etc/gremlin/gremlin.hcl", "the name of the gremlin manifest file", "GREMLIN_MANIFEST"},
		cli.IntFlag{FieldLoggingPort, 0, "the port to listen to for the firehost; speak websocket client", "GREMLIN_LOGGING_PORT"},
	}
	app.Action = Run
	app.Run(os.Args)
}

func Run(c *cli.Context) {
	filename := c.String(FieldManifest)
	m, err := manifest.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// 1. setup all the loggers
	publishers := map[string]broadcast.Publisher{}
	for id, _ := range m.Program {
		publisher := broadcast.New()
		publisher.Start()
		publishers[id] = publisher

		response := make(chan *broadcast.Subscription, 1)
		publisher.Subscribe(response)
		subscription := <-response
		go func() {
			for {
				data := <-subscription.Receive
				os.Stdout.Write(data)
			}
		}()
	}

	// 2. run configuration scripts

	// 3. start the application
	for id, program := range m.Program {
		publisher := publishers[id]
		e := executor.New(id, program, publisher, publisher)
		e.Run()
	}

	// 4. begin service checks
}
