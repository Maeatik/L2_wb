package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "go-calendar [command]",
}

// Execute запускаем сервер
func Execute() {
	rootCmd.AddCommand(apiServerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}
