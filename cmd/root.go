/*
Copyright © 2024 GABRIEL DIONIZIO PEREIRA dionizio.0808@gmail.com
*/
package cmd

import (
	"os"

	"github.com/Dionizio8/go-stress/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-stress",
	Short: "Stress test - go expert",
	Long: `CLI system in Go to perform load tests on a web service. 
	
The user must provide the service URL, the total number of requests and the number of simultaneous calls.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		log, _ := cmd.Flags().GetBool("log")

		cst := usecase.NewCreateStress(url, requests, concurrency, log)
		cst.Execute()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL of the service to be tested")
	rootCmd.Flags().IntP("requests", "r", 1, "Total number of requests")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent requests")
	rootCmd.Flags().BoolP("log", "l", false, "Log of each request in real time")
	rootCmd.MarkFlagRequired("url")
}
