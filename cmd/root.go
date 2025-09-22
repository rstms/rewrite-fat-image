/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "0.0.2",
	Use:     "rewrite-fat-image SRC_IMAGE DST_IMAGE",
	Short:   "rewrite a FAT disk image",
	Long: `
Reads a FAT formatted disk image from SRC_IMAGE.
Use the ffs library to generate a new FAT12 2.88MB output image
Use mtools mcat to copy the files from the source image to the destination
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := RewriteFATImage(args[0], args[1])
		cobra.CheckErr(err)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
}
