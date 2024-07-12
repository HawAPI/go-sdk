package main

import (
	"github.com/HawAPI/go-sdk/pkg/hawapi"
	"github.com/google/uuid"
)

func main() {
	client := hawapi.NewClient()
	client.WithOpts(hawapi.Options{
		Token:    "12",
		Language: "fr-FR",
		Endpoint: "http://localhost:8080/api",
	})

	_, _ = client.ListEpisodes()
	_, _ = client.FindEpisode(uuid.New())
	_, _ = client.RandomEpisode()
	_, _ = client.CreateEpisode(hawapi.CreateEpisode{})
	_, _ = client.PatchEpisode(uuid.New(), hawapi.PatchEpisode{})
	_ = client.DeleteEpisode(uuid.New())
}
