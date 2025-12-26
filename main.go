// Package neatapi provides a command-line entry point for the NeatLogic API SDK.
// This package demonstrates how to use the NeatLogic SDK with example operations.
package neatapi

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hejingwen098/neatapi/neatlogic"
)

// main is the entry point for the application.
// It demonstrates various operations using the NeatLogic SDK.
func main() {
	// Initialize neatClient
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	configPath := flag.String("config", filepath.Join(execDir, "config.yml"), "Config file path")

	neatClient := neatlogic.NewNeatClientWithConfigPath(*configPath)
	cientity, err := neatClient.SearchCientityByKeyword(1491357231226880, "keyword")
	if err != nil {
		fmt.Printf("Error searching entities: %v\n", err)
	} else {
		fmt.Printf("Found %d entities matching keyword '%s'\n", len(cientity), "keyword")
	}
	fmt.Println("NeatLogic SDK example completed.")
}
