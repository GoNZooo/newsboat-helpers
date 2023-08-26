package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/GoNZooo/newsboat-helpers/header"
)

func main() {
	home := os.Getenv("HOME")
	root := flag.String(
		"root",
		path.Join(home, ".newsboat/downloads"),
		"Root directory for podcast downloads",
	)
	overwrite := flag.Bool(
		"overwrite",
		false,
		"Overwrite existing files",
	)
	verbose := flag.Bool(
		"verbose",
		false,
		"Turn on verbose output",
	)

	flag.Parse()

	header, err := header.Parse(os.Stdin)
	if err != nil {
		_ = fmt.Errorf("Error parsing header: %v", err)
		os.Exit(1)
	}

	os.MkdirAll(*root, 0755)
	filepath := header.AsFilepath()
	path := path.Join(*root, filepath)
	downloadTo(header.PodcastUrl, path, *overwrite, *verbose)
}

func downloadTo(url string, filepath string, overwrite bool, verbose bool) error {
	if !overwrite && fileExists(filepath) {
		if verbose {
			fmt.Printf("File '%s' already exists, skipping\n", filepath)
		}

		return nil
	}

	directory := path.Dir(filepath)
	os.MkdirAll(directory, 0755)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	return !errors.Is(err, os.ErrNotExist)
}
