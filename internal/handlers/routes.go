package handlers

import (
	"github.com/gin-gonic/gin"
	"monopeelz/pocff-go/internal/handlers/todohdl"
)

type ApplicationHandler struct {
	todoHandler todohdl.Handler
}

func NewApplicationHandler(todoHandler todohdl.Handler) *ApplicationHandler {
	return &ApplicationHandler{todoHandler: todoHandler}
}

func GinRouteRegister(e *gin.Engine, h *ApplicationHandler) {
	// e.GET("/ping", h.todoHandler.ListTodoHandler)
	// e.Use(func(ctx *gin.Context) {
	// 	userID := ctx.GetHeader("X-User-ID")
	// 	m := &models.Context{
	// 		Userkey: userID,
	// 	}
	// 	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), featurehub.ContextModel, &m))
	// })
	e.GET("/random", h.todoHandler.RandomHandler)
	e.GET("/todo", h.todoHandler.ListTodoHandler)
	e.POST("/todo", h.todoHandler.SaveTodoHandler)
	e.DELETE("/todo/:id", h.todoHandler.DeleteTodoHandler)
	e.PATCH("/todo/:id", h.todoHandler.PartialUpdateTodoHandler)
}
