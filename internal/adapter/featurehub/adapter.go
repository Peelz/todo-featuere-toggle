package featurehub

import (
	"context"
	"github.com/featurehub-io/featurehub-go-sdk/pkg/models"
	streamingclient "github.com/featurehub-io/featurehub-go-sdk/pkg/streaming-client"
	"log"
)

type Adapter struct {
	// feature hub client with default context
	fhClient *streamingclient.ClientWithContext
	logger   *log.Logger
}

func (a Adapter) IsEnabled(ctx context.Context, n string, opts ...interface{}) bool {
	r, err := a.clientWithContext(ctx).GetBoolean(n)
	if err != nil {
		a.logger.Println(err)
	}
	return r
}

func (a Adapter) IsEnabledWithJSON(ctx context.Context, n string, opts ...interface{}) []byte {
	// TODO implement me
	panic("implement me")
}

func (a Adapter) IsEnabledNumber(ctx context.Context, n string, opts ...interface{}) float64 {
	var userID string
	rawUserID := ctx.Value("UserID")
	if _, ok := rawUserID.(string); ok {
		log.Println(ok, rawUserID.(string))
		userID = rawUserID.(string)
	}
	log.Println(&models.Context{Userkey: userID})
	r, err := a.fhClient.WithContext(&models.Context{Userkey: userID}).GetNumber(n)
	if err != nil {
		a.logger.Println(err)
	}
	return r
}

func (a Adapter) IsEnabledString(ctx context.Context, n string, opts ...interface{}) string {
	var userID string
	rawUserID := ctx.Value("UserID")
	if _, ok := rawUserID.(string); ok {
		log.Println(ok, rawUserID.(string))
		userID = rawUserID.(string)
	}
	r, err := a.fhClient.WithContext(&models.Context{Userkey: userID}).GetString(n)
	if err != nil {
		a.logger.Println(err)
	}
	return r
}

func (a Adapter) IsEnabledWithString(ctx context.Context, n string, opts ...interface{}) string {
	// TODO implement me
	panic("implement me")
}

func (a Adapter) IsEnabledByStrategy(ctx context.Context, n string, opts ...interface{}) interface{} {
	feature, err := a.fhClient.GetFeature(n)
	if err != nil {
		a.logger.Println("Get feature error", err)
	}
	return feature
}

func (a Adapter) clientWithContext(ctx context.Context) *streamingclient.ClientWithContext {
	_fhContext := ctx.Value(ContextModel)
	fhContext, ok := _fhContext.(*models.Context)
	if ok {
		return a.fhClient.WithContext(fhContext)
	}
	return a.fhClient
}

func NewAdapter(client *streamingclient.ClientWithContext, logger *log.Logger) *Adapter {
	return &Adapter{
		fhClient: client,
		logger:   logger,
	}
}
