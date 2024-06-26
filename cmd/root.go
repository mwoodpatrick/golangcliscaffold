/*
Copyright © 2024 Mark Wood-Patrick mwoodpatrick@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var Verbose bool
var Debug bool
var Highlight int
var Path string

var version = "0.0.1"
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
 Use:   "getsize",
 Version: version,
 Short: "List the size of a local directory.",
 Long: `This command will display the size of a directory with several different options.`,
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

    rootCmd.PersistentFlags().IntVarP(&Highlight, "highlight", "", 500, "Highlight files/directories over this threshold, in MB")
    viper.BindPFlag("highlight", rootCmd.PersistentFlags().Lookup("highlight"))

    rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "Define the path to scan.")
    rootCmd.MarkPersistentFlagRequired("path")
    viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}
