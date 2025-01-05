package controller

import (
	"context"
	"log"
	"os"
	"os/exec"
	"quick_launch_backend/internal/controller/dto"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func handleTextCmd(ctx context.Context, c *app.RequestContext) {
	req := dto.TextCmdRequest{}
	if err := c.Bind(&req); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// 从配置文件或环境变量中获取 Steam 路径
	steamPath := os.Getenv("STEAM_PATH")
	if steamPath == "" {
		steamPath = `D:\Program Files (x86)\Steam\Steam.exe`
	}

	// 创建命令并设置上下文超时
	cmd := exec.CommandContext(ctx, steamPath)

	// 设置超时时间
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start Steam: %v", err)
		c.JSON(consts.StatusInternalServerError, map[string]string{"error": "Failed to start Steam"})
		return
	}

	// 等待命令完成（可选）
	if err := cmd.Wait(); err != nil {
		log.Printf("Steam process exited with error: %v", err)
	}

	c.JSON(consts.StatusOK, req)
}
