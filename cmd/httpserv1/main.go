package main

import (
	"fmt"
	featurehubSDK "github.com/featurehub-io/featurehub-go-sdk"
	"github.com/featurehub-io/featurehub-go-sdk/pkg/models"
	"github.com/joho/godotenv"
	"monopeelz/pocff-go/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"monopeelz/pocff-go/internal/adapter/featurehub"
	"monopeelz/pocff-go/internal/handlers"
	"monopeelz/pocff-go/internal/handlers/todohdl"
	"monopeelz/pocff-go/internal/service/todosvc"
)

func main() {
	if os.Getenv("ENV_FILE") != "" {
		if err := godotenv.Load(os.Getenv("ENV_FILE")); err != nil {
			fmt.Println("Load env fail")
		}
	}
	cfg := config.LoadConfig()
	logger := log.New(os.Stdout, "", -1)
	logr := logrus.New()
	logr.SetLevel(logrus.TraceLevel)
	models.SetLogger(logr)

	featurehubConf, err := featurehubSDK.
		New(cfg.FeaturehubEdgeUrl, cfg.FeaturehubClientApiKey).
		WithLogLevel(logrus.TraceLevel).
		WithWaitForData(true).
		Connect()
	if err != nil {
		log.Fatal(err)
	}
	fhClient := featurehubConf.NewContext()
	featureFlag := featurehub.NewAdapter(fhClient, logger)
	useCase := todosvc.NewService(nil, featureFlag)
	handler := todohdl.New(useCase, fhClient)
	handlerFactory := handlers.NewApplicationHandler(handler)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	handlers.GinRouteRegister(r, handlerFactory)
	log.Fatal(r.Run(fmt.Sprintf(":%s", cfg.ServicePort)))
}
