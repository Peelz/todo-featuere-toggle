package todohdl

import (
	"fmt"
	"github.com/featurehub-io/featurehub-go-sdk/pkg/models"
	streamingclient "github.com/featurehub-io/featurehub-go-sdk/pkg/streaming-client"
	"github.com/gin-gonic/gin"
	"monopeelz/pocff-go/internal/service/todosvc"
)

type Handler struct {
	TodoUseCase *todosvc.Service
	fhClient    *streamingclient.ClientWithContext
}

func (h Handler) RandomHandler(ctx *gin.Context) {
	userKey := ctx.Query("user")
	ctxModel := &models.Context{
		Userkey: userKey,
	}
	result, err := h.fhClient.WithContext(ctxModel).GetBoolean("random")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, gin.H{
			"message": "400",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"value": result,
	})
}

func (h Handler) ListTodoHandler(ctx *gin.Context) {
	result, _ := h.TodoUseCase.ListTodo(ctx.Request.Context())
	ctx.JSON(200, result)
}

func (h Handler) SaveTodoHandler(ctx *gin.Context) {
	result, _ := h.TodoUseCase.SaveTodo(ctx.Request.Context())
	ctx.JSON(200, result)
}

func (h Handler) DeleteTodoHandler(ctx *gin.Context) {
	result, _ := h.TodoUseCase.SaveTodo(ctx.Request.Context())
	ctx.JSON(200, result)
}

func (h Handler) PartialUpdateTodoHandler(ctx *gin.Context) {
	result, _ := h.TodoUseCase.SaveTodo(ctx.Request.Context())
	ctx.JSON(200, result)
}

func New(u *todosvc.Service, fhClient *streamingclient.ClientWithContext) Handler {
	return Handler{
		TodoUseCase: u,
		fhClient:    fhClient,
	}
}
