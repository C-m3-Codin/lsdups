/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

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

type files struct {
	Path  []string
	Count int
}

func lsdups(cmd *cobra.Command, args []string) {
	start := time.Now() // Record the current time before the program logic

	// Place your program logic here
	// ...

	fmt.Println("lsdups called")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileMap := make(map[string]files)
	err = filepath.Walk(dir, getFiles(fileMap))
	fmt.Println("Directory is ", dir)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start) // Calculate the time elapsed since the start time
	fmt.Printf("Program execution time: %s\n", elapsed)
	listAllHashes(fileMap)

}

func getFiles(fileMap map[string]files) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		if !info.IsDir() {
			fileHash := fileHash(path)
			fmt.Println("file hash is ", fileHash)
			val, ok := fileMap[fileHash]
			if ok {
				val.Path = append(val.Path, path)
				val.Count++
				fileMap[fileHash] = val
				return nil
			}
			newFile := files{Path: []string{path}, Count: 1}
			fileMap[fileHash] = newFile
		}
		return nil
	}
}

func fileHash(filepath string) (hash string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func listAllHashes(filemap map[string]files) {
	for k, v := range filemap {
		// fmt.Println(k, v)
		if v.Count > 1 {
			fmt.Println(k, v)
		}
	}
}
