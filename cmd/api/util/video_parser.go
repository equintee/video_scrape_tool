package util

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var tempDir string

func init() {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Println("Error creating temp dir")
		panic(err)
	}
	log.Printf("Temporary diretory: %s\n", tempDir)
}

func ParseVideo(data []byte) (*os.File, error) {
	//Doesn't create under tempdir, also not deleting video after the function ends.
	file, _ := os.CreateTemp(tempDir, "*.mp4")
	defer file.Close()

	_, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	err = file.Sync()
	if err != nil {
		panic(err)
	}

	log.Printf("Parsed video file: %s\n", file.Name())
	return file, nil
}

func ParseVideoFromUrl(url string) (*os.File, error) {
	data := FetchVideo(url)
	video, err := ParseVideo(data)
	if err != nil {
		panic(err)
	}

	return video, nil
}

func FetchVideo(baseUrl string) []byte {
	parsedUrl, _ := url.Parse(baseUrl)
	request := http.Request{
		Method: "GET",
		URL:    parsedUrl,
	}

	response, err := http.DefaultClient.Do(&request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, _ := io.ReadAll(response.Body)

	return responseData
}
