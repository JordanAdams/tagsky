package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tagsky",
}

func Execute(ctx context.Context) (err error) {
	rootCmd.AddCommand(&ingestorCmd)

	err = rootCmd.ExecuteContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
