/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorbischoff/fileutilities/pkg/rename"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pathflag, _ := cmd.Flags().GetString("path")
		prefixflag, _ := cmd.Flags().GetString("prefix")
		rename.RenameFilesInDirectory(pathflag, prefixflag)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().String("path", "", "path to directory")
	renameCmd.PersistentFlags().String("prefix", "", "prefix for files")
}
