/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
     log "github.com/sirupsen/logrus"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
 Use:   "files",
 Short: "Show the largest files in the given path.",
 Long:  `Quickly scan a directory and find large files.`,
 Run: func(cmd *cobra.Command, args []string) {
  for key, value := range viper.GetViper().AllSettings() {
   log.WithFields(log.Fields{
    key: value,
   }).Info("Command Flag")
  }
 },
}

func init() {
	rootCmd.AddCommand(filesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
