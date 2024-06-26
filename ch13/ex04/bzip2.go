package bzip

import (
	"io"
	"os/exec"
)

type Writer struct {
	cmd   exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) (io.WriteCloser, error) {
	cmd := exec.Cmd{Path: "/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &Writer{cmd, stdin}, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *Writer) Close() error {
	if err := w.stdin.Close(); err != nil {
		return err
	}
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
