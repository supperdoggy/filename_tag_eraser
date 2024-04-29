# File Renamer

This Go program renames files in a specified folder by removing the text "[SPOTIFY-DOWNLOADER.COM] " from their names.

## Usage

1. Compile the program using the Go compiler:
`go build rename.go`

2. Run the compiled binary, providing the folder path as an argument:
`./rename <folder_path>`

## Description

The program reads the list of files in the specified folder. For each file whose name contains the text "[SPOTIFY-DOWNLOADER.COM]", it generates a new name by removing this text. Then, it renames the file accordingly.

## Example

Suppose there's a folder named "songs" containing the following files:
- "song1 [SPOTIFY-DOWNLOADER.COM].mp3"
- "song2 [SPOTIFY-DOWNLOADER.COM].mp3"
- "song3.mp3"

After running the program with the folder path "songs", the files will be renamed as follows:
- "song1.mp3"
- "song2.mp3"
- "song3.mp3"

## Notes

- This program assumes that the specified folder exists and that the program has permission to read from and write to it.
- It also assumes that the Go compiler is installed on the system.
- If no folder path is provided as an argument, the program will display a message prompting the user to provide one.
