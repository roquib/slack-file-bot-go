package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// Set environment variables for token and channel ID
	os.Setenv("SLACK_BOT_TOKEN", "your-token")
	os.Setenv("CHANNEL_ID", "channelid here")

	// Path to the file
	filePath := "plank.md"

	// Read the file content into memory
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error: Unable to read file:", err)
		return
	}

	// Check if the file is empty
	if len(fileContent) == 0 {
		fmt.Println("Error: File is empty.")
		return
	}

	// Create a new Slack client
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// Prepare the file content as a bytes.Reader
	fileReader := bytes.NewReader(fileContent)

	// Define parameters for the file upload
	params := slack.UploadFileV2Parameters{
		Channel:  os.Getenv("CHANNEL_ID"),
		Filename: "plank.md", // Set the filename in Slack
		Reader:   fileReader, // Pass the file content as a stream
	}

	// Upload the file using UploadFileV2
	uploadedFile, err := api.UploadFileV2(params)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	// Print file details on success
	fmt.Println("File uploaded successfully!")
	fmt.Println("ID:", uploadedFile.ID, "Title:", uploadedFile.Title)
}
