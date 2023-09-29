/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/MiguelCiulog/extrgo/pkg"
	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts the selected audio samples and outputs them into a file.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Extract(pkg.CLIInput{
			InputFile:      inputFile,
			OutputFile:     outputFile,
			CanReplaceFile: canReplaceFile,
		})
	},
}

var AudioTypes = []string{"mp3"}

var inputFile string
var outputFile string
var canReplaceFile bool

// Checks if file exists and cleans up the file path (for using in windows/unix paths).
func doesFileExist(filePath *string) bool {
	// Use filepath.Clean to ensure the path is in the correct format
	tmp_path := filepath.Clean(*filePath)
	filePath = &tmp_path

	// Use the Stat function from the os package to check if the file exists
	_, err := os.Stat(*filePath)

	// Check if the error is due to a non-existent file
	// if os.IsNotExist(err) {
	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else if err == nil {
		return true
	} else {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(extractCmd)

	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Specify input audio file.")
	err := rootCmd.MarkPersistentFlagRequired("input")
	err = rootCmd.MarkPersistentFlagFilename("input", AudioTypes...)
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Specify output file.")
	err = rootCmd.MarkPersistentFlagFilename("output", AudioTypes...)
	if err != nil {
		panic(err)
	}

	err = rootCmd.MarkPersistentFlagRequired("output")
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().
		BoolVarP(&canReplaceFile, "replace", "r", false, "Can replace output file.")
}

func check() {
	if inputFile == outputFile {
		panic("Input and output file can't be the same.")
	}

	if !doesFileExist(&inputFile) {
		panic("Input file not found.")
	}

	if !canReplaceFile && !doesFileExist(&outputFile) {
		panic("File already exists. Set -r or --replace to overwrite file.")
	}
}
