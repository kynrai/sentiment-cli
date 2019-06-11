package cmd

import (
	"fmt"
	"os"

	"github.com/kynrai/sentiment-cli/cmd/sentiment/cmd/sweep"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sweep.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "sentiment",
	Short: "Sentiment is a sentiment analysis tool",
	Long:  "A fast and simple sentiment CLI for twitter",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
