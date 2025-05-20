
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/iamgaru/human-presence-sim/internal/core"
	"github.com/iamgaru/human-presence-sim/internal/registry"
	"github.com/iamgaru/human-presence-sim/internal/telemetry"
)

func main() {
	_ = godotenv.Load()

	args := os.Args
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	runner := setupRunner()

	switch args[1] {
	case "--list-plugins":
		handleListPlugins()
	case "--validate-config":
		handleValidateConfig(args)
	default:
		handleExecuteRoutine(args[1], runner)
	}
}

func printUsage() {
	fmt.Println("Usage: simulator [--list-plugins|--validate-config <file>|<config.yaml|.json>]")
}

func setupRunner() *core.Runner {
	runner := core.NewRunner()
	for name, factory := range registry.All {
		runner.RegisterPlugin(factory())
	}
	return runner
}

func handleListPlugins() {
	telemetry.Log("Listing plugins", map[string]interface{}{"plugins": registry.All})
	for name := range registry.All {
		fmt.Printf(" - %s
", name)
	}
}

func handleValidateConfig(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: --validate-config <file.yaml|.json>")
		os.Exit(1)
	}
	path := args[2]
	if routine, err := core.LoadRoutine(path); err != nil {
		telemetry.Log("Config validation failed", map[string]interface{}{"error": err.Error()})
		os.Exit(1)
	} else {
		telemetry.Log("Config is valid", map[string]interface{}{"path": path})
		json.NewEncoder(os.Stdout).Encode(routine)
	}
}

func handleExecuteRoutine(path string, runner *core.Runner) {
	if err := runner.Run(path); err != nil {
		telemetry.Log("Execution failed", map[string]interface{}{"error": err.Error()})
		os.Exit(1)
	}
}
