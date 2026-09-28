package main

import (
	"os"
	"strings"

	"github.com/edheltzel/gap/internal/tui"
	"github.com/spf13/cobra"
)

func main() {
	if err := execute(os.Args[1:]); err != nil {
		_ = tui.Print(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func execute(args []string) error {
	cmd := newRoot()
	cmd.SetArgs(args)
	return cmd.Execute()
}

func newRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gap",
		Short: "Install Git-hosted agent distributions",
		Long:  "gap fetches a Git-hosted agent distribution, reads gap.toml, and installs listed files. Thin, idempotent installer — not a runtime.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SetHelpFunc(func(c *cobra.Command, _ []string) {
		if err := tui.Print(c.OutOrStdout(), helpText(c)); err != nil {
			c.PrintErrln(err)
		}
	})
	return cmd
}

func helpText(cmd *cobra.Command) string {
	var b strings.Builder
	if long := strings.TrimSpace(cmd.Long); long != "" {
		b.WriteString(long)
		b.WriteString("\n\n")
	}
	b.WriteString(cmd.UsageString())
	return strings.TrimSpace(b.String()) + "\n"
}
