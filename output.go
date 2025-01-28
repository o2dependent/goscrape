package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/o2dependent/go-scrape/utils"
)

func createOutputFile(url string) *os.File {
	directoryValid, err := utils.DirectoryExists(output)
	if !directoryValid || err != nil {
		log.Println(errors.New("directory is invalid"))
		os.Exit(1)
	}
	f, err := os.Create(output + strings.ReplaceAll(url, "/", "") + ".txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return f
}

func generateOutput(f *os.File, emails []string, numbers []string) {
	for _, email := range emails {
		f.WriteString(email + "\n")
	}

	if collectPhoneNumbers {
		f.WriteString("\n")
		for _, number := range numbers {
			f.WriteString(number + "\n")
		}
	}
}
