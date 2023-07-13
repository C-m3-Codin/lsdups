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
	"runtime"
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
	hash  string
}

func lsdups(cmd *cobra.Command, args []string) {
	fmt.Println("lsdups called")
	start := time.Now() // Record the current time before the program logic

	// Place your program logic here
	// ...
	fileMap := make(map[string]files)
	worker_count := runtime.GOMAXPROCS(0)
	// worker_count := 2

	work := make(chan string)
	files_done := make(chan files)
	worker_completed := make(chan bool)
	collection_done := make(chan bool)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < worker_count; i++ {

		go worker_fileHash(i, work, files_done, worker_completed)
	}

	go write_hashMap(fileMap, files_done, collection_done)

	err = filepath.Walk(dir, getFiles(fileMap, work))

	// fmt.Println("Directory is ", dir)
	if err != nil {
		log.Fatal(err)
	}
	close(work)
	for i := 0; i < worker_count; i++ {

		<-worker_completed
	}
	close(files_done)

	<-collection_done

	elapsed := time.Since(start) // Calculate the time elapsed since the start time
	fmt.Printf("Program execution time: %s\n", elapsed)

	// listAllHashes(fileMap)

}

func write_hashMap(fileMap map[string]files, files_done chan files, collection_done chan bool) {

	for i := range files_done {

		// fmt.Println("file hash is ", i.hash)
		val, ok := fileMap[i.hash]
		if ok {
			// fmt.Println("repeat found ")
			val.Path = append(val.Path, i.Path...)
			val.Count++
			fileMap[i.hash] = val

		} else {
			fileMap[i.hash] = i

		}
		// newFile := files{Path: []string{path}, Count: 1}

	}

	// fmt.Println("collection done")
	collection_done <- true

}

func worker_fileHash(workerID int, work chan string, done chan files, worker_completed chan bool) {
	for i := range work {
		// fmt.Println(i, " done by ", workerID)
		fileHash := fileHash(i)
		file := files{Path: []string{i}, Count: 1, hash: fileHash}
		done <- file

	}
	// fmt.Println("\n\n", workerID, " worker Completed")
	worker_completed <- true
}

func getFiles(fileMap map[string]files, work chan string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		if !info.IsDir() {
			// fileHash := fileHash(path)
			work <- path

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
