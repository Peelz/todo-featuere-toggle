package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	featurehubSDK "github.com/featurehub-io/featurehub-go-sdk"
	"github.com/featurehub-io/featurehub-go-sdk/pkg/models"
	streamingclient "github.com/featurehub-io/featurehub-go-sdk/pkg/streaming-client"
	"log"
	"monopeelz/pocff-go/config"
	"net/http"
	"os"
	"time"
)

const (
// featurehubApiKey  = "default/54a87485-5a1d-46d6-95ce-868775a16c73/5f83lDy6ppvZRE7KJQjlN24GSFll9i*a8bOSz7sm5B2UMPJg9fO"
// featurehubAddress = "http://192.168.0.25:8085"
// listenAddress     = ":8080"
// logLevel          = logrus.TraceLevel
)

func main() {
	ch := make(chan time.Time)

	fmt.Println(os.Getenv("ENV_FILE"))
	// if os.Getenv("ENV_FILE") != "" {
	// 	if err := godotenv.Load(os.Getenv("ENV_FILE")); err != nil {
	// 		fmt.Println("Load env fail")
	// 	}
	// }
	// featurehubApiKey := os.Getenv("FEATUREHUB_API_KEY")
	// if featurehubApiKey == "" {
	// 	log.Fatal("featurehub api key not exist")
	// 	os.Exit(1)
	// }
	cnf := config.LoadConfig()

	client, err := streamingclient.NewStreamingClient(featurehubSDK.New(cnf.FeaturehubEdgeUrl, cnf.FeaturehubClientApiKey))
	if err != nil {
		log.Fatal(err)
	}
	client.AddNotifierFeature("save_todo", func(state *models.FeatureState) {
		// log.Println("list_todo change", state.Value, state.Type, state.Strategies)
		// log.Printf("%+v\n", state.Strategies)
		stateJson, _ := json.Marshal(state)
		_, _ = http.Post("https://webhook.site/89f3b394-f928-42a7-a982-87cfda40a722", "application/json", bytes.NewBuffer(stateJson))
		ch <- time.Now()
	})
	client.ReadinessListener(func() {
		fmt.Println("starting listen change from featurehub")
		ch <- time.Now()
	})
	client.Start()

	for {
		select {
		case timestamp := <-ch:
			fmt.Println("received something", timestamp.String())
		}
	}
}
