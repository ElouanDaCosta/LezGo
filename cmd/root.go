/*
Copyright © 2024 NAME HERE elouandacostapeixoto@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "lezgo",
	Short:   "A simple build and dependency management tool for Go, offering multiple build profiles and enhanced workflows for developers",
	Long: `lezgo is an advanced build and dependency management CLI tailored for Go projects. 
	It simplifies your development workflow with features like multiple build profiles, 
	cross-compilation, and streamlined project management. 
	Whether you’re building for debugging, release, or testing,
	lezgo adapts to your needs with intuitive commands and configurations.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lezgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
