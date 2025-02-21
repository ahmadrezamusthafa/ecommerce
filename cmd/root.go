package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ecommerce",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.GetConfig()

	b, _ := json.Marshal(cfg)
	fmt.Println(string(b))
}
