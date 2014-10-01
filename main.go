package main

import (
	"github.com/codegangsta/cli"
	"github.com/savaki/gremlind/executor"
	"github.com/savaki/gremlind/manifest"
	"log"
	"os"
)

const (
	FieldManifest = "manifest"
)

func main() {
	app := cli.NewApp()
	app.Name = "gremlind"
	app.Usage = "gremlin daemon to run inside a container"
	app.Flags = []cli.Flag{
		cli.StringFlag{FieldManifest, "/etc/gremlin/gremlin.hcl", "the name of the gremlin manifest file", "GREMLIN_MANIFEST"},
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

	for id, program := range m.Program {
		e := executor.New(id, program, os.Stdout, os.Stderr)
		e.Run()
	}

	// start or connect to the logger
	// run configuration scripts
	// start the application
	// begin service checks
}
