/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golangWrappers",
	Short: "\nWrappers for common tasks in Go",
	Long:  "\nWrappers for common tasks in Go",
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golangWrappers.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	fmt.Println(`
	 __                                                 _                 
	/__  _    \    / ._ _. ._  ._   _  ._ _   |_       |_ _. |_  ._ _     
	\_| (_)    \/\/  | (_| |_) |_) (/_ | _>   |_) \/   | (_| |_) |  /_ \/ 
			       |   |                  /                    /  
	`)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
