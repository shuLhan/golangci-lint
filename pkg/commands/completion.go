package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (e *Executor) initCompletion() {
	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "Output completion script",
	}
	e.rootCmd.AddCommand(completionCmd)

	bashCmd := &cobra.Command{
		Use:   "bash",
		Short: "Output bash completion script",
		RunE:  e.executeBashCompletion,
	}
	completionCmd.AddCommand(bashCmd)

	zshCmd := &cobra.Command{
		Use:   "zsh",
		Short: "Output zsh completion script",
		RunE:  e.executeZshCompletion,
	}
	completionCmd.AddCommand(zshCmd)
}

func (e *Executor) executeBashCompletion(cmd *cobra.Command, args []string) error {
	err := cmd.Root().GenBashCompletion(os.Stdout)
	if err != nil {
		return fmt.Errorf("%s: unable to generate bash completions", err.Error())
	}

	return nil
}

func (e *Executor) executeZshCompletion(cmd *cobra.Command, args []string) error {
	err := cmd.Root().GenZshCompletion(os.Stdout)
	if err != nil {
		return fmt.Errorf("unable to generate zsh completions: %w", err)
	}
	// Add extra compdef directive to support sourcing command directly.
	// https://github.com/spf13/cobra/issues/881
	// https://github.com/spf13/cobra/pull/887
	fmt.Println("compdef _golangci-lint golangci-lint")

	return nil
}
