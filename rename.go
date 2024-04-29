package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// get folder path from arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide the folder path as an argument.")
		return
	}

	// Define the folder path where the files are located
	folderPath := os.Args[1]

	// Read the list of files in the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Loop through all files in the folder
	for _, file := range files {
		// Check if the file name contains the text "[SPOTIFY-DOWNLOADER.COM] "
		if strings.Contains(file.Name(), "[SPOTIFY-DOWNLOADER.COM]") {
			// Generate the new file name by removing the text
			newName := strings.ReplaceAll(file.Name(), "[SPOTIFY-DOWNLOADER.COM] ", "")

			// Rename the file
			err := os.Rename(folderPath+"/"+file.Name(), folderPath+"/"+newName)
			if err != nil {
				fmt.Println("Error renaming file:", err)
				continue
			}
			fmt.Printf("Renamed '%s' to '%s'\n", file.Name(), newName)
		}
	}
}
