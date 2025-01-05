package controller

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterGroupRoute(h *server.Hertz) {
	// Simple group: v1
	v1 := h.Group("/v1")
	{
		v1.POST("/handleTextCmd", handleTextCmd)
	}
}
