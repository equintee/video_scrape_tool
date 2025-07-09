package util

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
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

func SaveVideo(data []byte) (*os.File, error) {
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

func SaveVideoFromUrl(url string) (*os.File, error) {
	data := FetchVideo(url)
	video, err := SaveVideo(data)
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

func ExtractAudioFromVideo(file *os.File) ([]byte, error) {
	filePath := "C:\\Users\\equinte\\AppData\\Local\\Temp\\312159831.mp3"
	err := ffmpeg_go.Input(file.Name()).Audio().Output(filePath).Run()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return os.ReadFile(filePath)
}
