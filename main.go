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

type ChanWriter struct {
	messages chan []byte
}

func (c *ChanWriter) Write(p []byte) (n int, err error) {
	c.messages <- p
	return len(p), nil
}

func Run(c *cli.Context) {
	filename := c.String(FieldManifest)
	m, err := manifest.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	publishers := map[string]broadcast.Publisher{}
	for id, program := range m.Program {
		messages := make(chan []byte, 4096)
		publisher := broadcast.New((<-chan []byte)(messages))
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

		writer := &ChanWriter{messages: messages}
		e := executor.New(id, program, writer, writer)
		e.Run()
	}

	// start or connect to the logger
	// run configuration scripts
	// start the application
	// begin service checks
}
