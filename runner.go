package ytdlp

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Runner struct {
	ExecutablePath string
}

func NewRunner(executablePath string) *Runner {
	return &Runner{ExecutablePath: executablePath}
}

func (r *Runner) Run(listener Listener, arg ...string) error {
	if arg == nil {
		return fmt.Errorf("command is nil")
	}
	if listener == nil {
		return fmt.Errorf("listener is nil")
	}
	if r.ExecutablePath == "" {
		return fmt.Errorf("executable path is empty")
	}

	command := exec.Command(r.ExecutablePath, arg...)
	command.Stderr = &lineWriter{handler: listener.OnStderr}
	command.Stdout = &lineWriter{handler: listener.OnStdout}

	if err := command.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			listener.OnExit(exitErr.ExitCode())
			return nil
		}
		return err
	}

	listener.OnExit(0)
	return nil
}

type lineWriter struct {
	buffer  bytes.Buffer
	handler func(string)
}

func (w *lineWriter) Write(p []byte) (int, error) {
	w.buffer.Write(p)

	for {
		line, err := w.buffer.ReadString('\n')
		if err != nil {
			break
		}

		if len(line) > 0 && line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		w.handler(line)
	}
	return len(p), nil
}
