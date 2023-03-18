/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorbischoff/fileutilities/pkg/rename"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "renames all files in directory numbered from to the amount of files",
	Long:  `By victorbischoff`,
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
