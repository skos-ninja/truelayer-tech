package main

import (
	"github.com/skos-ninja/truelayer-tech/svc/pokemon"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "truelayer",
}

func init() {
	// Register a flag for all commands that allows setting the http port
	root.PersistentFlags().Int("port", 8080, "Port to listen on")

	// Add sub commands here
	root.AddCommand(pokemon.CMD)
}

func main() {
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
