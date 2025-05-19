package util

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func ParseVideo(data []byte) {
	file, _ := os.Create("video.mp4")
	defer file.Close()

	_, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	file.Sync()
}

func ParseVideoFromUrl(url string) {
	data := FetchVideo(url)
	ParseVideo(data)
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
