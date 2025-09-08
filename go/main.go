// main.go
package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var configFile string
	var configPath string

	var rootCmd = &cobra.Command{
		Use:   "shroobago",
		Short: "shroobada deployment",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig(configFile, configPath)
		},
		Run: func(cmd *cobra.Command, args []string) {
			name := viper.GetString("name")
			age := viper.GetInt("age")
			fmt.Printf("Hello %s, you are %d years old!\n", name, age)
		},
	}

	// Add config flags
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "Config file name (supports .yaml and .json)")
	rootCmd.PersistentFlags().StringVar(&configPath, "config-path", ".", "Config file path")

	// Add application flags
	rootCmd.Flags().StringP("name", "n", "World", "Name to greet")
	rootCmd.Flags().IntP("age", "a", 0, "Your age")

	// Bind flags to viper
	viper.BindPFlags(rootCmd.Flags())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig(configFile, configPath string) {
	if configFile != "" {
		// Build full path to config file
		fullPath := filepath.Join(configPath, configFile)

		// Use SetConfigFile - it automatically detects type from extension
		viper.SetConfigFile(fullPath)

		// Try to read config file
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Warning: Could not read config file: %v\n", err)
		} else {
			fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
		}
	}
}
