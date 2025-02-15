package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jordanadams/tagsky/internal/ingestor"
	"github.com/spf13/cobra"
)

var ingestorCmd = cobra.Command{
	Use: "ingestor",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, _ := signal.NotifyContext(cmd.Context(), os.Kill, syscall.SIGABRT, syscall.SIGTERM)

		err := ingestor.Start(ctx)
		if err != nil {
			return fmt.Errorf("failed to start ingestor: %w", err)
		}

		return nil
	},
}
