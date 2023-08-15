package main

import (
	"context"
	"log"

	"github.com/make-software/casper-go-sdk/sse"
)

func main() {
	client := sse.NewClient("http://94.130.10.55:9999/events/main")
	defer client.Stop()
	client.RegisterHandler(
		sse.DeployProcessedEventType,
		func(ctx context.Context, rawEvent sse.RawEvent) error {
			deploy, err := rawEvent.ParseAsDeployProcessedEvent()
			if err != nil {
				return err
			}
			log.Printf("Deploy hash: %s", deploy.DeployProcessed.DeployHash)
			return nil
		})
	lastEventID := 1234
	client.Start(context.TODO(), lastEventID)
}
