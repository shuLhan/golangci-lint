package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/golangci/golangci-lint/pkg/lint/linter"
)

func (e *Executor) initLinters() {
	e.lintersCmd = &cobra.Command{
		Use:   "linters",
		Short: "List current linters configuration",
		Run:   e.executeLinters,
	}
	e.rootCmd.AddCommand(e.lintersCmd)
	e.initRunConfiguration(e.lintersCmd)
}

func (e *Executor) executeLinters(_ *cobra.Command, args []string) {
	if len(args) != 0 {
		e.log.Fatalf("Usage: golangci-lint linters")
	}

	enabledLintersMap, err := e.EnabledLintersSet.GetEnabledLintersMap()
	if err != nil {
		log.Fatalf("Can't get enabled linters: %s", err)
	}

	fmt.Println("Enabled by your configuration linters:")
	enabledLinters := make([]*linter.Config, 0, len(enabledLintersMap))
	for _, linter := range enabledLintersMap {
		enabledLinters = append(enabledLinters, linter)
	}
	printLinterConfigs(enabledLinters)

	var disabledLCs []*linter.Config
	for _, lc := range e.DBManager.GetAllSupportedLinterConfigs() {
		if enabledLintersMap[lc.Name()] == nil {
			disabledLCs = append(disabledLCs, lc)
		}
	}

	fmt.Println("\nDisabled by your configuration linters:")
	printLinterConfigs(disabledLCs)

	os.Exit(0)
}
