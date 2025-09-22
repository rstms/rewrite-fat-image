/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/rstms/ffs/image"
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
		ftype := 12
		switch {
		case ViperGetBool("12"):
			ftype = 12
		case ViperGetBool("16"):
			ftype = 16
		case ViperGetBool("32"):
			ftype = 32
		}
		size := ViperGetInt64("size")
		src := args[0]
		dst := args[1]
		err := image.RewriteImage(dst, src, ftype, int64(size))
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
	CobraInit(rootCmd)
	OptionSwitch(rootCmd, "12", "", "set FAT12 type")
	OptionSwitch(rootCmd, "16", "", "set FAT16 type")
	OptionSwitch(rootCmd, "32", "", "set FAT32 type")
	OptionInt(rootCmd, "size", "", 2880*1024, "image size")
	rootCmd.MarkFlagsMutuallyExclusive("12", "16", "32")
}
