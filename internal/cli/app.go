package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/vincentkoc/graincrawl/internal/buildinfo"
	"github.com/vincentkoc/graincrawl/internal/output"
)

type App struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (a App) Run(ctx context.Context, args []string) error {
	stdout := a.Stdout
	if stdout == nil {
		stdout = io.Discard
	}
	flags, rest := parseGlobalFlags(args)
	if flags.Help || len(rest) == 0 {
		_, err := io.WriteString(stdout, usage)
		return err
	}
	cmd, cmdArgs := rest[0], rest[1:]
	switch cmd {
	case "version":
		return a.runVersion(stdout, flags)
	case "help":
		_, err := io.WriteString(stdout, usage)
		return err
	default:
		_ = ctx
		_ = cmdArgs
		return fmt.Errorf("unknown command %q", cmd)
	}
}

func (a App) runVersion(w io.Writer, flags GlobalFlags) error {
	info := buildinfo.Current()
	if flags.JSON {
		return output.WriteEnvelope(w, info)
	}
	output.PrintKV(w, "version", info.Version)
	output.PrintKV(w, "commit", info.Commit)
	output.PrintKV(w, "date", info.Date)
	return nil
}
