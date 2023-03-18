/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorbischoff/fileutilities/pkg/audiobitrate"
)

var audiobitrateCmd = &cobra.Command{
	Use:   "audiobitrate",
	Short: "Will convert wav file bit depth to 16bit",
	Long:  `By victorbischoff`,
	Run: func(cmd *cobra.Command, args []string) {
		pathflag, _ := cmd.Flags().GetString("path")
		audiobitrate.ConvertFolder24to16Wav(pathflag)
	},
}

func init() {
	rootCmd.AddCommand(audiobitrateCmd)
	audiobitrateCmd.PersistentFlags().String("path", "", "path to directory")
}
