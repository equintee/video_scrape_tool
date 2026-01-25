package __

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var musicRecognitionService MusicRecoginitonServiceClient

func init() {
	con, err := grpc.NewClient("shazam-api:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := NewMusicRecoginitonServiceClient(con)
	musicRecognitionService = client
}

func GetMusicRecognitionService() MusicRecoginitonServiceClient {
	return musicRecognitionService
}
