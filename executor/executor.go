package executor

import (
	"github.com/savaki/gremlind/manifest"
	"io"
	"os/exec"
)

func New(id string, p manifest.Program, stdout, stderr io.Writer) *Executor {
	return &Executor{
		id:      id,
		program: p,
		stdout:  stdout,
		stderr:  stderr,
	}
}

type Executor struct {
	id      string
	program manifest.Program
	stdout  io.Writer
	stderr  io.Writer
}

func (e *Executor) Id() string {
	return e.id
}

func (e *Executor) Run() error {
	return e.runOnce()
}

func (e *Executor) runOnce() error {
	cmd := exec.Command(e.program.Cmd[0], e.program.Cmd[1:]...)
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr

	return cmd.Run()
}
