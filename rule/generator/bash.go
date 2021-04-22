package generator

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// compile time implementation check
var _ IGenerator = (*Bash)(nil)

type Bash struct {
	Command string

	file *os.File
}

func NewBash(command string) (*Bash, error) {
	file, err := genBash(command)
	if err != nil {
		return nil, err
	}

	return &Bash{
		Command: command,
		file:    file,
	}, nil
}

func (b *Bash) Run(ctx context.Context, env map[string]string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "/bin/bash", b.file.Name())

	kvenv := make([]string, 0, len(env))
	for k, v := range env {
		kvenv = append(kvenv, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = kvenv
	return cmd
}

func (b *Bash) Generate(ctx context.Context, env map[string]interface{}) (string, error) {
	senv := make(map[string]string)
	for k, v := range env {
		// TODO: implement
		switch v := v.(type) {
		case []string:
			senv[k] = strings.Join(v, ",")
		case string:
			senv[k] = v
		}
	}
	cmd := b.Run(ctx, senv)

	output, err := cmd.Output()
	if err != nil {
		return "", ErrRunCommand
	}
	return string(output), nil
}

func (b *Bash) Close() error {
	b.file.Close()
	os.Remove(b.file.Name())
	return nil
}

func genBash(cmd string) (*os.File, error) {
	temp, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		return nil, err
	}

	io.WriteString(temp, cmd)
	return temp, nil
}
