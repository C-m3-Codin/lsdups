/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// lsdupsCmd represents the lsdups command
var lsdupsCmd = &cobra.Command{
	Use:   "lsdups",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: lsdups,
}

func init() {
	rootCmd.AddCommand(lsdupsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsdupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsdupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func lsdups(cmd *cobra.Command, args []string) {
	fmt.Println("lsdups called")

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(dir, getFile)

	fmt.Println("Directory is ", dir)
	if err != nil {
		log.Fatal(err)
	}

}

func getFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)

	if !info.IsDir() {

		fileHash(path)
	}
	return nil
}

func fileHash(filepath string) {
	fmt.Println("In file Hash")
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))

}
